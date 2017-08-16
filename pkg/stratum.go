package mf

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// StratumPool represents a pool
type StratumPool struct {
	addr     string
	port     int
	username string
	password string
	hc       http.Client
}

// NewStratumPool creates a new stratum pool
func NewStratumPool(addr string, port int) (*StratumPool, error) {
	var sp StratumPool
	sp.addr = addr
	sp.port = port
	return &sp, nil
}

// AddCredentials adds informations about account on pool
func (sp *StratumPool) AddCredentials(username, password string) {
	sp.username = username
	sp.password = password
}

// GetWork requests work from stratum pool
func (sp *StratumPool) GetWork() (string, error) {
	if sp.username == "" || sp.password == "" {
		return "", errors.New("No username or password provided")
	}

	content := []byte("{\"method\":\"getwork\",\"params\":[],\"id\":0}")
	req, err := http.NewRequest("POST", "http://"+sp.addr+":"+strconv.Itoa(sp.port), bytes.NewBuffer(content))
	if err != nil {
		return "", err
	}

	// Build request
	req.SetBasicAuth(sp.username, sp.password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", strconv.Itoa(len(content)))

	// Exec request
	resp, err := sp.hc.Do(req)
	if err != nil {
		return "", err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
