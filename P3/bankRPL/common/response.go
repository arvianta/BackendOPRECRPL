package common

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"error"`
	Data    any    `json:"data"`
}
