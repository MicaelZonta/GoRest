package models

import (
	"GoRest/core/model"
	"GoRest/infra/repository/postgres/connection"
)

func Get(codigo int64) (t model.Tarefa, err error) {
	conn, err := connection.OpenConnection()
	if err != nil {
		return model.Tarefa{}, err
	}
	defer conn.Close()

	//Query
	row := conn.QueryRow("SELECT * FROM tarefas WHERE codigo=$1", codigo)
	err = row.Scan(&t.Codigo, &t.Titulo, &t.Descricao, &t.Completa)

	//Return
	return t, err
}
