package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

type Camera struct {
	name string
}

func (p Phone) Start(){
	fmt.Printf("%v开始工作\n", p.name)
}
func (c Camera) Start(){
	fmt.Printf("%v开始工作\n", c.name)
}
func (p Phone) Stop(){
	fmt.Printf("%v停止工作\n", p.name)
}
func (c Camera) Stop(){
	fmt.Printf("%v停止工作\n", c.name)
}
func (p Phone) Call(){
	fmt.Printf("%v拨打电话\n", p.name)
}

type Computer struct {
}
func (com Computer) working(usb Usb){
	usb.Start()
	usb.Stop()
	if p, ok := usb.(Phone); ok{
		p.Call()
	}
}
func main() {
	var usblince [3]Usb
	usblince[0] = Phone{
		name : "小米",
	}
	usblince[1] = Camera{
		name : "佳能",
	}
	usblince[2] = Phone{
		name : "Apple",
	}
	var computer Computer
	for _, value := range usblince{
		computer.working(value)
	}
}