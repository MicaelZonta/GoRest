package postgres

import (
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres/connection"
	"GoRest/infra/repository/postgres/entity"
)

func AtualizarTarefaRepository(uC *loghelper.Usercontext, codigo int64, tE entity.TarefaEntity) (int64, error) {

	conn, err := connection.OpenConnection()
	if err != nil {
		loghelper.LogError(uC, loghelper.E00003, err)
		return 0, err
	}
	defer conn.Close()

	//Query
	loghelper.LogDebug(uC, loghelper.D00007)
	res, err := conn.Exec("UPDATE tarefas SET titulo = $1 , descricao = $2, completa = $3 WHERE codigo = $4", tE.Titulo, tE.Descricao, tE.Completa, codigo)
	loghelper.LogDebug(uC, loghelper.D00007)

	if err != nil {
		loghelper.LogError(uC, loghelper.E00003, err)
		return 0, err
	}

	//Return
	return res.RowsAffected()
}
