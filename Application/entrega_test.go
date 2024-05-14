package application

import (
	"fintregas/model"
	"testing"
)

func TestEntregarPedido(t *testing.T) {

	t.Run("Teste  pedido Formalizao?", func(t *testing.T) {

		//Arrange
		esperado := false
		pedido := PedidoEntrega{}
		actual := pedido.FormalizarPedido(&model.Cliente{})

		//Act

		//Assert
		if esperado != actual {
			t.Errorf("%t != %t", esperado, actual)
			t.Fail()
		}
	})

	t.Run("Teste  Pagamento Formalizao?", func(t *testing.T) {

		//Arrange
		esperado := true
		pedido := PedidoEntrega{}
		actual := pedido.ValidarPagamento(&model.Cliente{})
		//Act

		//Assert
		if esperado != actual {
			t.Errorf("%t != %t", esperado, actual)
			t.Fail()
		}

	})
}
