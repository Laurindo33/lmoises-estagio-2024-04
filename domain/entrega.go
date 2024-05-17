package domain

type Entrega struct {
	ID           string
	Status       string
	Destinatario string
}

type EntregaService interface {
	NotificarDestinatario(entrega Entrega) error
	AtualizarFila(entrega Entrega) error
}
