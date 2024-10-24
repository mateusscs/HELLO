package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3 //Numero de vezes que os sites serão monitorados
const delay = 5          //Atraso entre cada teste
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

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://www.alura.com.br", "https://www.twitch.tv/",
		"https://www.globo.com/"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println(i, site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Error:", resp.StatusCode)
	}
}
