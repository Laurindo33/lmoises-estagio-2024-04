package model

type Cliente struct {
	Id           int
	Nome         string
	Contacto     string
	Destinatario Destinatario
	Encomenda    Encomenda
}
