package controllers

type Msg struct {
	Msg string
}

func NewMsg(msg string) Msg {
	return Msg {
		Msg: msg,
	}
}