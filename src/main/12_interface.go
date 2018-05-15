package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connect()
	Disconnect()
}

type TV interface {
	Name
}

type Name interface {
	Name() string
}

type PhoneConnector struct {
	name string
}

type TVConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	fmt.Println("Name() called: ", pc.name)
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connect:", pc.name)
}

func (pc PhoneConnector) Disconnect() {
	fmt.Println("Disconnect:", pc.name)
}

func (tv TVConnector) Name() string {
	fmt.Println("Name() called: ", tv.name)
	return tv.name
}

func TurnOff(tv TV) {
	fmt.Println(tv.Name(), "turn off")
}

func GetType(bb interface{}) {
	switch v := bb.(type) {
	case PhoneConnector:
		fmt.Println(v.name, " is PhoneConnector")
	default:
		fmt.Println("Unknown device")
	}
}

func interface_usb() {
	var usb USB
	usb = PhoneConnector{"iPhone"}
	usb.Name()
	usb.Connect()
	usb.Disconnect()
	GetType(usb)
	TurnOff(usb)

	var tv TV
	tv = TVConnector{"mi tv"}
	tv.Name()
	GetType(tv)
	TurnOff(tv)
}
