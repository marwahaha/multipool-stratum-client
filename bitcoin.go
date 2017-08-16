package mf

// Source: https://forum.minergate.com/viewtopic.php?f=9&t=1987

// RequestBitcoin sent from client to server
type RequestBitcoin struct {
	ID      int      `json:"id"`
	Method  string   `json:"method"`
	Version string   `json:"version,omitempty"`
	Params  []string `json:"params"`
}

// ResponseBitcoin received from server
type ResponseBitcoin struct {
	ID      int    `json:"id"`
	Version string `json:"jsonrpc"`
}

// SetID set transaction ID
func (r *RequestBitcoin) SetID(id int) {
	r.ID = id
}

// BitcoinStratumPool represent a Bitcoin Stratum Pool
type BitcoinStratumPool struct {
	StratumPool // Inheritance
}

// GetWork requests work from stratum pool
func (sp *StratumPool) GetWork(username, password string) (string, error) {
	// Authorize
	r := RequestBitcoin{
		Method: "mining.authorize",
		Params: []string{username, password},
	}
	sp.send(&r)
	_, err := sp.receive()
	if err != nil {
		return "", err
	}
	_, err = sp.receive() /* TODO */

	// Subscribe
	r = RequestBitcoin{
		Method: "mining.subscribe",
	}
	sp.send(&r)
	_, err = sp.receive()
	if err != nil {
		return "", err
	}
	_, err = sp.receive() /* TODO */

	return "OK", nil
}
