package Rm_localizacao

import (
	"testing"
)

func TestNewRastreamento_localizacao(t *testing.T) {
	lt := NewRastreamento_localizacao()
	lt.TasProdutolocalizacao(1, "Armazém A")

	if len(lt.Produtolocalizacao) != 1 {
		t.Errorf("Esperado 1 local rastreado, Tem %d", len(lt.Produtolocalizacao))
	}
}
