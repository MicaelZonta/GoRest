package postgres

import (
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres/connection"
	"GoRest/infra/repository/postgres/entity"
)

type TarefaRepository interface {
	Init(uC loghelper.UserContext) TarefaRepository
	InserirTarefaExecute(t entity.TarefaEntity) (entity.TarefaEntity, error)
	ListarTarefasExecute() (tarefaSlice []entity.TarefaEntity, err error)
	BuscarTarefaExecute(codigo int64) (t entity.TarefaEntity, err error)
	DeletarTarefaExecute(codigo int64) (int64, error)
	AtualizarTarefaExecute(codigo int64, tE entity.TarefaEntity) (int64, error)
}

type TarefaRepositoryImpl struct {
	uC loghelper.UserContext
}

func (_this TarefaRepositoryImpl) Init(uC loghelper.UserContext) TarefaRepository {
	_this.uC = uC
	return _this
}

func (_this TarefaRepositoryImpl) InserirTarefaExecute(t entity.TarefaEntity) (entity.TarefaEntity, error) {

	conn, err := connection.OpenConnection()
	if err != nil {
		_this.uC.LogError(loghelper.E00003, err)
		return t, err
	}
	defer conn.Close()

	//Query
	_this.uC.LogDebug(loghelper.D00001)
	sql := "INSERT INTO tarefas(titulo, descricao, completa) VALUES ($1, $2, $3) returning codigo"
	err = conn.QueryRow(sql, t.Titulo, t.Descricao, t.Completa).Scan(&t.Codigo)
	_this.uC.LogDebug(loghelper.D00002)

	//Return
	return t, err
}

func (_this TarefaRepositoryImpl) ListarTarefasExecute() (tarefaSlice []entity.TarefaEntity, err error) {
	_this.uC.SendValue("4", "4")
	conn, err := connection.OpenConnection()
	if err != nil {
		_this.uC.LogError(loghelper.E00003, err)
		return tarefaSlice, err
	}
	defer conn.Close()

	//Query
	_this.uC.LogDebug(loghelper.D00001)
	rows, err := conn.Query("SELECT * FROM tarefas")
	if err != nil {
		_this.uC.LogError(loghelper.E00003, err)
		return tarefaSlice, err
	}

	for rows.Next() {
		var tE entity.TarefaEntity
		err = rows.Scan(&tE.Codigo, &tE.Titulo, &tE.Descricao, &tE.Completa)
		if err != nil {
			_this.uC.LogError(loghelper.E00003, err)
			continue
		}
		tarefaSlice = append(tarefaSlice, tE)
	}
	_this.uC.LogDebug(loghelper.D00001)

	//Return
	return tarefaSlice, err
}

func (_this TarefaRepositoryImpl) BuscarTarefaExecute(codigo int64) (t entity.TarefaEntity, err error) {

	conn, err := connection.OpenConnection()
	if err != nil {
		_this.uC.LogError(loghelper.E00006, err)
		return t, err
	}
	defer conn.Close()

	//Query
	_this.uC.LogDebug(loghelper.D00007)
	row := conn.QueryRow("SELECT * FROM tarefas WHERE codigo=$1", codigo)
	err = row.Scan(&t.Codigo, &t.Titulo, &t.Descricao, &t.Completa)
	_this.uC.LogDebug(loghelper.D00008)

	//Return
	return t, err
}

func (_this TarefaRepositoryImpl) DeletarTarefaExecute(codigo int64) (int64, error) {

	conn, err := connection.OpenConnection()
	if err != nil {
		_this.uC.LogError(loghelper.E00006, err)
		return 0, err
	}
	defer conn.Close()

	//Query
	_this.uC.LogDebug(loghelper.D00001)
	res, err := conn.Exec("DELETE FROM tarefas WHERE codigo = $1", codigo)
	_this.uC.LogDebug(loghelper.D00001)

	if err != nil {
		_this.uC.LogError(loghelper.E00006, err)
		return 0, err
	}

	//Return
	return res.RowsAffected()
}

func (_this TarefaRepositoryImpl) AtualizarTarefaExecute(codigo int64, tE entity.TarefaEntity) (int64, error) {

	conn, err := connection.OpenConnection()
	if err != nil {
		_this.uC.LogError(loghelper.E00003, err)
		return 0, err
	}
	defer conn.Close()

	//Query
	_this.uC.LogDebug(loghelper.D00007)
	res, err := conn.Exec("UPDATE tarefas SET titulo = $1 , descricao = $2, completa = $3 WHERE codigo = $4", tE.Titulo, tE.Descricao, tE.Completa, codigo)
	_this.uC.LogDebug(loghelper.D00007)

	if err != nil {
		_this.uC.LogError(loghelper.E00003, err)
		return 0, err
	}

	//Return
	return res.RowsAffected()
}
