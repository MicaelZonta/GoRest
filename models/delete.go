package models

import "GoRest/db"

func Delete(codigo int64) (int64, error) {
	conn, err := db.OpenConnection()
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
