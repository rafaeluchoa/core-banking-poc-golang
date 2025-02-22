package api

// @Description Response
// @Param code body string true "Code"
// @Param message body string true "Message"
type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
