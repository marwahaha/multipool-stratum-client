package mf

// Source: https://forum.minergate.com/viewtopic.php?f=9&t=1987

// RequestCryptonote sent from client to server
type RequestCryptonote struct {
	ID      int               `json:"id"`
	Method  string            `json:"method"`
	Version string            `json:"version,omitempty"`
	Params  map[string]string `json:"params"`
}

// ResponseCryptonote received from server
type ResponseCryptonote struct {
	ID      int    `json:"id"`
	Version string `json:"jsonrpc"`
}

// SetID set transaction ID
func (r *RequestCryptonote) SetID(id int) {
	r.ID = id
}

// CryptonoteStratumPool represent a Cryptonote Stratum Pool
type CryptonoteStratumPool struct {
	StratumPool // Inheritance
}

// GetWork requests work from stratum pool
func (sp *CryptonoteStratumPool) GetWork(username, password string) (string, error) {
	r := RequestCryptonote{
		Method: "login",
		Params: map[string]string{
			"login": username,
			"pass":  password,
		},
	}
	sp.send(&r)
	_, err := sp.receive()
	if err != nil {
		return "", err
	}
	return "OK", nil
}
