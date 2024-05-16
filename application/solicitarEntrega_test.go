package application

import (
	"Fintregas/domain"
	"testing"
)

func TestClienteSolicitarEntrega(t *testing.T) {
	t.Run("Cliente solicita uma entrega", func(t *testing.T) {
		// Dado
		
		encomenda := domain.Encomenda{
			EncomendaID:"10",
			Largura:     20,
			Altura:      23,
			Peso:        12,
			MoradaDestinatario: "Bengo",
		}

		Esperado :=SolicitarEntrega(encomenda)
		Actual:="Benguela"



		if Esperado != Actual {
			t.Errorf("Erro ao solicitar entrega: %v Diferente de %v",Esperado, Actual )
		}

		
	})

}
