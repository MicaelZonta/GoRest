package models

import "GoRest/db"

func Insert(t Tarefa) (codigo int64, err error) {
	//Open
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	//Query
	sql := "INSERT INTO tarefas(titulo, descricao, completa) VALUES ($1, $2, $3) returning codigo"
	err = conn.QueryRow(sql, t.Titulo, t.Descricao, t.Completa).Scan(&codigo)

	//Return
	return codigo, err
}
