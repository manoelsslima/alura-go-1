package main

import (
	"fmt"
)

func main() {
	fmt.Println("Informe a opção desejada")
	fmt.Println("1 - Iniciar")
	fmt.Println("2 - Executar ação")
	fmt.Println("9 - Sair da aplicação")
	var opcao int
	fmt.Scanf("%d", &opcao)
	fmt.Println("A opção escolhida foi", opcao)

	if opcao == 1 {
		fmt.Println("Iniciando...")
	} else if opcao == 2 {
		fmt.Println("Executando ação...")
	} else if opcao == 9 {
		fmt.Println("Saindo...")
	} else {
		fmt.Println("Opção inválida!")
	}

}
