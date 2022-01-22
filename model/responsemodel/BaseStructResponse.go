package responsemodel

//BaseStructRequest struct
type BaseStructRequest struct {
	Result     string      `json:"result"` //success warning error
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"` //200 423 424 500
	Data       interface{} `json:"data"`
}
