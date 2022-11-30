package client_one

import "github.com/Izhar7QR/PracticasGo/Practica/singleton/singleton.go"

func incrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
