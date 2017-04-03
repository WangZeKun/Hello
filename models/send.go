package models

type Send struct {
	message string
	data    interface{}
}

func SendMessage(message string, data interface{}) (send Send) {
	send.message = message
	send.data = data
	return
}
