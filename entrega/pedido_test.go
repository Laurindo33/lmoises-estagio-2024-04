package entrega

import (
	"testing"
)

func TestFormalizar(t *testing.T) {
	pedido := Pedido{
		dimensoes: 20,
		peso:      20,
		Endereco:  "Luanda",
		contatos:  "932838238",
	}

	service := NewPedidoServico()
	esperado := "Pedido formalizados com sucesso"
	resultado := service.FormalizeOrder(&pedido)
	if resultado != esperado {
		t.Errorf("FormalizPedido() retornou um resultado inesperado, pego: %s, esperado: %s", resultado, esperado)
	}
}
