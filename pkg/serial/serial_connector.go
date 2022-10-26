package serialConnector

import (
	"fmt"
	"go.bug.st/serial"
	"log"
)

type SerialConnection struct {
	baudRate int
	portStr string
	port serial.Port
	buff []byte
}

func CreateSerialConnection(baudRate int, portStr string) *SerialConnection {

	mode := &serial.Mode{
		BaudRate: baudRate,
	}

	// printPorts()

	port, err := serial.Open(portStr, mode)
	if err != nil {
		fmt.Println(err)
		panic("Unable to open serial port")
	}

	sc := new(SerialConnection)
	sc.baudRate = baudRate
	sc.portStr = portStr
	sc.port = port
	sc.buff = make([]byte, 500)
	
	fmt.Printf("Serial connection open [%v]\n", portStr)
	
	return sc
}

func (x SerialConnection) Write(command string) {
	_, err := x.port.Write([]byte(command))
	if err != nil {
		log.Fatal(err)
	}
}

func (x SerialConnection) WriteAndWaitForResp(command string) (string, error) {
	x.Write(command)
	// fmt.Printf("Sent %v", command)
	return x.ReadLine()
}

func (x SerialConnection) ReadLine() (string, error) {
	n, err := x.port.Read(x.buff)
	return string(x.buff[:n]), err
}

func printPorts() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}
