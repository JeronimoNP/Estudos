package main

import (
	"Estudos/routes"
	"fmt"
	"net/http"
)

// Função principal de inicialização do http
func main() {
	//configurando resposta de escrita ao entrar no servidor, avisando que está funcionando.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})
	//local onde vai ficar os arquivos de requisição de pagina principal.
	fs := http.FileServer(http.Dir("../Estudos/Front-End/main"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	routes.Teste()

	http.ListenAndServe(":80", nil)
}
