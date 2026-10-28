package agilentim540_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	agilentim540 "github.com/devicehub-go/agilent-im540"
	"github.com/devicehub-go/unicomm"
	"github.com/devicehub-go/unicomm/protocol/unicommtcp"
)

func TestMain(t *testing.T) {
	sensor := agilentim540.New(unicomm.Options{
		Protocol: unicomm.TCP,
		TCP: unicommtcp.TCPOptions{
			Host:         "10.0.3.35",
			Port:         4003,
			ReadTimeout:  1000 * time.Millisecond,
			WriteTimeout: 1000 * time.Millisecond,
		},
	})
	if err := sensor.Connect(); err != nil {
		log.Fatal(err)
	}
	defer sensor.Disconnect()

	fmt.Println(sensor.GetEmission())
}
