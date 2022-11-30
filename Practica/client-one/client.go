package client_one

import "practicas/Practica/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
