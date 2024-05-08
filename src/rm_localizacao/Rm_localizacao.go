package rm_localizacao

type RastreamentoLocalizacao struct {
	Produtolocalizacao map[int]string
}

func NovoRastreamentoLocalizacao() *RastreamentoLocalizacao {
	return &RastreamentoLocalizacao{
		Produtolocalizacao: make(map[int]string),
	}
}

func (lt *RastreamentoLocalizacao) TasProdutolocalizacao(produtoID int, localizacao string) {
	lt.Produtolocalizacao[produtoID] = localizacao
}
