package application

import (
	"Fintregas/domain"
	"fmt"
)

type PedidoService struct {
	entregaService domain.EntregaService
}

func (s *PedidoService) Entregar(pedido domain.Pedido) error {

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
