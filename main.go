package main

import (
	"fmt"
	"github.com/hypersafe/simpleq/message"
)
func main(){
	
	msg := message.NewMsg()
	msg.From = "hello"
	msg.To ="world"
	msg.Content = "today"
	msg.GenerateCheckSum()	
	fmt.Println(msg.Checksum)
	fmt.Println(msg.ToJson())
}