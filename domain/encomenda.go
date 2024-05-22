package domain

type Encomenda struct {
	ID        string
	Peso      float64
	Dimensoes [3]float64
	Morada    string
	Contacto  string
	Status    string
	Pagamento bool
}
