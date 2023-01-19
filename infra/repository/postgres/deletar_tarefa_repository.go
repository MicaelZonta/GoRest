package postgres

import (
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres/connection"
)

func DeletarTarefaRepository(uC *loghelper.Usercontext, codigo int64) (int64, error) {

	conn, err := connection.OpenConnection()
	if err != nil {
		loghelper.LogError(uC, loghelper.E00006, err)
		return 0, err
	}
	defer conn.Close()

	//Query
	loghelper.LogDebug(uC, loghelper.D00001)
	res, err := conn.Exec("DELETE FROM tarefas WHERE codigo = $1", codigo)
	loghelper.LogDebug(uC, loghelper.D00001)

	if err != nil {
		loghelper.LogError(uC, loghelper.E00006, err)
		return 0, err
	}

	//Return
	return res.RowsAffected()
}
