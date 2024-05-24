package server

import (
	"Fintregas/domain"
	"testing"
)

func TestEncomendaRecebida(t *testing.T) {
	caminhoarquivo := "test_encomendas.csv"
	repositorio := domain.NovoSimularEncomendaRepository(caminhoarquivo)
	servico := ServicoReceberEncomenda{repositorio: repositorio}

	encomenda := domain.Encomenda{
		ID:                  "123",
		Peso:                2.5,
		Dimensoes:           [3]float64{1, 2, 4},
		Morada:              "Luanda",
		ContatoDestinatario: "934-656-575",
		Localizacao:         "vianada (estalagem)",
		PagamentoEfetuado:   true,
	}

	err := servico.EncomendaRecebida(encomenda)
	if err != nil {
		t.Errorf("Erro ao receber encomenda: %v", err)
	}

	t.Run("Testando o ID da encomenda", func(t *testing.T) {
		pedido := ServicoReceberEncomenda{}
		esperado := pedido.FormalizarEncomenda(domain.Encomenda{})
		if esperado == nil {
			t.Error("ID da encomenda não pode ser vazio")
		}
	})

}

func TestFormalizarEncomenda(t *testing.T) {
	caminhoarquivo := "test_encomendas.csv"
	repositorio := domain.NovoSimularEncomendaRepository(caminhoarquivo)
	servico := ServicoReceberEncomenda{repositorio: repositorio}

	encomenda := domain.Encomenda{
		ID:                  "123",
		Peso:                2.5,
		Dimensoes:           [3]float64{1, 2, 4},
		Morada:              "Luanda",
		ContatoDestinatario: "934-656-575",
		Localizacao:         "vianada (estalagem)",
		PagamentoEfetuado:   true,
	}

	err := servico.FormalizarEncomenda(encomenda)
	if err != nil {
		t.Errorf("Erro ao formalizar encomenda: %v", err)
	}

	salvadoEncomenda, err := repositorio.ProcurarporID(encomenda.ID)
	if err != nil {
		t.Errorf("Erro ao encontrar encomenda: %v", err)
	}
	if salvadoEncomenda.ID != encomenda.ID {
		t.Errorf("Encomenda salva incorretamente: esperado %s, obtido %s", encomenda.ID, salvadoEncomenda.ID)
	}
}
