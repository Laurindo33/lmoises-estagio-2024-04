package application

import "fintregas/model"

/*import (
	"fintregas/model"
)
*/

type PedidoEntrega struct{}

func (e *PedidoEntrega) FormalizarPedido(cliente *model.Cliente) bool {
	cliente.Encomenda.PagamentoFormalizado = true 

	return cliente.Encomenda.PagamentoFormalizado

}

//ordemEntrgaFormalizada

func (e *PedidoEntrega) ValidarPagamento(cliente *model.Cliente) bool {
	cliente.Encomenda.PagamentoFormalizado=true
	return cliente.Encomenda.PagamentoFormalizado
}
