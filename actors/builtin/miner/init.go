package miner

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/rickiey/loggo"
	"time"
)

var JetNats nats.JetStreamContext

type PenaltyMsg struct {
	ToAddr       string    `json:"to_addr"`
	FromAddr     string    `json:"from_addr"`
	Height       int64     `json:"height"`
	Amount       string    `json:"amount"`
	TimeAt       time.Time `json:"time_at"`
	CallFunction string    `json:"call_function"` // 调用方法
	SubCause     string    `json:"sub_cause"`
}

const NatsPenaltyMessageStream = "lotus_penalty_msg"

func PubPenaltyMsg(to, from string, height int64, amount, callFunc, subcause string) {

	if JetNats == nil {
		return
	}

	pmsg := PenaltyMsg{
		ToAddr:       to,
		FromAddr:     from,
		Height:       height,
		Amount:       amount,
		CallFunction: callFunc,
		SubCause:     subcause,
		TimeAt:       time.Now(),
	}

	b, _ := json.Marshal(pmsg)

	puback, err := JetNats.Publish(NatsPenaltyMessageStream+"."+callFunc, b)
	if err != nil {
		loggo.Error()
		return
	}
	loggo.Info(puback, string(b))
}
