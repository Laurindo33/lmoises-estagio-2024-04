package domain

import (
	"testing"
)

func TestRegistrarEntregaConcluida(t *testing.T) {
	t.Run("Registrar entrega concluída com sucesso", func(t *testing.T) {
		entrega := EntregaConcluida{}

		// Dados da entrega concluída
		encomendaID := "123"
		destinatario := "João"

		// Registrar entrega concluída
		entrega.RegistrarEntregaConcluida(encomendaID, destinatario)

		// Verificar se os dados foram corretamente atribuídos
		if entrega.encomendaID != encomendaID || entrega.destinatario != destinatario {
			t.Errorf("Os dados da entrega concluída não foram corretamente atribuídos: esperado %s, %s, recebido %s, %s", encomendaID, destinatario, entrega.encomendaID, entrega.destinatario)
		}
	})
}
