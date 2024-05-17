package application

import (
	"Fintregas/domain"
	"fmt"
)

type PedidoService struct {
	entregaService domain.EntregaService
}
type EntregaService struct {
	entregaService domain.Entrega
}

func (s *PedidoService) SolicitarPedido(pedido domain.Pedido) error {

	if pedido.Morada != "Luanda" {
		return fmt.Errorf("Pedido não está ser Feito em Luanda")
	}
	return nil
}

func (s *PedidoService) EntregarPedido(pedido domain.Pedido) error {

	if !pedido.Pago || !pedido.Formalizado {
		return fmt.Errorf("pedido não está pago ou formalizado")
	}

	pedido.Entrega = domain.Entrega{
		ID:           "123",
		Status:       "Em andamento",
		Destinatario: pedido.Contacto,
	}

	return nil
}
