package mf

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

// Request sent from client to server
type Request struct {
	ID     int    `json:"id"`
	Method string `json:"method"`
}

// Response received from server
type Response struct {
	ID      int    `json:"id"`
	Version string `json:"jsonrpc"`
}

// Connect to StratumPool
func (sp *StratumPool) Connect(addr string, port int) error {
	var err error

	sp.conn, err = net.Dial("tcp", addr+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	sp.reader = bufio.NewReader(sp.conn)

	// Check version
	if err := sp.checkVersion(); err != nil {
		return err
	}

	return nil
}

func (sp *StratumPool) checkVersion() error {
	// Request version
	p := Request{Method: "version"}
	if err := sp.send(p); err != nil {
		return err
	}

	// Get response
	resp, err := sp.receive()
	if err != nil {
		return err
	}

	// Process version
	if resp.Version == "" {
		sp.jsonRPCVersion = 1.0
	} else {
		ver, err := strconv.ParseFloat(resp.Version, 64)
		if err != nil {
			return err
		}
		sp.jsonRPCVersion = ver
	}

	fmt.Println("Version detected:", sp.jsonRPCVersion)
	return nil
}

func (sp *StratumPool) send(req Request) error {
	// Set ID
	sp.idCount++
	req.ID = sp.idCount

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

func (sp *StratumPool) receive() (Response, error) {
	var r Response

	// Get response
	txt, err := sp.reader.ReadString('\n')
	if err != nil {
		return r, err
	}
	fmt.Print("Server> ", txt)

	// Transform to JSON
	err = json.Unmarshal([]byte(txt), &r)
	return r, err
}
