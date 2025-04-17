package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	Result int `json:"result"`
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	// Обрабатываем только POST-запросы
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	// Разбор формы
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка разбора формы", http.StatusBadRequest)
		return
	}
	// Получение значений
	aStr := r.FormValue("a")
	bStr := r.FormValue("b")
	// Преобразование в числа
	a, err1 := strconv.Atoi(aStr)
	b, err2 := strconv.Atoi(bStr)
	if err1 != nil || err2 != nil {
		http.Error(w, "Неверные числа", http.StatusBadRequest)
		return
	}

	res := Response{Result: a + b}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
func main() {
	http.HandleFunc("/sum", sumHandler)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
