package domain

import (
	"fmt"
)

// EntregaConcluida representa uma entrega concluída
type EntregaConcluida struct {
	encomendaID    string
	destinatario   string
}

// RegistrarEntregaConcluida registra uma entrega concluída
func (e *EntregaConcluida) RegistrarEntregaConcluida(encomendaID, destinatario string) {
	e.encomendaID = encomendaID
	e.destinatario = destinatario

	fmt.Println("Entrega concluída registrada com sucesso!")
}
