package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	// var opcao = capturaOpcaoMenu()
	opcao := capturaOpcaoMenu()

	switch opcao {
	case 1:
		fmt.Println("Iniciando...")
	case 2:
		fmt.Println("Executando ação...")
	case 9:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Opção inválida!")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	fmt.Println("Bem-vindo ao sistema!")
	fmt.Println("")
}

func exibeMenu() {
	fmt.Println("Informe a opção desejada")
	fmt.Println("1 - Iniciar")
	fmt.Println("2 - Executar ação")
	fmt.Println("9 - Sair da aplicação")
}

func capturaOpcaoMenu() int {
	var opcao int
	fmt.Scanf("%d", &opcao)
	fmt.Println("A opção escolhida foi", opcao)
	return opcao
}
