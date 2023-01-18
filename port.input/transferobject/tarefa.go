package model

import "GoRest/infra/repository/postgres/entity"

type Tarefa struct {
	Codigo    int64  `json:"codigo"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Completa  bool   `json:"completa"`
}

func (tT *Tarefa) FromModel(tM entity.Tarefa) {
	tT.Codigo = tM.Codigo
	tT.Titulo = tM.Titulo
	tT.Descricao = tM.Titulo
	tT.Completa = tM.Completa
}
