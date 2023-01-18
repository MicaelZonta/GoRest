package model

import "GoRest/infra/repository/postgres/entity"

type Tarefa struct {
	Codigo    int64  `json:"codigo"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Completa  bool   `json:"completa"`
}

func (tM *Tarefa) FromEntity(tE entity.Tarefa) {
	tM.Codigo = tE.Codigo
	tM.Titulo = tE.Titulo
	tM.Descricao = tE.Titulo
	tM.Completa = tE.Completa
}
