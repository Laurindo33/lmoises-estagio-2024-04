package application

import "fintregas/model"

type PedidoEntrega struct{}

func (e *PedidoEntrega) FormalizarPedido(cliente *model.Cliente) bool {
	cliente.Encomenda.PagamentoFormalizado = true

	return true

}

func (e *PedidoEntrega) ValidarPagamento(cliente *model.Cliente) bool {
	cliente.Encomenda.PagamentoFormalizado = true
	return true
}

func (e *PedidoEntrega) VericarDestinatario(cliente *model.Cliente) string {
	
	cliente.Destinatario.Morada = "Luanda"
	morada:=cliente.Destinatario.Morada

	return morada
}
