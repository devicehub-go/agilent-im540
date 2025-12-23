package agilentim540

import (
	"github.com/devicehub-go/agilent-im540/protocol"
	"github.com/devicehub-go/unicomm"
)

type AgilentIM540 = protocol.AgilentIM540

/*
Creates a new instance of Agilent Ionivac IM540 tha allows to
communicate and control the connected pressure sensor
*/
func New(options unicomm.Options) *AgilentIM540 {
	options.Delimiter = protocol.CRLF
	return &AgilentIM540{
		Communication: unicomm.New(options),
	}
}
