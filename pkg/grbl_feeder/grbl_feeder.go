package grblfeeder

import (
	"fmt"
	serialConnector "plotter/pkg/serial"
	"strings"
	"time"
)


type GrblFeeder struct {
	serialConnector *serialConnector.SerialConnection
}

func NewGrblFeeder(port string) *GrblFeeder {
	grblfeeder := new(GrblFeeder)
	grblfeeder.serialConnector = serialConnector.CreateSerialConnection(115200, port)

	fmt.Println("Waiting until grbl get ready...")
	grblfeeder.WaitUntilIdle()
	fmt.Println("GRBL ready")
	return grblfeeder
}

func (x GrblFeeder) WaitUntilIdle() {
	var isReady bool = false
	for !isReady {
		resp, _ := x.serialConnector.WriteAndWaitForResp("?\n")

		if len(resp) > 4 && string(resp[1:5]) == "Idle" {
			x.serialConnector.ReadLine()
			isReady = true
		}

		time.Sleep(time.Millisecond * 500)
	}
}

func (x GrblFeeder) SendGcode(gcode string) {
	fmt.Printf("%v\n", gcode)
	resp, err := x.serialConnector.WriteAndWaitForResp(gcode + "\n")
	if (err != nil) {
		fmt.Printf("Error [%v] happend while sending gcode", err)
	}
	if (strings.TrimRight(resp, "\r\n") != "ok") {
		fmt.Printf("Error [%v] happend while sending gcode", resp)
	}
}

func (x GrblFeeder) SendGcodes(gcode []string) {
	for _, line := range gcode {
		if (len(line) > 0 && line[:1] != ";") {
			x.SendGcode(line)
		}
	}
	x.WaitUntilIdle()
	fmt.Println("Draw complete")
}