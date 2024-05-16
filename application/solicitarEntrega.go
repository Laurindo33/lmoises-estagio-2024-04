package application

import (
	"Fintregas/domain"
)

// Entrega representa o processo de entrega de uma encomenda
type Entrega struct{}

// SolicitarEntrega solicita uma entrega para a encomenda especificada
func SolicitarEntrega(encomenda domain.Encomenda) string {
	var retorno string

	if encomenda.MoradaDestinatario != "Luanda" && encomenda.ContactoDestinatario == "" {
		retorno = "Solicitação Negada"
	} else {
		retorno = "Solicitação Aceite"
	}
	return retorno
}
func (f *FilaDeEntregas) AdicionarEntrega(entrega EntregaConcluida) {
	f.entregas = append(f.entregas, entrega)
}
