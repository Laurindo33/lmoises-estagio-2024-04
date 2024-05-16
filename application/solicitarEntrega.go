package application

import (
	"Fintregas/domain"
)

// Cliente representa um cliente que solicita uma entrega


// Encomenda representa uma encomenda a ser entregue

func NovoCliente(cliente *domain.Cliente) *domain.Cliente {
	return &domain.Cliente{
		Nome:     cliente.Nome,
		Endereco: cliente.Endereco,
		Contacto: cliente.Contacto,
	}
}

func NovaEncomenda(encomenda *domain.Encomenda) *domain.Encomenda {
	return &domain.Encomenda{
		Largura:     encomenda.Largura,
		Altura:       encomenda.Altura,
		Peso:         encomenda.Peso,
		EncomendaID:  encomenda.EncomendaID,
	}
}
func  SolicitarEntrega(encomenda *domain.Encomenda) (string, error) {

	//pedidoID := uuid.New().String()
	pedidoID:="1"

	// Neste exemplo, estamos simplesmente retornando o identificador do pedido
	return pedidoID, nil
}
