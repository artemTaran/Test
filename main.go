package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var applications = make([]string, 50)
var delAppCounts = make([]string, 50)
var showCounts = make([]int, 50)

//var applications [50]string
//var delAppCounts [50]string

func main() {
	tomain()
}

func myHt() {
	http.HandleFunc("/request", _request)
	http.HandleFunc("/admin/requests", _admin_requests)
	http.ListenAndServe(":8080", nil)
}

func _request(w http.ResponseWriter, r *http.Request) {
	randNumber := randInt(0, 50)
	w.Write([]byte(fmt.Sprintf("%v", applications[randNumber])))
	showCounts[randNumber]++
}

func _admin_requests(w http.ResponseWriter, r *http.Request) {

	for range applications {
		w.Write([]byte(fmt.Sprintf("%v", applications)))
	}

	w.Write([]byte(fmt.Sprintf("%v", applications)))
}

func tomain() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(applications); i++ { //Инициализация заявок
		applications[i] = randomString(2)
		delAppCounts[i] = applications[i]
	}
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Done()
	defer wg.Done()
	go listUpdate(applications)
	go myHt()
	wg.Wait()
}

//Рандомная замена элемента слайса
func listUpdate(applications []string) {
	for true {
		apSize := randInt(0, 50)
		(applications)[apSize] = randomString(2)
		delAppCounts = append(delAppCounts, applications[apSize]) //добавление удалённого элемента в отдельный слайс
		time.Sleep(time.Millisecond * 200)
		//fmt.Println(applications)
		fmt.Println(delAppCounts)
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
