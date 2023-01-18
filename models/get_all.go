package models

import (
	"GoRest/core/model"
	"GoRest/infra/repository/postgres/connection"
)

func GetAll() (tarefaSlice []model.Tarefa, err error) {
	conn, err := connection.OpenConnection()
	if err != nil {
		return []model.Tarefa{}, err
	}
	defer conn.Close()

	//Query
	rows, err := conn.Query("SELECT * FROM tarefas")
	if err != nil {
		return []model.Tarefa{}, err
	}

	for rows.Next() {
		var t model.Tarefa
		err = rows.Scan(&t.Codigo, &t.Titulo, &t.Descricao, &t.Completa)
		if err != nil {
			continue
		}
		tarefaSlice = append(tarefaSlice, t)
	}

	//Return
	return tarefaSlice, err
}
