package model

type TarefaModel struct {
	Codigo    int64  `json:"codigo"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Completa  bool   `json:"completa"`
}
