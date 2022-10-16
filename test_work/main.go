package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

//ab -n 10000 -c 10 http://localhost:8080/admin/request

type App struct {
	applications string
	showCounts   int
}

var app = make([]App, 50)

func main() {
	tomain()
}

func httpFunc() {
	http.HandleFunc("/request", _request)
	http.HandleFunc("/admin/requests", _adminRequests)
	http.ListenAndServe(":8080", nil)
}

func _request(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		randNumber := randInt(0, 50)
		w.Write([]byte(fmt.Sprintf("%s", app[randNumber].applications)))
		app[randNumber].showCounts++
		return
	}
	return
}

func _adminRequests(w http.ResponseWriter, r *http.Request) {
	_, ok := w.(http.Flusher)
	if !ok {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		for i := 0; i < len(app); i++ {
			//w.WriteHeader(http.StatusOK) //superfluous response.WriteHeader call from main._adminRequests
			w.Write([]byte(fmt.Sprintf("%s : %v\n", app[i].applications, app[i].showCounts)))
		}
		return
	}
	return
}

func tomain() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(app); i++ { //Инициализация заявок
		app[i].applications = randomString(2)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Done()
	defer wg.Done()
	go listUpdate()
	go httpFunc()
	wg.Wait()
}

//Рандомная замена элемента слайса
func listUpdate() {
	for true {
		apSize := randInt(0, 50)
		app = append(app, app[apSize])
		app[apSize].applications = randomString(2)
		app[apSize].showCounts = 0
		time.Sleep(time.Millisecond * 200)
	}
}

//Получение рандомной строки длинной 2 символа
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(97, 123))
	}
	return string(bytes)
}

//рандомное число в указаном диапазоне
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
