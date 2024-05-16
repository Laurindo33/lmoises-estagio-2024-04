package application

import (
	"Fintregas/domain"
	"testing"
)

func TestClienteSolicitarEntrega(t *testing.T) {
	t.Run("Cliente solicita uma entrega", func(t *testing.T) {
		// Dado
		/*cliente := domain.Cliente{
			Nome:     "João",
			Endereco: "Rua A, 123",
			Contacto: "joao@example.com",
		}*/
		encomenda := domain.Encomenda{
			EncomendaID: "10",
			Largura:     20,
			Altura:      23,
			Peso:        12,
		}

		//ClientEsperado := NovoCliente(&cliente)
		//EncomendaEsperada := NovaEncomenda(&encomenda)

		pedidoID, err := SolicitarEntrega(&encomenda)

		if err != nil {
			t.Errorf("Erro ao solicitar entrega: %v", err)
		}

		if pedidoID == "2" {
			t.Error("Identificador de pedido vazio")
		}
	})

}
