package protocol

import (
	"fmt"
	"strings"

	"github.com/devicehub-go/unicomm"
)

const (
	CR   string = "\x0D"
	LF   string = "\x0A"
	CRLF string = CR + LF
	ETX  string = "\x03"
	ACK  string = "\x06"
	NAK  string = "\x15"
	ENQ  string = "\x05"
)

type AgilentIM540 struct {
	Communication unicomm.Unicomm
}

/*
Establishes a connection with the device
*/
func (a *AgilentIM540) Connect() error {
	return a.Communication.Connect()
}

/*
Closes the connection with the device
*/
func (a *AgilentIM540) Disconnect() error {
	return a.Communication.Disconnect()
}

/*
Checks if the device is connected
*/
func (a *AgilentIM540) IsConnected() bool {
	return a.Communication.IsConnected()
}

/*
Requests a message to the device and returns the response
*/
func (a *AgilentIM540) Request(message string) (string, error) {
	var invalidResponse string = "invalid"

	if !a.IsConnected() {
		return invalidResponse, fmt.Errorf("device is not connected")
	}
	if !strings.Contains(message, CRLF) {
		message += CRLF
	}
	if err := a.Communication.Write([]byte(message)); err != nil {
		return invalidResponse, err
	}
	ackResponse, err := a.Communication.ReadUntil(CRLF)
	if err != nil {
		return invalidResponse, err
	}
	if err := a.Communication.Write([]byte(ENQ)); err != nil {
		return invalidResponse, err
	}
	response, err := a.Communication.ReadUntil(CRLF)
	if err != nil {
		return invalidResponse, err
	}

	strAckResponse := string(ackResponse)
	if strings.Contains(strAckResponse, NAK) {
		return invalidResponse, fmt.Errorf("invalid request, got '%s'", strAckResponse)
	}

	return string(response[:len(response)-2]), nil
}
