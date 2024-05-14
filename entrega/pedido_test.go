package entrega

import (
	"testing"
)

func TestFormalizarPedido(t *testing.T) {

	//Assert

	pedido := Pedido{
		dimensoes: []int{1,2},
		peso:      20,
		Endereco:  "Luanda",
		contatos:  "932838238",
	}

	esperado := "Pedido formalizado com sucesso"

	//Act
	
	service := NewPedidoService()
	resultado := service.FormalizarPedido(&pedido)

	//Assert

	if resultado != esperado {
		t.Errorf(" pego: %s, esperado: %s", resultado, esperado)
	}
}
