package entrega

type PedidoService struct{}

func NewPedidoServico() *PedidoService {
	return &PedidoService{}
}

func (os *PedidoService) FormalizeOrder(order *Pedido) string {
	return order.FormalizarPedido()
}
