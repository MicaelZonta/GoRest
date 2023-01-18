package models

import (
	"GoRest/infra/repository/postgres/connection"
)

func Delete(codigo int64) (int64, error) {
	conn, err := connection.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	//Query
	res, err := conn.Exec("DELETE FROM tarefas WHERE codigo = $1", codigo)
	if err != nil {
		return 0, err
	}

	//Return
	return res.RowsAffected()
}
