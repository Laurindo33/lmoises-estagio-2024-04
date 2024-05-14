package application

import (
	"fintregas/model"
	"testing"
)

func TestEntregarPedido(t *testing.T) {

	t.Run("Teste Validar Pedido de Encomenda", func(t *testing.T) {

		//Arrange
		esperado := true
		pedido := PedidoEntrega{}
		actual := pedido.FormalizarPedido(&model.Cliente{})

		//Act

		//Assert
		if esperado != actual {
			t.Errorf("%t != %t", esperado, actual)
			t.Fail()
		}
	})

	t.Run("Teste  Validar Pagamento Pagamento da encomenda ", func(t *testing.T) {

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

	t.Run("Teste Verificar Morada válida", func(t *testing.T) {
		//Arrange
		esperado := "Luanda"
		pedido := PedidoEntrega{}
		actual := pedido.VericarDestinatario(&model.Cliente{})
		//Act

		//Assert
		if esperado != actual {
			t.Errorf("%s != %s", esperado, actual)
			t.Fail()
		}
	})
}
