package main

import "fmt"

func main(){


	var a , b int

	fmt.Scanf("%d%d",&a,&b)

	sol := 0
	c := 0
	for ;a>0; {
		sol += a
		c += a
		a = int(c/b)
		c %= b
	}
	
	fmt.Println(sol)


}
