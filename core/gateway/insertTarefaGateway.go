package gateway

import (
	"GoRest/core/model"
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres"
	"GoRest/infra/repository/postgres/entity"
)

func SalvarTarefaGateway(uC *loghelper.Usercontext, t model.Tarefa) (model.Tarefa, error) {

	loghelper.LogDebug(uC, loghelper.I00007)
	tE := entity.Tarefa{}
	tE.FromModel(t)

	tE, err := postgres.InsertTarefaRepository(uC, tE)
	t.FromEntity(tE)
	loghelper.LogDebug(uC, loghelper.I00008)

	return t, err
}
