package recomendacao

import (
	"sort"
)

type Produto struct {
	ID            int
	Data_validade string
	Quantidade    int
}

type Recomendacao struct {
	Produtos []Produto
}

func NewRecomendacao() *Recomendacao {
	return &Recomendacao{}
}

func (r *Recomendacao) Addprodutos(produtoID int, datavalidade string, quantidade int) {
	r.Produtos = append(r.Produtos, Produto{
		ID:            produtoID,
		Data_validade: datavalidade,
		Quantidade:    quantidade,
	})
}

func (r *Recomendacao) Recomendar() []int {
	sort.Slice(r.Produtos, func(i, j int) bool {
		return r.Produtos[i].Data_validade < r.Produtos[j].Data_validade
	})

	var recomendarIDs []int
	for _, product := range r.Produtos {
		recomendarIDs = append(recomendarIDs, product.ID)
	}

	return recomendarIDs
}
