package message


import (
	"testing"
	"gotest.tools/assert"
)
func TestCheckSum(t *testing.T) {
	// t.Error() // to indicate test failed
	msg := NewMsg()
	msg.From = "hello"
	msg.To ="world"
	msg.Content = "today"
	msg.GenerateCheckSum()
	assert.Equal(t,msg.Checksum, "11857333f874c5b93476e9baca1bc4e9")
	assert.Equal(t,msg.VerifyCheckSum(), true)
}

func TestMarshal(t *testing.T){
	msg, err := NewMsgFromJson(`{"From":"hello","To":"world","Content":"today","Checksum":"11857333f874c5b93476e9baca1bc4e9"}`)
	assert.Equal(t, err, nil)
	assert.Equal(t, msg.From,"hello")
	assert.Equal(t, msg.To,"world")
	assert.Equal(t, msg.Content,"today")
	assert.Equal(t, msg.Checksum,"11857333f874c5b93476e9baca1bc4e9")
	
}