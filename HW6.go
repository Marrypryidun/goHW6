package main

import (
	"fmt"
	"sync"
)

type MyStruct struct {
	v   map[int]worker
	mux sync.RWMutex
}
type worker struct {
	person people
	position string
}
type people struct {
	name string
	age int
}

func ToPeople(w worker)people{

	return w.person
}
func (p *MyStruct) createAndPrintBosses(i int,w *worker,waitgroup *sync.WaitGroup) {
	p.mux.Lock()
	w.person.age++
	p.v[i]=*w
	println(i,"-",p.v[i].person.name,p.v[i].person.age)
	p.mux.Unlock()
	waitgroup.Done()
}
func (p *MyStruct) createAndPrint(i int,w *worker) {
	p.mux.Lock()
	w.person.age++
	p.v[i]=*w
	println(i,"-",p.v[i].person.name,p.v[i].person.age)
	p.mux.Unlock()

}
func main()  {
	//конфертація типів
	var w =worker{
		person:
			people{
			name: "Marry",
			age: 18},
			position: "student",
	}
	var w2 =worker{
		person:
		people{
			name: "Jim",
			age: 22},
		position: "developer",
	}
	b:=ToPeople(w);
	/*b := *(*people)(unsafe.Pointer(&w))//можливий і такий варіант*/
	fmt.Println(b)

	// перебір начальників і робітників
	bosses:=MyStruct{v:make(map[int]worker)}
	workers:=MyStruct{v:make(map[int]worker)}
	var waitgroup sync.WaitGroup
	waitgroup.Add(9)
	for i:=1; i<10; i++{
		go bosses.createAndPrintBosses(i,&w2,&waitgroup)
	}
	waitgroup.Wait()
	for i:=1; i<10; i++{
		workers.createAndPrint(i,&w)
	}

}
