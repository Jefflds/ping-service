package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func ping(url string, expectedResult string) {
	cmd := exec.Command("cmd.exe", "/C", "start", "cmd.exe")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Erro ao abrir o terminal: %v\n", err)
		return
	}

	for {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Erro ao fazer ping: %v\n", err)
		} else {
			if resp.StatusCode == http.StatusOK {
				var body = make([]byte, 64)
				_, err := resp.Body.Read(body)
				if err == nil && string(body) == expectedResult {
					fmt.Println("Ping bem-sucedido: mensagem correta recebida.")
					fmt.Println(string(body))
				} else {
					fmt.Println("Ping bem-sucedido, mas a mensagem não é a esperada.")
					fmt.Println(string(body))
				}
			} else {
				fmt.Printf("Ping para %s: Status %s\n", url, resp.Status)
			}
			resp.Body.Close()
		}
		time.Sleep(30 * time.Minute)
	}
}

func main() {
	url := "https://diy-crypto-api.onrender.com/"
	go ping(url, "Welcome to DIY Crypto API Server!")

	select {}
}
