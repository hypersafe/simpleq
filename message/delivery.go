package message

import (
	// log "github.com/sirupsen/logrus"
	"errors"
)

type DeliveryProp struct{
	TempDir string
	RecipientDir string
}


func NewDeliveryProp( tempdir string, recipientdir string) DeliveryProp{
	return DeliveryProp{TempDir: tempdir, RecipientDir: recipientdir}
}

func( delivery *DeliveryProp) deliver(msg Msg) error{

	
	
	// if msg.CheckValidMsg(){
	// 	msgjson, jerr := msg.ToJson()
	// 	if jerr!= nil{
	// 		log.Error("unable to parse message, rejecting delivery")
	// 		return jerr
	// 	}

	// }

	// tmp directory
	// moves file to folder
	
	return errors.New("todo")
}