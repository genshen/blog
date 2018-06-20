package utils

type SimpleJsonResponse struct {
	Status   int         `json:"status"`
	Error    interface{} `json:"error"`
	Addition interface{} `json:"addition"`
}
