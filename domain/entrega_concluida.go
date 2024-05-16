package domain

import (
	"fmt"
)

// EntregaConcluida representa uma entrega concluída
type EntregaConcluida struct {
	encomendaID    string
	destinatario   string
	dataConclusao  string
}

// RegistrarEntregaConcluida registra uma entrega concluída
func (e *EntregaConcluida) RegistrarEntregaConcluida(encomendaID, destinatario, dataConclusao string) {
	e.encomendaID = encomendaID
	e.destinatario = destinatario
	e.dataConclusao = dataConclusao

	fmt.Println("Entrega concluída registrada com sucesso!")
}
