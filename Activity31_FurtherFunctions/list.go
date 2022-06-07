package main

type list []*game

func (l list) print(){
	for _, item := range l{
		item.print()
	}
}