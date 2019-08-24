package message
import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	)
type Msg struct {
	From 		string
	To 			string
	Content 	string
	Checksum	string
}



func NewMsg() Msg{
	return Msg{ }
}

func NewMsgFromJson(jsonstr string) (Msg, error){
	msg := Msg{}
	err := json.Unmarshal([]byte(jsonstr),&msg)
	return msg, err

}

func (msg *Msg)  GenerateCheckSum(){
	sumbuf := bytes.Buffer{}
	sumbuf.WriteString(msg.From)
	sumbuf.WriteString(msg.To)
	sumbuf.WriteString(msg.Content)
	
	hash := md5.Sum(sumbuf.Bytes())
	msg.Checksum = hex.EncodeToString(hash[:])

}


func (msg *Msg)  VerifyCheckSum() bool{
	if msg.Checksum=="" {
		return false
	}
	sumbuf := bytes.Buffer{}
	sumbuf.WriteString(msg.From)
	sumbuf.WriteString(msg.To)
	sumbuf.WriteString(msg.Content)
	
	hash := md5.Sum(sumbuf.Bytes())
	checksum := hex.EncodeToString(hash[:])
	if checksum == msg.Checksum{
		return true;
	}
	return false;


}

func (msg *Msg) ToJson() (string, error) {
	buf := bytes.Buffer{}
	buf.WriteString(msg.From)
	buf.WriteString(msg.To)
	buf.WriteString(msg.Content)
	
	jsonstr, err := json.Marshal(msg)
	if(err!=nil){
		log.Error("Unable to parse message : "+msg.From);
	}

	return string(jsonstr), err
	
	

}