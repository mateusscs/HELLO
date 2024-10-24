package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {
	introducao()
	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibir logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
	
}

func introducao() {
	nome := "Mateus"
	versao := 1.23
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento(){
	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"

	fmt.Println(sites)
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200{
		fmt.Println("Site:",site,"foi carregado com sucesso")
	}else{
		fmt.Println("Site:",site,"esta com problemas. Error:",resp.StatusCode)
	}

}


func exibeNomes(){
	nomes := []string{""}
}