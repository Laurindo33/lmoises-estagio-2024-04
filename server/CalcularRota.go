package server

import (
	"Fintregas/domain"
	"fmt"
)

type CalcularRotaSevico struct{
	repositorio domain.EncomendaRepositorio
}

func NovoCalcularRotaSevico(repositorio domain.EncomendaRepositorio)*CalcularRotaSevico{
	return &CalcularRotaSevico{repositorio: repositorio}
}

func (s *CalcularRotaSevico)CalcularRota(encomendaID string)error{
	encomenda, err := s.repositorio.ProcurarporID(encomendaID)
	if err != nil {
        return fmt.Errorf("erro ao encontrar encomenda: %v", err)
    }
	rotaCalculada := fmt.Sprintf("Rota para entrega da encomenda %s foi calculada.", encomendaID)
    fmt.Println(rotaCalculada)

	return s.NotificarDestinatario(encomenda)
}

func (m *CalcularRotaSevico)NotificarDestinatario(encomenda domain.Encomenda)error{
	mensagem := fmt.Sprintf("Destinatário %s foi notificado sobre a entrega da encomenda %s.", encomenda.ContatoDestinatario, encomenda.ID)
    fmt.Println(mensagem)
    return nil
}