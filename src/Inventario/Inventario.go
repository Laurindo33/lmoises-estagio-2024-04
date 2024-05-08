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

type DadosLote struct {
	lotID                  int
	produtoID              int
	dataFabricacao         string
	dataValidade           string
	quantidade             int
	localizacaoArmazem     string
	condicoesarmazenamento string
	status                 string
	Informaçõesfornecedor  string
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

func (inv *Inventario) AddLot(config DadosLote, anexos ...string) {
	_, exists := inv.Produtos[config.produtoID]
	if !exists {
		panic("Produto não existe")
	}

	inv.Lotes[config.lotID] = Lote{
		ID:                     config.lotID,
		DataFabricacao:         config.dataFabricacao,
		DataValidade:           config.dataValidade,
		Quantidade:             config.quantidade,
		LocalizacaoArmazem:     config.localizacaoArmazem,
		CondicoesArmazenamento: config.condicoesarmazenamento,
		Status:                 config.status,
		HistoricoMovimento:     []string{},
		InformaçõesFornecedor:  "",
		ObservacaoNota:         "",
	}
}
