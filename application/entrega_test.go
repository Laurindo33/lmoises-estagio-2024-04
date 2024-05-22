package application

import (
	"Fintregas/domain"
	"Fintregas/repository"
	"testing"
)

func TestReceberEncomenda(t *testing.T) {
	Repo := repository.NewMemoryRepository()
	Service := NewEntregaService(Repo)

	Encomenda := domain.Encomenda{
		ID:        "123",
		Peso:      2.5,
		Morada:    "Luanda",
		Contacto:  "9438843737",
		Dimensoes: [3]float64{1, 2, 3},
		Status:    "Recebido",
		Pagamento: true,
	}

	err := Service.ReceberEncomenda(Encomenda)
	if err != nil {
		t.Errorf("erro ao Receber a encomenda: %v", err)
	}
}

func TestVerificarPagamento_Efectuado(t *testing.T) {
	//Arrange
	Repo := repository.NewMemoryRepository()
	Service := NewEntregaService(Repo)
	esperado := false
	encomenda := domain.Encomenda{
		Pagamento: true,
	}
	err := Service.VerificarPagamento(encomenda)

	//Act

	//Assert
	if esperado != err {
		t.Errorf("O pagamento não está formalizado")
	}
}

func testFormalizarPagamento(t *testing.T) {
	//arrange
	Repo := repository.NewMemoryRepository()
	Service := NewEntregaService(Repo)
	Encomenda := domain.Encomenda{
		ID:        "123",
		Peso:      2.5,
		Morada:    "Luanda",
		Contacto:  "9438843737",
		Dimensoes: [3]float64{1, 2, 3},
		Status:    "Recebido",
		Pagamento: false,
	}
	actual := Service.FormalizarPagamento(Encomenda.ID, Encomenda)
	esperado:="Pagamento formalizado com sucesso"

	//Act

	//Assert
	if esperado != actual {
		t.Errorf("Erro ao formalizar o pagamento, %v diferente de %v",esperado, actual)

	}

}
