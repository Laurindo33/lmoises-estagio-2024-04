package Inventario

type Produto struct {
	ID        int
	Nome      string
	Descricao string
	Categoria string
}

type Lote struct {
	ID                     int
	DataFabricacao         string
	DataValidade           string
	Quantidade             int
	LocalizacaoArmazem     string
	CondicoesArmazenamento string
	Status                 string
	HistoricoMovimento     []string
	InformaçõesFornecedor  string
	ObservacaoNota         string
}

type Inventario struct {
	Produtos map[int]Produto
	Lotes    map[int]Lote
}

func NovoInventario() *Inventario {
	return &Inventario{
		Produtos: make(map[int]Produto),
		Lotes:    make(map[int]Lote),
	}
}

func (i *Inventario) AddProduto(id int, nome, descricao, categoria string) {

	i.Produtos[id] = Produto{
		ID:        id,
		Nome:      nome,
		Descricao: descricao,
		Categoria: categoria,
	}

}

func (inv *Inventario) AddLot(lotID int, produtoID int, dataFabricacao string, dataValidade string, quantidade int, localizacaoArmazem string, condicoesarmazenamento string, status string, Informaçõesfornecedor string) {
	_, exists := inv.Produtos[produtoID]
	if !exists {
		panic("Produto não existe")
	}

	inv.Lotes[lotID] = Lote{
		ID:                     lotID,
		DataFabricacao:         dataFabricacao,
		DataValidade:           dataValidade,
		Quantidade:             quantidade,
		LocalizacaoArmazem:     localizacaoArmazem,
		CondicoesArmazenamento: condicoesarmazenamento,
		Status:                 status,
		HistoricoMovimento:     []string{},
		InformaçõesFornecedor:  "",
		ObservacaoNota:         "",
	}
}
