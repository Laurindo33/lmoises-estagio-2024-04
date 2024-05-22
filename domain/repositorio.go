package domain
type EntregaRepository interface {
	AddEncomenda(encomenda Encomenda) error
	GetEncomenda(id string) (Encomenda, error)
	UpdateEncomenda(encomenda Encomenda) error
	// Outras funções necessárias
   }
   