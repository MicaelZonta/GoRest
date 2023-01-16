package models

import "GoRest/db"

func Get(codigo int64) (t Tarefa, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return Tarefa{}, err
	}
	defer conn.Close()

	//Query
	row := conn.QueryRow("SELECT * FROM tarefas WHERE codigo=$1", codigo)
	err = row.Scan(&t.Codigo, &t.Titulo, &t.Descricao, &t.Completa)

	//Return
	return t, err
}
