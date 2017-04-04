package controllers

type Send struct {
	Message string
	Data    interface{}
}

func sendMessage(message string, data interface{}) (send Send) {
	send.Message = message
	send.Data = data
	return
}
