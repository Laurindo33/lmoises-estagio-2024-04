package recomendacao

import (
	"testing"
)

func TestRecomendacao(t *testing.T) {
	rec := NovaRecomendacao()
	rec.AddProdutos(1, "2024-05-15", 100)
	rec.AddProdutos(2, "2024-05-10", 50)
	rec.AddProdutos(3, "2024-05-20", 200)

	recomendado := rec.Recomendar()

	expected := []int{2, 1, 3}
	for i, id := range recomendado {
		if id != expected[i] {
			t.Errorf("Esperado ID %d, tem ID %d", expected[i], id)
		}
	}
}
