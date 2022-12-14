package main

import (
	"sync"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func testReg(t *testing.T) {
	RegisterType("person", GetTypeInfo(Person{}))
}

func testGetInst1(t *testing.T) {
	v := GetInstance("person").(*Person)
	t.Log(v.Age == 0)
	t.Log(v.Name == "")
	v.Name = "fuck"
	v.Age = 100
}

func testGetInst2(t *testing.T) {
	v := GetInstance("person").(*Person)
	t.Log(v.Age == 100)
	t.Log(v.Name == "fuck")
}

func testGetInst3(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		v := GetInstance("person").(*Person)
		v.Name = "fuck1"
		v.Age = 101
		wg.Done()
	}()

	go func() {
		v := GetInstance("person").(*Person)
		t.Log("-------------------1")
		t.Log(v.Name)
		t.Log(v.Age)
		t.Log("-------------------2")
		wg.Done()
	}()
	wg.Wait()
}

func TestAll(t *testing.T) {
	t.Run("Reg", testReg)
	t.Run("testGetInst1", testGetInst1)
	t.Run("testGetInst2", testGetInst2)
	t.Run("testGetInst3", testGetInst3)
}
