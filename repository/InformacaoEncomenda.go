package repository

import (
	"Fintregas/domain"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	
)

type InformarEncomendarepeositorio struct {
	caminhoarquivo string
	encomendas map[string]domain.Encomenda
}

func NovoInformacaoEncomendaRepository(caminhoarquivo string)*InformarEncomendarepeositorio{
	repositorio := &InformarEncomendarepeositorio{
		caminhoarquivo: caminhoarquivo,
		encomendas: make(map[string]domain.Encomenda),
	}
	repositorio.Entrarparaficheiro()
	return repositorio
}

func(r *InformarEncomendarepeositorio) Entrarparaficheiro()error{
	ficheiro, err := os.Open(r.caminhoarquivo)
	if err != nil {
		if os.IsNotExist(err){
			return nil
		}
		return err
	}
	defer ficheiro.Close()

	escrever := csv.NewReader(ficheiro)
	listar, err := escrever.ReadAll()

	if err != nil{
		return err
	}

	for _,lista := range listar{
		if len(lista) < 9{
			continue
		}
		peso, _ := strconv.ParseFloat(lista[1], 64)
		dim1, _ := strconv.ParseFloat(lista[2], 64)
		dim2, _ := strconv.ParseFloat(lista[3], 64)
		dim3, _ := strconv.ParseFloat(lista[4], 64)
		pago, _ := strconv.ParseBool(lista[7])
		entregue := false
		if len(lista) >= 10{
			entregue,_ = strconv.ParseBool(lista[8])
		}

		encomenda := domain.Encomenda{
			ID: lista[0],
			Peso: peso,
			Dimensoes: [3]float64{dim1,dim2,dim3},
			Morada: lista[5],
			ContatoDestinatario: lista[6],
			PagamentoEfetuado: pago,
			Entregue: entregue,
		}
		r.encomendas[encomenda.ID] = encomenda
	}
	return nil
}

func(r *InformarEncomendarepeositorio) Salvar(encomenda domain.Encomenda)error{
	r.encomendas[encomenda.ID] = encomenda

	ficheiro, err := os.OpenFile(r.caminhoarquivo, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		return fmt.Errorf("erro ao abrir o arquivo: %v", err)
	}
	defer ficheiro.Close()

	lercsv := csv.NewWriter(ficheiro)
	defer lercsv.Flush()
	for _, encomenda := range r.encomendas{
		lista := []string{
			encomenda.ID,
			fmt.Sprintf("%.2f", encomenda.Peso),
			fmt.Sprintf("%.2f", encomenda.Dimensoes[0]),
			fmt.Sprintf("%.2f", encomenda.Dimensoes[1]),
			fmt.Sprintf("%.2f", encomenda.Dimensoes[2]),
			encomenda.Morada,
			encomenda.ContatoDestinatario,
			fmt.Sprintf("%t", encomenda.PagamentoEfetuado),
			fmt.Sprintf("%t", encomenda.Entregue),
		}
		if err := lercsv.Write(lista); err != nil{
			return fmt.Errorf("erro ao escrever no arquivo CSV: %v", err)
		}
	}
	return nil
}

func (r *InformarEncomendarepeositorio) ProcurarporID(id string)(domain.Encomenda, error){
	encomenda, existe := r.encomendas[id]
	if !existe{
		return domain.Encomenda{},fmt.Errorf("encomenda não encontrada")
	}
	return encomenda, nil
}