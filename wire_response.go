package gourd

type wire_response interface {
	encode() string
}

type generic_wire_response struct{}

func (response *generic_wire_response) encode() string {
	return ""
}

type wire_response_success struct {
}

func (response *wire_response_success) encode() string {
	return ""
}
