package database

type BaseEndPoint struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message "`
	Response interface{} `json:"response"`
}
