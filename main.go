package main

import (
	"Fintregas/domain"
	"Fintregas/server"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	repositorio := domain.NovoSimularEncomendaRepository("encomendas.csv")
	receberEncomenda := server.NovoReceberEncomendaService(repositorio)
	calcularRotaService := server.NovoCalcularRotaSevico(repositorio)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("=== Entregas de Encomendas ===")
		fmt.Println("1. Receber Encomenda")
		fmt.Println("2. Calcular Rota e Notificar Destinatario")
		fmt.Println("3. Sair")

		scanner.Scan()
		opcao := scanner.Text()

		switch opcao{
		case "1":
			encomendaID := ReceberDadosEncomenda(scanner)
			err := receberEncomenda.EncomendaRecebida(encomendaID)
			if err != nil{
				fmt.Println("Erro ao receber encomenda:", err)
			}else{
				fmt.Println("Encomenda recebida com sucesso.")
			}
		case "2":
			fmt.Print("Digita o ID de encomenda: ")
			scanner.Scan()
			encomendaID := scanner.Text()
			err := calcularRotaService.CalcularRota(encomendaID)
			if err != nil{
				fmt.Println("Erro ao Calcular rota:", err)
			}else{
				fmt.Println("Rota calculada e destinatario notificado com sucesso")
			}
		case "3":
			fmt.Println("Sain...")
			return
		default:
			fmt.Println("Opção Inválida. Tente Novamente.")
		}
	}
}

func ReceberDadosEncomenda(scanner *bufio.Scanner) domain.Encomenda {
	fmt.Print("Digite o ID da encomenda: ")
	scanner.Scan()
	id := scanner.Text()

	fmt.Print("Digite o peso da encomenda: ")
	scanner.Scan()
	peso,_:=strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Digite as dimensões (Comprimento, Largura, Altura) separadas por vírgulas : ")
	scanner.Scan()
	dimensoesInput := scanner.Text()
	dimensoesParts := strings.Split(dimensoesInput, ",")
	dimensoes := [3]float64{}
	
	for i, d := range dimensoesParts{
		dimensoes[i], _ = strconv.ParseFloat(strings.TrimSpace(d), 64)
	}

	fmt.Print("Digite Morada: ")
	scanner.Scan()
	morada := scanner.Text()

	fmt.Print("Digite Localização: ")
	scanner.Scan()
	localizacao := scanner.Text()

	fmt.Print("Digite o contato do destinatário: ")
	scanner.Scan()
	contato := scanner.Text()

	fmt.Print("O pagamento foi efetuado? (true/false): ")
	scanner.Scan()
	pagamentoEfetuado, _:= strconv.ParseBool(scanner.Text())

	return domain.Encomenda{
		ID: id,
		Peso: peso,
		Dimensoes: dimensoes,
		Morada: morada,
		ContatoDestinatario: contato,
		Localizacao: localizacao,
		PagamentoEfetuado: pagamentoEfetuado,
	}
}
