package entrega

type Pedido struct {
	dimensoes []int
	peso      int
	Endereco  string
	contatos  string
	//formalizada bool
	//encomendaPaga bool
}

func (o *Pedido) FormalizarPedido() string {
	// Lógica para formalizar o pedido
	return "Pedido formalizado com sucesso"
}
