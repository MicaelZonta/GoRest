package postgres

import (
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres/connection"
	"GoRest/infra/repository/postgres/entity"
)

func BuscarTarefaRepository(uC *loghelper.Usercontext, codigo int64) (t entity.TarefaEntity, err error) {

	conn, err := connection.OpenConnection()
	if err != nil {
		loghelper.LogError(uC, loghelper.E00006, err)
		return t, err
	}
	defer conn.Close()

	//Query
	loghelper.LogDebug(uC, loghelper.D00007)
	row := conn.QueryRow("SELECT * FROM tarefas WHERE codigo=$1", codigo)
	err = row.Scan(&t.Codigo, &t.Titulo, &t.Descricao, &t.Completa)
	loghelper.LogDebug(uC, loghelper.D00008)

	//Return
	return t, err
}
