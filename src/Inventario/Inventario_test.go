package Inventario

import (
	"testing"
)

func TestInventario(t *testing.T) {

	// Teste para verificar se a função NewInventory() retorna um ponteiro para Inventario não nulo
	t.Run("NovoInventario retorna um ponteiro não nulo para Inventario", func(t *testing.T) {
		inventario := NovoInventario()
		if inventario == nil {
			t.Errorf("NewInventory() returned nil")
		}
	})

}

func TestAddProduct(t *testing.T) {
	inventario := NovoInventario()

	inventario.AddProduto(1, "Produto A", "Descrição A", "Categoria A")

	if len(inventario.Produtos) != 1 {
		t.Errorf("Esperado 1 Produto, Tem %d", len(inventario.Produtos))
	}

}

func TestAddLot(t *testing.T) {
	inventario := NovoInventario()
	inventario.AddProduto(1, "Produto A", "Descrição A", "Categoria A")
	config := DadosLote {
		lotID                  : 1,
		produtoID              : 1,
		dataFabricacao         : "2024-10-01",
		dataValidade           : "2024-12-02",
		quantidade             : 100,
		localizacaoArmazem     : "Local",
		condicoesarmazenamento : "Fresco",
		status                 :"diponivel",
		Informaçõesfornecedor  : "",
	}
	inventario.AddLot(config)

	if len(inventario.Lotes) != 1 {
		t.Errorf("Esperado 1 lot, Tem %d", len(inventario.Lotes))
	}
}
