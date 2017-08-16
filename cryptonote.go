package msc

// Source: https://forum.minergate.com/viewtopic.php?f=9&t=1987

// CryptonoteStratumPool represent a Cryptonote Stratum Pool
type CryptonoteStratumPool struct {
	StratumPool // Inheritance
}

// CryptonoteWork represent a work assign to the worker
type CryptonoteWork struct {
	Blob   string `json:"blob"`
	Target string `json:"target"`
	JobID  string `json:"job_id"`
	TTL    int    `json:"time_to_live"`
}

// GetWork requests work from stratum pool
func (sp *CryptonoteStratumPool) GetWork(username, password string) (CryptonoteWork, error) {
	resp := cryptonoteResponse{}

	// Get job
	r := cryptonoteRequest{
		Method: "login",
		Params: map[string]string{
			"login": username,
			"pass":  password,
		},
	}
	sp.send(&r)
	err := sp.receive(&resp)
	if err != nil {
		return CryptonoteWork{}, err
	}
	/* TODO: Check error in response */

	// Get work
	return resp.Result.Job, nil
}

// cryptonoteRequest sent from client to server
type cryptonoteRequest struct {
	ID      int               `json:"id"`
	Method  string            `json:"method"`
	Version string            `json:"version,omitempty"`
	Params  map[string]string `json:"params"`
}

type cryptonoteResult struct {
	Job    CryptonoteWork `json:"job"`
	Status string         `json:"status"`
	ID     string         `json:"id"`
}

// cryptonoteResponse received from server
type cryptonoteResponse struct {
	ID      int              `json:"id"`
	Version string           `json:"jsonrpc"`
	Result  cryptonoteResult `json:"result,omitempty"`
}

// SetID set transaction ID
func (r *cryptonoteRequest) setID(id int) {
	r.ID = id
}
