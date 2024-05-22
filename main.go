package intregas

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	Scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Digite um comando (criar, obter, atualizar, sair):")
		Scanner.Scan()
		input := strings.TrimSpace(Scanner.Text())

		switch input {
		case "criar":
			fmt.Println("Criando uma nova encomenda…")
			// Código para criar encomenda

		case "atualizar":
			fmt.Println("Atualizando uma encomenda…")
			// Código para atualizar encomenda
		case "sair":
			fmt.Println("Saindo…")
			return
		default:
			fmt.Println("Comando não reconhecido.")
		}
	}
}
