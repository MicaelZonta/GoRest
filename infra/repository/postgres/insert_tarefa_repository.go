package postgres

import (
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres/connection"
	"GoRest/infra/repository/postgres/entity"
)

func InserirTarefaRepository(uC *loghelper.Usercontext, t entity.TarefaEntity) (entity.TarefaEntity, error) {

	conn, err := connection.OpenConnection()
	if err != nil {
		loghelper.LogError(uC, loghelper.E00003, err)
		return t, err
	}
	defer conn.Close()

	//Query
	loghelper.LogDebug(uC, loghelper.D00001)
	sql := "INSERT INTO tarefas(titulo, descricao, completa) VALUES ($1, $2, $3) returning codigo"
	err = conn.QueryRow(sql, t.Titulo, t.Descricao, t.Completa).Scan(&t.Codigo)
	loghelper.LogDebug(uC, loghelper.D00002)

	//Return
	return t, err
}
