package domain

type Pedido struct {
	ID          string
	Dimensoes   string
	Peso        float64
	Morada      string
	Contacto    string
	Pago        bool
	Formalizado bool
	Entrega     Entrega
}
