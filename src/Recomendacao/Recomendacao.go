package recomendacao

import (
	"sort"
)

type Produto struct {
	ID           int
	DataValidade string
	Quantidade   int
}

type Recomendacao struct {
	Produtos []Produto
}

func NovaRecomendacao() *Recomendacao {
	return &Recomendacao{}
}

func (r *Recomendacao) AddProdutos(produtoID int, datavalidade string, quantidade int) {
	r.Produtos = append(r.Produtos, Produto{
		ID:           produtoID,
		DataValidade: datavalidade,
		Quantidade:   quantidade,
	})
}

func (r *Recomendacao) Recomendar() []int {
	sort.Slice(r.Produtos, func(i, j int) bool {
		return r.Produtos[i].DataValidade < r.Produtos[j].DataValidade
	})

	var recomendarIDs []int
	for _, produto := range r.Produtos {
		recomendarIDs = append(recomendarIDs, produto.ID)
	}

	return recomendarIDs
}
