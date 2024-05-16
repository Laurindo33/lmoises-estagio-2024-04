package application

import (
    "testing"
    "Fintregas/domain"
)

func TestEncomendaEntregue(t *testing.T) {
    t.Run("Solicitar entrega com encomenda válida", func(t *testing.T) {
        encomenda := domain.Encomenda{
            EncomendaID :"1",
            Largura     :23.4,
            Altura      :21.3,
            Peso        :23.4,
            NomeCliente	     :"Paulo Ferreira",
            ContactoDestinatario : "932827742",
            NomeDestinatario : "Mateus silva",
            MoradaDestinatario :"Luanda",
        }

      
     
    })
}
