package loghelper

import (
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

func CreateUserContext(r *http.Request) (uc Usercontext) {

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

func LogInfo(us Usercontext, code string, msg string) {
	logContent(log.Info(), us, code, msg)
}

func LogDebug(us Usercontext, code string, msg string) {
	logContent(log.Debug(), us, code, msg)
}

func LogWarn(us Usercontext, code string, msg string) {
	logContent(log.Warn(), us, code, msg)
}

func LogError(us Usercontext, code string, msg string) {
	logContent(log.Error(), us, code, msg)
}

func LogFatal(us Usercontext, code string, msg string) {
	logContent(log.Fatal(), us, code, msg)
}

func LogPanic(us Usercontext, code string, msg string) {
	logContent(log.Panic(), us, code, msg)
}

func logContent(zlog *zerolog.Event, us Usercontext, code string, msg string) {
	zlog = zlog.Str("correlation-id", us.CorrelationId).
		Str("transaction-id", us.TransactionId)

	for chave, valor := range us.UserValues {
		zlog = zlog.Str(chave, valor)
	}
	zlog.Msg(msg)
}
