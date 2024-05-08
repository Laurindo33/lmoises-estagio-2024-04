package main

import (
	"fmt"

	"Acme_lda/Inventario"
	recomendacao "Acme_lda/Recomendacao"
	RmLocalizacao "Acme_lda/rm_localizacao"
	"encoding/csv"
	"os"
	"strconv"
)

// geral
func criarCsv(dados [][]string, nomeArquivo string) {
	// Abrir o arquivo CSV existente (ou criar um novo se não existir)
	file, err := os.OpenFile(nomeArquivo+".csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Ler o conteúdo do arquivo CSV existente
	existingData, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}

	// Combinar os dados existentes com os novos dados
	allData := append(existingData, dados...)

	// Posicionar o cursor no início do arquivo para reescrever todos os dados
	file.Seek(0, 0)

	// Criar um escritor CSV para o arquivo
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escrever todos os dados (existentes + novos) de volta para o arquivo
	for _, row := range allData {
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}
}

func Inventariosave() {
	//Dados Produtos
	id_produto := 1
	nome_produto := "Produto A"
	nome_discrcao := "Discrição"
	nome_categoria := "Categoria"

	//Dados Lotes
	id_lote := 1
	data_fabrico := "2024-10-05"
	data_validade := ""
	quantidade := 12
	localizacao_armazem := ""
	condicoes_armazenamento := ""
	status := ""
	Informaçõesfornecedor := ""

	inventario := Inventario.NovoInventario()

	inventario.AddProduto(1, "Produto A", "Discrição", "Categoria A")

	//inventario.AddLot(1, 1, "", "", 100, "", "", "", "")

	DadosProduto := [][]string{
		{"Id_Produto: " + strconv.Itoa(id_produto) + " || Produto B : " + nome_produto + " || Discrição : " + nome_discrcao + " || Categoria: " + nome_categoria},
	}
	criarCsv(DadosProduto, "Produto")

	DadosLot := [][]string{
		{"Id_lote: " + strconv.Itoa(id_lote) + " || Id_Produto: " + strconv.Itoa(id_produto) + " || Data de fabrico : " + data_fabrico + " || Data de Validade: " + data_validade + " || quantidade:" + strconv.Itoa(quantidade) + " || Localização do Armazem : " + localizacao_armazem + " || Condições do Armazem: " + condicoes_armazenamento + " || status: " + status + " || Imformações do fornecedor : " + Informaçõesfornecedor},
	}
	criarCsv(DadosLot, "Lote")
}

func RmLocalizacaosave() {
	Id_localizacao := 1
	localizacao := "Armazém A"

	Rl := RmLocalizacao.NovoRastreamentoLocalizacao()
	Rl.TasProdutolocalizacao(1, "Armazém A")
	Rl.TasProdutolocalizacao(2, "Armazém B")

	DadosLocalizacao := [][]string{
		{"Id_localização: " + strconv.Itoa(Id_localizacao) + " || Localização : " + localizacao},
	}
	criarCsv(DadosLocalizacao, "Localização")
}

func recomendacaosave() {

	rec := recomendacao.NovaRecomendacao()

	rec.AddProdutos(1, "2024-05-09", 100)
	rec.AddProdutos(2, "2024-05-10", 50)
	rec.AddProdutos(4, "2024-05-20", 200)

	DadosRecomendacao := [][]string{
		{"Id_recomendação : " + strconv.Itoa(2) + " || data de validade : " + "2024-05-19" + " || Quantidade : " + strconv.Itoa(100)},
	}
	criarCsv(DadosRecomendacao, "Recomendação")

	recomendedado := rec.Recomendar()
	fmt.Println("Recomendado produtos IDs para envio:", recomendedado)
}

func main() {
	Inventariosave()
	RmLocalizacaosave()
	recomendacaosave()
}
