package main

func main(){
	spending := []float64{9.5, 8, 10.2, 7.4}
	fmt.Println(len(spending))
	fmt.Println(cap(spending))
	
	spending[2] = 9.8

	fmt.Println(spending)
}