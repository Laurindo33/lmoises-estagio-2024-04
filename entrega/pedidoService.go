package entrega

type PedidoService struct{}

func NewPedidoService() *PedidoService {
	return &PedidoService{}
}

func (os *PedidoService) FormalizarPedido(order *Pedido) string {
	return order.FormalizarPedido()
}
