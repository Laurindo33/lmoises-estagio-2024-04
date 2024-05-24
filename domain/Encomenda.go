package domain

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Encomenda struct {
	ID                  string
	Peso                float64
	Dimensoes           [3]float64
	Morada              string
	ContatoDestinatario string
	Localizacao         string
	PagamentoEfetuado   bool
	Entregue            bool
}

type EncomendaRepositorio interface {
	Salvar(encomenda Encomenda) error
	ProcurarporID(id string) (Encomenda, error)
}

type SimuladoEncomendaRepositorio struct {
	encomendas map[string]Encomenda
	csvFile    string
}

func NovoSimularEncomendaRepository(csvFile string) *SimuladoEncomendaRepositorio {

	return &SimuladoEncomendaRepositorio{
		csvFile:    csvFile,
		encomendas: make(map[string]Encomenda),
	}
}

func (m *SimuladoEncomendaRepositorio) Salvar(encomenda Encomenda) error {
	if m.encomendas == nil {
		m.encomendas = make(map[string]Encomenda)
	}
	m.encomendas[encomenda.ID] = encomenda
	file, err := os.OpenFile(m.csvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo CSV: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		fmt.Sprintf("ID: %v", encomenda.ID),
		fmt.Sprintf("Peso: %.1f ", encomenda.Peso),
		fmt.Sprintf("Dimesoes: %.1f,%.1f,%.1f ", encomenda.Dimensoes[0], encomenda.Dimensoes[1], encomenda.Dimensoes[2]),
		fmt.Sprintf("Morada: %v ", encomenda.Morada),
		fmt.Sprintf("Contato Destinatario: %v ", encomenda.ContatoDestinatario),
		fmt.Sprintf("Pagamento Efetuado: %t", encomenda.PagamentoEfetuado),
		fmt.Sprintf("Entregue: %t", encomenda.Entregue),
	}

	if err := writer.Write(record); err != nil {
		return fmt.Errorf("erro ao escrever no arquivo CSV: %v", err)
	}
	return nil
}

func (m *SimuladoEncomendaRepositorio) ProcurarporID(id string) (Encomenda, error) {
	encomenda, exists := m.encomendas[id]
	if !exists {
		return Encomenda{}, fmt.Errorf("encomenda não encontrada")
	}
	return encomenda, nil
}
