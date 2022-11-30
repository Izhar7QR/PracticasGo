package client_two

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
