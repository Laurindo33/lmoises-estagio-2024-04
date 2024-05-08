package Rm_localizacao

type Rastreamento_localizacao struct {
	Produtolocalizacao map[int]string
}

func NewRastreamento_localizacao() *Rastreamento_localizacao {
	return &Rastreamento_localizacao{
		Produtolocalizacao: make(map[int]string),
	}
}

func (lt *Rastreamento_localizacao) TasProdutolocalizacao(produtoID int, localizacao string) {
	lt.Produtolocalizacao[produtoID] = localizacao
}
