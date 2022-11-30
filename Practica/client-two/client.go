package client_two

import "practicas/Practica/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
