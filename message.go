package anet

type Message struct {
	Api     string
	Payload interface{}
}

func NewMessage(api string, payload interface{}) *Message {
	return &Message{
		Api:     api,
		Payload: payload,
	}
}
