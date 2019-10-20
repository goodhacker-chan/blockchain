package baseBlockchain

import (
	"blockchain/tools"
	"bytes"
	"encoding/gob"
	"log"
)

// 事务输出
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

// 签证输出
func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := tools.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

// 检查公开秘钥的所有者是否可以输出
func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// 创建一个新的输出
func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

// 收藏事务输出
type TXOutputs struct {
	Outputs []TXOutput
}

// 序列化收藏输出
func (outs TXOutputs) Serialize() []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(outs)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// 反序列化收藏输出
func DeserializeOutputs(data []byte) TXOutputs {
	var outputs TXOutputs

	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}

	return outputs
}