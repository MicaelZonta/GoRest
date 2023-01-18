package loghelper

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Usercontext struct {
	CorrelationId string
	TransactionId string
	UserValues    map[string]string
}

func CreateUserContext(r *http.Request) (uc *Usercontext) {

	//Correlation
	correlationId := r.Context().Value("X-Correlation-Id").(string)
	if correlationId == "" {
		correlationId = uuid.New().String()
	}
	uc.CorrelationId = correlationId

	//Transaction
	uc.TransactionId = uuid.New().String()

	//Return UserContext
	return uc
}

func LogInfo(uC *Usercontext, code string) {
	logContent(log.Info(), uC, code, "")
}

func LogDebug(uC *Usercontext, code string) {
	logContent(log.Debug(), uC, code, "")
}

func LogWarn(uC *Usercontext, code string) {
	logContent(log.Warn(), uC, code, "")
}

func LogError(uC *Usercontext, code string, err error) {
	logContent(log.Error(), uC, code, fmt.Sprintf("%v", err))
}

func LogFatal(uC *Usercontext, code string, err error) {
	logContent(log.Fatal(), uC, code, fmt.Sprintf("%v", err))
}

func LogPanic(uC *Usercontext, code string, err error) {
	logContent(log.Panic(), uC, code, fmt.Sprintf("%v", err))
}

func logContent(zlog *zerolog.Event, uC *Usercontext, code string, err string) {
	zlog = zlog.Str("correlation-id", uC.CorrelationId).
		Str("transaction-id", uC.TransactionId).
		Str("error", err)
	for chave, valor := range uC.UserValues {
		zlog = zlog.Str(chave, valor)
	}
	zlog.Msg(LogCodeMessage[code])
}
