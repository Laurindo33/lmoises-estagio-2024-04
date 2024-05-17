package application

import (
	"Fintregas/domain"
	"testing"
)

type entregaService struct{}

func (m *entregaService) AtualizarFila(entrega domain.Entrega) error {
	return nil
}

func (m *entregaService) NotificarDestinatario(entrega domain.Entrega) error {
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

	mockService := &entregaService{}
	service := PedidoService{entregaService: mockService}

	err := service.EntregarPedido(pedido)

	if err != nil {
		t.Errorf("Esperava-se sucesso, mas recebeu erro: %v", err)
	}
}

func TestEntregarPedido_NaoPago_Erro(t *testing.T) {
	pedido := domain.Pedido{
		Pago:        false,
		Formalizado: true,
	}

	mockService := &entregaService{}
	service := PedidoService{entregaService: mockService}

	err := service.EntregarPedido(pedido)

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

	mockService := &entregaService{}
	service := PedidoService{entregaService: mockService}

	err := service.EntregarPedido(pedido)

	if err == nil {
		t.Error("Esperava-se um erro, mas nenhum foi retornado")
	} else if err.Error() != "pedido não está pago ou formalizado" {
		t.Errorf("Erro esperado 'pedido não está pago ou formalizado', mas obteve: %v", err)
	}
}

func TestSolicitarPedido_PedidoFeitoEmLuanda(t *testing.T) {
	// Configuração do cenário de teste
	pedido := domain.Pedido{
		Morada: "Luanda",
	}

	pedidoService := PedidoService{}

	// Chamada da função a ser testada
	err := pedidoService.SolicitarPedido(pedido)

	// Verificação do resultado esperado
	if err != nil {
		t.Errorf("Esperava-se que não ocorresse nenhum erro, mas ocorreu o seguinte erro: %v", err)
	}
}

func TestSolicitarPedido_PedidoNaoFeitoEmLuanda(t *testing.T) {
	// Configuração do cenário de teste
	pedido := domain.Pedido{
		Morada: "Outro Lugar",
	}

	pedidoService := PedidoService{}

	// Chamada da função a ser testada
	err := pedidoService.SolicitarPedido(pedido)

	// Verificação do resultado esperado
	expectedError := "Pedido não está ser Feito em Luanda"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Esperava-se o erro: %v, mas recebeu: %v", expectedError, err)
	}
}
