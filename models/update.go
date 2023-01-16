package models

import "GoRest/db"

func Update(codigo int64, t Tarefa) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	//Query
	res, err := conn.Exec("UPDATE tarefas SET titulo = $1 , descricao = $2, completa = $3 WHERE codigo = $4", t.Titulo, t.Descricao, t.Completa, codigo)
	if err != nil {
		return 0, err
	}

	//Return
	return res.RowsAffected()
}
