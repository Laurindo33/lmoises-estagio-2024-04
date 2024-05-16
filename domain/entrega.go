package domain

import (
	"fmt"
)

// EntregaConcluida representa uma entrega concluída
type EntregaConcluida struct {
	encomendaID    string
	FilaDeEntregas *FilaDeEntregas

}
type FilaDeEntregas struct {
	entregas []EntregaConcluida
}
// RegistrarEntregaConcluida registra uma entrega concluída
func (e *EntregaConcluida) RegistrarEntregaConcluida(encomendaID string, fila *FilaDeEntregas) {
	e.encomendaID = encomendaID
	e.FilaDeEntregas = fila
	fila.AdicionarEntrega(*e)

}

func (f *FilaDeEntregas) AdicionarEntrega(entrega EntregaConcluida) {
	f.entregas = append(f.entregas, entrega)
}

func (f *FilaDeEntregas) RemoverEntregaConcluida() {
	if len(f.entregas) > 0 {
		fmt.Printf("Entrega concluída para %s removida da fila.\n", f.entregas[0].encomendaID)
		f.entregas = f.entregas[1:]
	} else {
		fmt.Println("Não há entregas concluídas na fila.")
	}
}