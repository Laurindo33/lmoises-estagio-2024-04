package rm_localizacao

import (
	"testing"
)

func TestNewRastreamentoLocalizacao(t *testing.T) {
	lt := NovoRastreamentoLocalizacao()
	lt.TasProdutolocalizacao(1, "Armazém A")

	if len(lt.Produtolocalizacao) != 1 {
		t.Errorf("Esperado 1 local rastreado, Tem %d", len(lt.Produtolocalizacao))
	}
}
