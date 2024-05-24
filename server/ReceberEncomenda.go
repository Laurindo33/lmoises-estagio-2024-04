package server

import (
	"Fintregas/domain"
	"fmt"
	"log"
)

type ServicoReceberEncomenda struct {
	repositorio domain.EncomendaRepositorio
}



func NovoReceberEncomendaService(repositorio domain.EncomendaRepositorio) *ServicoReceberEncomenda {
	return &ServicoReceberEncomenda{repositorio: repositorio}
}

func (s *ServicoReceberEncomenda) EncomendaRecebida(encomenda domain.Encomenda) error {
	if encomenda.Morada != "Luanda" {
		return fmt.Errorf("encomenda fora de Luanda não são aceitas")
	}
	if encomenda.PagamentoEfetuado {
		return s.FormalizarEncomenda(encomenda)
	}
	return nil
}

func (s *ServicoReceberEncomenda) FormalizarEncomenda(encomenda domain.Encomenda) error {
	if encomenda.ID == "" {
		return fmt.Errorf("ID da encomenda não pode ser vazio")
	}
	if encomenda.Peso <= 0 {
		return fmt.Errorf("peso da encomenda deve ser maior que zero")
	}
	for i, dim := range encomenda.Dimensoes {
		if dim <= 0 {
			return fmt.Errorf("dimensão %d deve ser maior que zero", i)
		}
	}
	if encomenda.ContatoDestinatario == "" {
		return fmt.Errorf("contato do destinatário não pode ser vazio")
	}

	if err := s.repositorio.Salvar(encomenda); err != nil {
		return fmt.Errorf("erro ao salvar a encomenda: ")
	}

	log.Println("Encomenda formalizada com sucesso")
	return nil
}
