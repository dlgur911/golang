package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func NewActor(name string, hp int, speed float64) *Actor {
	/*
		내가 한 것
		var actor = &Actor{}
		actor.Name = name
		actor.HP = hp
		actor.Speed = speed
		return actor
	*/
	// return &Actor{name, hp, speed}
	/*
		var a = Actor{name, hp, speed}
		return &a
	*/
}

func main() {
	var data Data
	data.value = 1
	data.data[100] = 123

	ChangeData(&data)
	fmt.Printf("value: %d, data[100]: %d", data.value, data.data[100])
	fmt.Println()

	var p *Data = &data
	p.value = 2
	p.data[199] = 200
	//ChangeData(p)
	fmt.Printf("value: %d, data[100]: %d", data.value, data.data[199])
	fmt.Println()

	var actor = NewActor("금토끼", 99, 100)
	fmt.Println(actor.Speed)
	fmt.Println(actor.Name)
}
