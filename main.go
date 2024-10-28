package main

import (
	"log"
	"net/http"
)

// viewHandler ответ на определенный запрос
func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!")
}
func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut, web!")
}
func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, web!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)         // регистрируем обработчик для сервера, сервер всегда ищет обработчики
	http.HandleFunc("/salut", frenchHandler)          // регистрируем обработчик для сервера, сервер всегда ищет обработчики
	http.HandleFunc("/namaste", hindiHandler)         // регистрируем обработчик для сервера, сервер всегда ищет обработчики
	err := http.ListenAndServe("localhost:8080", nil) // обрабатывает запрос браузера и отвечает на него, грубо говоря он ищет хендлеры
	log.Fatal(err)
}
