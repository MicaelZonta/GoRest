package entity

import "GoRest/core/model"

type Tarefa struct {
	Codigo    int64
	Titulo    string
	Descricao string
	Completa  bool
}

func (tE *Tarefa) FromModel(tM model.Tarefa) {
	tE.Codigo = tM.Codigo
	tE.Titulo = tM.Titulo
	tE.Descricao = tM.Descricao
	tE.Completa = tM.Completa
}
