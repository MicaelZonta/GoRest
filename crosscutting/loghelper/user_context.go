package loghelper

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type UserContext interface {
	Init(correlationId string) UserContext
	LogInfo(code string)
	LogDebug(code string)
	LogWarn(code string)
	LogError(code string, err error)
	LogFatal(code string, err error)
	LogPanic(code string, err error)
	SendValue(chave string, valor string)
}

type UserContextImpl struct {
	CorrelationId string
	TransactionId string
	UserValues    map[string]string
}

func (_this UserContextImpl) Init(correlationId string) UserContext {

	//Correlation
	if correlationId == "" {
		correlationId = uuid.New().String()
	}
	_this.CorrelationId = correlationId

	//Transaction
	_this.TransactionId = uuid.New().String()

	_this.UserValues = map[string]string{}

	return _this
}

func (_this UserContextImpl) SendValue(chave string, valor string) {
	_this.UserValues[chave] = valor
}

func (_this UserContextImpl) LogInfo(code string) {
	_this.logContent(log.Info(), code, "")
}

func (_this UserContextImpl) LogDebug(code string) {
	_this.logContent(log.Debug(), code, "")
}

func (_this UserContextImpl) LogWarn(code string) {
	_this.logContent(log.Warn(), code, "")
}

func (_this UserContextImpl) LogError(code string, err error) {
	_this.logContent(log.Error(), code, fmt.Sprintf("%v", err))
}

func (_this UserContextImpl) LogFatal(code string, err error) {
	_this.logContent(log.Fatal(), code, fmt.Sprintf("%v", err))
}

func (_this UserContextImpl) LogPanic(code string, err error) {
	_this.logContent(log.Panic(), code, fmt.Sprintf("%v", err))
}

func (_this UserContextImpl) logContent(zlog *zerolog.Event, code string, err string) {
	zlog = zlog.Str("correlation-id", _this.CorrelationId).
		Str("transaction-id", _this.TransactionId).
		Str("error", err)
	for chave, valor := range _this.UserValues {
		zlog = zlog.Str(chave, valor)
	}
	zlog.Msg(LogCodeMessage[code])
}
