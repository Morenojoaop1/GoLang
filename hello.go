package main

import (
	"fmt"
)

func main() {

	nome := "João"
	versao := 1.2

	fmt.Println("Olá senhor", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar o Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
}
