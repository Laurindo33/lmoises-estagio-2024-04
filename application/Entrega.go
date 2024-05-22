package application

import (
	"Fintregas/domain"
	"errors"
)

type EntregaService struct {
	Repo domain.EntregaRepository
}

func NewEntregaService(repo domain.EntregaRepository) *EntregaService {
	return &EntregaService{Repo: repo}
}

func (s *EntregaService) ReceberEncomenda(encomenda domain.Encomenda) error {
	if encomenda.Morada != "Luanda" {
		return errors.New("A modada não pode ser Fora de Luanda")
	}
	return s.Repo.AddEncomenda(encomenda)
}

func (s *EntregaService) FormalizarPagamento(id string , encomenda domain.Encomenda) string {
	encomenda, err := s.Repo.GetEncomenda(id)
	

	
	if err != nil {
		encomenda.Pagamento = true
		return  "Pagamento formalizado com sucesso"
	}
	return "Erro ao Formalizar"
}

func (s *EntregaService) VerificarPagamento(encomenda domain.Encomenda) bool {
	if encomenda.Pagamento == false {
		return true
	}

	return false
}
