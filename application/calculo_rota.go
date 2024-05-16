package application

import "strconv"

// CalculoRota representa a lógica para calcular a rota mais eficiente
type CalculoRota struct{}

// CalcularRota calcula a rota mais eficiente para a entrega
func CalcularRota(viatura string, numeroEncomenda int) string {

	rota := viatura + " : " + strconv.Itoa(numeroEncomenda)

	return rota
}
