package msc

// Source: https://forum.minergate.com/viewtopic.php?f=9&t=1987

// BitcoinStratumPool represent a Bitcoin Stratum Pool
type BitcoinStratumPool struct {
	StratumPool // Inheritance
}

// BitcoinWork represent a work assign to the worker
type BitcoinWork struct {
}

// GetWork requests work from stratum pool
func (sp *StratumPool) GetWork(username, password string) (string, error) {
	resp := bitcoinResponse{}

	// Authorize
	r := bitcoinRequest{
		Method: "mining.authorize",
		Params: []string{username, password},
	}
	sp.send(&r)
	err := sp.receive(&resp)
	if err != nil {
		return "", err
	}
	/* TODO: Check error in response */

	err = sp.receive(&resp)
	if err != nil {
		return "", err
	}
	/* TODO: Check error in response */

	// Subscribe
	r = bitcoinRequest{
		Method: "mining.subscribe",
	}
	sp.send(&r)
	err = sp.receive(&resp)
	if err != nil {
		return "", err
	}
	/* TODO: Check error in response */

	err = sp.receive(&resp)
	if err != nil {
		return "", err
	}
	/* TODO: Check error in response */

	return "OK", nil
}

// bitcoinRequest sent from client to server
type bitcoinRequest struct {
	ID      int      `json:"id"`
	Method  string   `json:"method"`
	Version string   `json:"version,omitempty"`
	Params  []string `json:"params"`
}

// bitcoinResponse received from server
type bitcoinResponse struct {
	ID      int    `json:"id"`
	Version string `json:"jsonrpc"`
	Error   string `json:"error"`
}

// SetID set transaction ID
func (r *bitcoinRequest) setID(id int) {
	r.ID = id
}
