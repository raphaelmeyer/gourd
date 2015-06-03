package gourd

type wire_response interface {
	encode() string
}

type wire_response_success struct {
}

func (response *wire_response_success) encode() string {
	return ""
}
