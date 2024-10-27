package main

import (
	"log"
	"net/http"
)

// viewHandler ответ на определенный запрос
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)            // регистрируем обработчик для сервера, сервер всегда ищет обработчики
	err := http.ListenAndServe("localhost:8080", nil) // обрабатывает запрос браузера и отвечает на него, грубо говоря он ищет хендлеры
	log.Fatal(err)
}
