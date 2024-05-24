package server

import (
	"Fintregas/domain"
	"fmt"
)

type EntregarEncomendaService struct{
	repositorio domain.EncomendaRepositorio
}

func NovoEntregarEncomendaService(repositorio domain.EncomendaRepositorio)*EntregarEncomendaService{
	return &EntregarEncomendaService{repositorio: repositorio}
}

func (s *EntregarEncomendaService)EntregarEncomenda(encomendaID string)error{
	encomenda, err := s.repositorio.ProcurarporID(encomendaID)
	if err != nil{
		return fmt.Errorf("erro ao encontrar encomenda: %v", err)
	}

	encomenda.Entregue = true

	err = s.repositorio.Salvar(encomenda)
	if err != nil{
		return fmt.Errorf("erro ao salvar a encomenda: %v",err)
	}

	fmt.Print("Encomenda entregue com sucesso!")
	return s.NotificalCliente(encomenda)
}

func (s *EntregarEncomendaService)NotificalCliente(encomenda domain.Encomenda)error{
	mensagem := fmt.Sprintf("Encomenda ID: %s entregue com sucesso para %s", encomenda.ID, encomenda.ContatoDestinatario) 
	fmt.Print(mensagem)
	return nil
}