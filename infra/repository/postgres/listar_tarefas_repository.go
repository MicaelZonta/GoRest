package postgres

import (
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres/connection"
	"GoRest/infra/repository/postgres/entity"
)

func ListarTarefasRepository(uC *loghelper.Usercontext) (tarefaSlice []entity.TarefaEntity, err error) {

	conn, err := connection.OpenConnection()
	if err != nil {
		loghelper.LogError(uC, loghelper.E00003, err)
		return tarefaSlice, err
	}
	defer conn.Close()

	//Query
	loghelper.LogDebug(uC, loghelper.D00001)
	rows, err := conn.Query("SELECT * FROM tarefas")
	if err != nil {
		loghelper.LogError(uC, loghelper.E00003, err)
		return tarefaSlice, err
	}

	for rows.Next() {
		var tE entity.TarefaEntity
		err = rows.Scan(&tE.Codigo, &tE.Titulo, &tE.Descricao, &tE.Completa)
		if err != nil {
			loghelper.LogError(uC, loghelper.E00003, err)
			continue
		}
		tarefaSlice = append(tarefaSlice, tE)
	}
	loghelper.LogDebug(uC, loghelper.D00001)

	//Return
	return tarefaSlice, err
}
