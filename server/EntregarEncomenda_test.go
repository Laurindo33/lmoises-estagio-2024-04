package server

import (
	"Fintregas/domain"
	"Fintregas/repository"
	"os"
	"testing"
)

func TestEntregarEncomenda(t *testing.T) {
	caminhoarquivo := "test_encomendas.csv"
	repositorio := repository.NovoInformacaoEncomendaRepository(caminhoarquivo)
	servico := NovoEntregarEncomendaService(repositorio)

	encomenda := domain.Encomenda{
		ID:                  "123",
		Peso:                2.5,
		Dimensoes:           [3]float64{10, 5, 3},
		Morada:              "Luanda",
		ContatoDestinatario: "test@example.com",
		Localizacao:         "vianada (estalagem)",
		PagamentoEfetuado:   true,
		Entregue:            true,
	}
	repositorio.Salvar(encomenda)
	err := servico.EntregarEncomenda(encomenda.ID)
	if err != nil {
		t.Errorf("Erro ao entregar encomenda: %v", err)
	}

	listaEncomenda, err := repositorio.ProcurarporID(encomenda.ID)
	if err != nil {
		t.Errorf("Erro ao encontrar encomenda: %v", err)
	}
	if !listaEncomenda.Entregue {
		t.Errorf("A encomeda não foi marcada entregue")
	}

	os.Remove(caminhoarquivo)
}

func TestNotificarCliente(t *testing.T) {
	caminhoarquivo := "test_encomendas.csv"
	repositorio := repository.NovoInformacaoEncomendaRepository(caminhoarquivo)
	servico := NovoEntregarEncomendaService(repositorio)

	encomenda := domain.Encomenda{
		ID:                  "123",
		Peso:                2.5,
		Dimensoes:           [3]float64{10, 5, 3},
		Morada:              "Luanda",
		ContatoDestinatario: "test@example.com",
		Localizacao:         "vianada (estalagem)",
		PagamentoEfetuado:   true,
		Entregue:            true,
	}
	repositorio.Salvar(encomenda)

	err := servico.NotificalCliente(encomenda)
	if err != nil {
		t.Errorf("Erro ao notificar cliente: %v", err)
	}

	os.Remove(caminhoarquivo)
}
