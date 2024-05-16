package application

import (
	"testing"
)

func TestCalcularRota(t *testing.T) {
	t.Run("Calculo da rota de encomendas", func(t *testing.T) {

		//arrange
		viatura := "Toyota Hiace"
		numeroEncomenda := 1
		esperado := "Toyota Hiace : 1"
		actual := CalcularRota(viatura, numeroEncomenda)

		//assert
		if esperado != actual {
			t.Errorf("erro com a rota %v  diferente de %v", esperado, actual)
			t.Fail()

		}

	})
}
