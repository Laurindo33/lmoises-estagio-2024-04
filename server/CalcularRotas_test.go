package server

import (
	"testing"
	"Fintregas/domain"
)

func TestCalcularRota(t* testing.T){
	caminhoarquivo := "test_encomendas.csv"
	repositorio := domain.NovoSimularEncomendaRepository(caminhoarquivo)
	servico := NovoCalcularRotaSevico(repositorio)

	encomenda := domain.Encomenda{
		ID:"123",
		Peso:2.5,
		Dimensoes: [3]float64{10, 5, 3},
		Morada: "Luanda",
		ContatoDestinatario: "934-656-575",
		PagamentoEfetuado: true,
	}
	repositorio.Salvar(encomenda)
	err := servico.CalcularRota(encomenda.ID)
    if err != nil {
        t.Errorf("Erro ao calcular rota: %v", err)
    }
}

func TestNotificarDestinatario(t *testing.T) {
	caminhoarquivo := "test_encomendas.csv"
    repositorio := domain.NovoSimularEncomendaRepository(caminhoarquivo)
    service := NovoCalcularRotaSevico(repositorio)

    encomenda := domain.Encomenda{
		ID: "123",
		Peso: 2.5, 
		Dimensoes: [3]float64{10, 5, 3}, 
		Morada: "Luanda", 
		ContatoDestinatario: "934-656-575", 
		PagamentoEfetuado: true, 
		Entregue: true,
	}
    repositorio.Salvar(encomenda)

    err := service.NotificarDestinatario(encomenda)
    if err != nil {
        t.Errorf("Erro ao notificar destinatário: %v", err)
    }
}