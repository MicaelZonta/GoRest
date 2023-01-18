package loghelper

// string mapping
const (
	I00001 string = "I00001"
	I00002 string = "I00002"
	I00003 string = "I00003"
	I00004 string = "I00004"
	I00005 string = "I00005"
	I00006 string = "I00006"
	I00007 string = "I00007"
	I00008 string = "I00008"
	E00001 string = "E00001"
	E00002 string = "E00002"
	E00003 string = "E00003"
	E00004 string = "E00004"
	E00005 string = "E00005"
	W00001 string = "W00001"
)

var LogCodeMessage = map[string]string{
	I00001: "Buscando Tarefa - START",
	I00002: "Buscando Tarefa - END",
	I00003: "Criando Tarefa - USECASE - START",
	I00004: "Criando Tarefa - USECASE - END",
	I00005: "Conectando na Base Tarefa - Repository - START",
	I00006: "Conectando na Base Tarefa - Repository - END",
	I00007: "Criando Tarefa - GATEWAY - START",
	I00008: "Criando Tarefa - GATEWAY - END",
	E00001: "Erro ao buscar codigo da tarefa no Header.",
	E00002: "Erro ao consultar na base",
	E00003: "Erro ao conectar na base",
	W00001: "Busca sem resultado",
}
