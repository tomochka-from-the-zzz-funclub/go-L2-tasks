package main

import "fmt"

type TVChannel interface {
	WriteAMessage()
}

type TV struct {
	TVChannel
}

func NewTV() *TV {
	return &TV{TV1{}}
}

func (tv *TV) SetChannel(tvchannel TVChannel) {
	tv.TVChannel = tvchannel
}

func (tv *TV) NextChannel() {
	switch tv.TVChannel.(type) {
	case TV1:
		tv.TVChannel = Russia1{}
	case Russia1:
		tv.TVChannel = Friday{}
	case Friday:
		tv.TVChannel = TV1{}
	}
}

type TV1 struct {
}

func (tv1 TV1) WriteAMessage() {
	fmt.Println("You are at TV1 channel now")
}

type Russia1 struct {
}

func (ru1 Russia1) WriteAMessage() {
	fmt.Println("You are at Russia 1 channel now")
}

type Friday struct {
}

func (fr Friday) WriteAMessage() {
	fmt.Println("You are at Friday channel now")
}

func main() {
	tv := NewTV()
	tv.SetChannel(Friday{})
	tv.WriteAMessage()
	tv.NextChannel()
	tv.WriteAMessage()
	tv.NextChannel()
	tv.WriteAMessage()
}
