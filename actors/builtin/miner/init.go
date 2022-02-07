package miner

import (
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rickiey/loggo"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Logoutput io.Writer
var db *gorm.DB
var err error

func init() {

	Logoutput = &lumberjack.Logger{
		Filename: "/var/log/lotus/specs-actors.log",

		// default 100
		MaxSize:    300, // megabytes
		MaxBackups: 60,
		MaxAge:     60,   //days
		Compress:   true, // disabled by default
	}

	loggo.SetLog(loggo.NewZapLog("debug", Logoutput))
	loggo.Info("loggo configuration done.")

	db, err = gorm.Open(sqlite.Open("/var/tmp/penalty_messages.db"), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(PenaltyMsg{})
	if err != nil {
		log.Println(err)
		return
	}
}

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

	pmsg := PenaltyMsg{
		ToAddr:       to,
		FromAddr:     from,
		Height:       height,
		Amount:       amount,
		CallFunction: callFunc,
		SubCause:     subcause,
		TimeAt:       time.Now(),
	}

	db.Create(pmsg)

	loggo.Info(pmsg)

	if JetNats == nil {
		return
	}

	b, err := json.Marshal(pmsg)
	if err != nil {
		loggo.Error("json.Marshal(pmsg) failed: ", err)
		return
	}

	puback, err := JetNats.Publish(NatsPenaltyMessageStream+"."+callFunc, b)
	if err != nil {
		loggo.Error(err)
		return
	}
	loggo.Info(puback, string(b))
}
