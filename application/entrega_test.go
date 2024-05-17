package application

import (
	"Fintregas/domain"
	"testing"
)

type mockEntregaService struct{}

func (m *mockEntregaService) AtualizarFila(entrega domain.Entrega) error {
	return nil
}

func (m *mockEntregaService) NotificarDestinatario(entrega domain.Entrega) error {
	return nil
}

func TestEntregarPedido_PagoEFormalizado_Sucesso(t *testing.T) {
	pedido := domain.Pedido{
		Pago:        true,
		Formalizado: true,
		Entrega: domain.Entrega{
			ID:           "123",
			Status:       "Aguardando",
			Destinatario: "John Doe",
		},
	}

	mockService := &mockEntregaService{}
	service := PedidoService{entregaService: mockService}

	err := service.Entregar(pedido)

	if err != nil {
		t.Errorf("Esperava-se sucesso, mas recebeu erro: %v", err)
	}
}

func TestEntregarPedido_NaoPago_Erro(t *testing.T) {
	pedido := domain.Pedido{
		Pago:        false,
		Formalizado: true,
	}

	mockService := &mockEntregaService{}
	service := PedidoService{entregaService: mockService}

	err := service.Entregar(pedido)

	if err == nil {
		t.Error("Esperava-se um erro, mas nenhum foi retornado")
	} else if err.Error() != "pedido não está pago ou formalizado" {
		t.Errorf("Erro esperado 'pedido não está pago ou formalizado', mas obteve: %v", err)
	}
}

func TestEntregarPedido_NaoFormalizado_Erro(t *testing.T) {
	pedido := domain.Pedido{
		Pago:        true,
		Formalizado: false,
	}

	mockService := &mockEntregaService{}
	service := PedidoService{entregaService: mockService}

	err := service.Entregar(pedido)

	if err == nil {
		t.Error("Esperava-se um erro, mas nenhum foi retornado")
	} else if err.Error() != "pedido não está pago ou formalizado" {
		t.Errorf("Erro esperado 'pedido não está pago ou formalizado', mas obteve: %v", err)
	}
}

func TestEntregar_PedidoNaoPago(t *testing.T) {
	// Configuração do cenário de teste
	pedido := domain.Pedido{
		Pago:        false,
		Formalizado: true,
	}
	pedidoService := PedidoService{&mockEntregaService{}}

	// Chamada da função a ser testada
	err := pedidoService.Entregar(pedido)

	// Verificação do resultado esperado
	expectedError := "pedido não está pago ou formalizado"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Esperava-se o erro: %v, mas recebeu: %v", expectedError, err)
	}
}
