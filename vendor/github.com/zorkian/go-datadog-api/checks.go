package datadog

type Check struct {
	Check     string   `json:"check"`
	HostName  string   `json:"host_name"`
	Status    Status   `json:"status"`
	Timestamp string   `json:"timestamp,omitempty"`
	Message   string   `json:"message,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

type Status int

const (
	OK Status = iota
	WARNING
	CRITICAL
	UNKNOWN
)

// PostCheck posts the result of a check run to the server
func (client *Client) PostCheck(check Check) error {
	return client.doJsonRequest("POST", "/v1/check_run",
		check, nil)
}
