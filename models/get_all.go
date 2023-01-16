package models

import "GoRest/db"

func GetAll() (tarefaSlice []Tarefa, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return []Tarefa{}, err
	}
	defer conn.Close()

	//Query
	rows, err := conn.Query("SELECT * FROM tarefas")
	if err != nil {
		return []Tarefa{}, err
	}

	for rows.Next() {
		var t Tarefa
		err = rows.Scan(&t.Codigo, &t.Titulo, &t.Descricao, &t.Completa)
		if err != nil {
			continue
		}
		tarefaSlice = append(tarefaSlice, t)
	}

	//Return
	return tarefaSlice, err
}
