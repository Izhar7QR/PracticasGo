package client_one

func incrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
