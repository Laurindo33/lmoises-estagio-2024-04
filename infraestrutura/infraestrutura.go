package infraestrutura

// Camada de Infraestrutura

import (
	"Fintregas/domain"
	"fmt"
)

// FakeEntregaService é uma implementação falsa do serviço de entrega para testes
type FakeEntregaService struct{}

// CalcularRota simula o cálculo de rota
func (s *FakeEntregaService) CalcularRota(pedido domain.Pedido) (string, error) {
	return "Rota calculada", nil
}

// NotificarDestinatario simula a notificação do destinatário
func (s *FakeEntregaService) NotificarDestinatario(entrega domain.Entrega) error {
	fmt.Println("Destinatário notificado:", entrega.Destinatario)
	return nil
}

// AtualizarFila simula a atualização da fila
func (s *FakeEntregaService) AtualizarFila(entrega domain.Entrega) error {
	fmt.Println("Fila atualizada para:", entrega.Destinatario)
	return nil
}
