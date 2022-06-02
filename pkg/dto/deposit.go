package dto

import (
	"fmt"
	kafkaMessages "goka-test/proto/kafka"

	"google.golang.org/protobuf/proto"
)

type DepositCodec struct{}

func (d *DepositCodec) Encode(v interface{}) ([]byte, error) {
	message := v.(*kafkaMessages.Deposit)
	fmt.Println("Encode ", message)
	return proto.Marshal(message)
}

func (d *DepositCodec) Decode(data []byte) (interface{}, error) {

	var m kafkaMessages.Deposit
	err := proto.Unmarshal(data, &m)
	fmt.Println("decode ", m.WalletId)
	return &m, err
}

type AboveThresholdCodec struct{}

func (d *AboveThresholdCodec) Encode(v interface{}) ([]byte, error) {
	message := v.(*kafkaMessages.AboveThreshold)
	fmt.Println("Encode ", message)
	return proto.Marshal(message)
}

func (d *AboveThresholdCodec) Decode(data []byte) (interface{}, error) {

	var m kafkaMessages.AboveThreshold
	err := proto.Unmarshal(data, &m)
	fmt.Println("decode ", &m)
	return &m, err
}
