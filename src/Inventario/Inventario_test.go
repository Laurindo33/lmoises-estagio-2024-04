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
	inventario.AddLot(1, 1, "2024-04-30", "2024-05-30", 100, "Armazém A", "condições A", "Disponivel", "Recomendado usar em lugares frescos")

	if len(inventario.Lotes) != 1 {
		t.Errorf("Esperado 1 lot, Tem %d", len(inventario.Lotes))
	}
}
