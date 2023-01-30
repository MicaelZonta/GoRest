package mock

import (
	"GoRest/crosscutting/loghelper"
	"GoRest/infra/repository/postgres"
	"GoRest/infra/repository/postgres/entity"
)

type TarefaRepositoryMock struct {
	CodigoRetorno             int64
	LinhasRetornadasDeletar   int64
	LinhasRetornadasAtualizar int64
	BuscarTarefaEntity        entity.TarefaEntity
	CriarTarefaEntity         entity.TarefaEntity
	ListarTarefaEntity        []entity.TarefaEntity
}

func (_this TarefaRepositoryMock) Init(uC loghelper.UserContext) postgres.TarefaRepository {
	return _this
}

func (_this TarefaRepositoryMock) InserirTarefaExecute(t entity.TarefaEntity) (entity.TarefaEntity, error) {
	return _this.CriarTarefaEntity, nil
}

func (_this TarefaRepositoryMock) ListarTarefasExecute() (tarefaSlice []entity.TarefaEntity, err error) {
	return _this.ListarTarefaEntity, nil
}

func (_this TarefaRepositoryMock) BuscarTarefaExecute(codigo int64) (t entity.TarefaEntity, err error) {
	_this.BuscarTarefaEntity.Codigo = codigo
	return _this.BuscarTarefaEntity, nil
}

func (_this TarefaRepositoryMock) DeletarTarefaExecute(codigo int64) (int64, error) {
	return _this.LinhasRetornadasDeletar, nil
}

func (_this TarefaRepositoryMock) AtualizarTarefaExecute(codigo int64, tE entity.TarefaEntity) (int64, error) {
	return _this.LinhasRetornadasAtualizar, nil
}
