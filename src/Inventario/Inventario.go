package Inventario

type Produto struct {
	ID        int
	Nome      string
	Descricao string
	Categoria string
}

type Lote struct {
	ID                      int
	Data_fabricacao         string
	Data_validade           string
	Quantidade              int
	Localizacao_armazem     string
	Condicoes_armazenamento string
	Status                  string
	Historico_movimento     []string
	Informações_fornecedor  string
	Observacao_nota         string
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

func (inv *Inventario) AddLot(lotID int, produtoID int, datafabricacao string, datavalidade string, quantidade int, localizacaoarmazem string, condicoesarmazenamento string, status string, Informaçõesfornecedor string) {
	_, exists := inv.Produtos[produtoID]
	if !exists {
		panic("Produto não existe")
	}

	inv.Lotes[lotID] = Lote{
		ID:                      lotID,
		Data_fabricacao:         datafabricacao,
		Data_validade:           datavalidade,
		Quantidade:              quantidade,
		Localizacao_armazem:     localizacaoarmazem,
		Condicoes_armazenamento: condicoesarmazenamento,
		Status:                  status,
		Historico_movimento:     []string{},
		Informações_fornecedor:  "",
		Observacao_nota:         "",
	}
}
