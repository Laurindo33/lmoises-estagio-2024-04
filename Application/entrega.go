package application

import "fintregas/model"

type PedidoEntrega struct{}

func (e *PedidoEntrega) FormalizarPedido(cliente *model.Cliente) bool {
	cliente.Encomenda.PagamentoFormalizado = true

	return cliente.Encomenda.PagamentoFormalizado

}

func (e *PedidoEntrega) ValidarPagamento(cliente *model.Cliente) bool {
	cliente.Encomenda.PagamentoFormalizado = true
	return cliente.Encomenda.PagamentoFormalizado
}

func (e *PedidoEntrega) VericarDestinatario(cliente *model.Cliente) string {
	cliente.Destinatario.Morada = "Luanda"

	return cliente.Destinatario.Morada
}
