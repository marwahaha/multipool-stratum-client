package msc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

// StratumPool represents a pool
type StratumPool struct {
	jsonRPCVersion float64
	conn           net.Conn
	reader         *bufio.Reader
	idCount        int
}

// RequestInterface is an interface for requests
type RequestInterface interface {
	setID(id int)
}

// ResponseInterface is an interface for responses
type ResponseInterface interface {
}

// Connect to StratumPool
func (sp *StratumPool) Connect(addr string, port int) error {
	var err error

	sp.conn, err = net.Dial("tcp", addr+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	sp.reader = bufio.NewReader(sp.conn)

	return nil
}

func (sp *StratumPool) send(req RequestInterface) error {
	// Set ID
	sp.idCount++
	req.setID(sp.idCount)

	// Transform to text
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	txt := string(bytes) + "\n"

	// Send request
	fmt.Print("Client> ", txt)
	fmt.Fprintf(sp.conn, txt)

	return nil
}

func (sp *StratumPool) receive(resp ResponseInterface) error {
	// Get response
	txt, err := sp.reader.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Print("Server> ", txt)

	// Transform to JSON
	err = json.Unmarshal([]byte(txt), &resp)
	return err
}
