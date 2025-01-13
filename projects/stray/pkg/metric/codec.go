package metric

import "github.com/munnerz/goautoneg"

// A Codec performs encoding of API responses.
type Codec interface {
	// ContentType returns the MIME time that this Codec emits.
	ContentType() MIMEType

	// CanEncode determines if this Codec can encode resp.
	CanEncode(resp *Response) bool

	// Encode encodes resp, ready for transmission to an API consumer.
	Encode(resp *Response) ([]byte, error)
}

type MIMEType struct {
	Type    string
	SubType string
}

func (m MIMEType) String() string {
	return m.Type + "/" + m.SubType
}

func (m MIMEType) Satisfies(accept goautoneg.Accept) bool {
	if accept.Type == "*" && accept.SubType == "*" {
		return true
	}

	if accept.Type == m.Type && accept.SubType == "*" {
		return true
	}

	if accept.Type == m.Type && accept.SubType == m.SubType {
		return true
	}

	return false
}
