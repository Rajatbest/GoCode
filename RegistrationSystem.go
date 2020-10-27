package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){

	var n int 
	var m map[string]int 
	r := bufio.NewReader(os.Stdin)
	m = make(map[string]int)

	fmt.Fscanf(r, "%d\n", &n)

	var s string
	for i := 0 ; i < n ; i ++ {

		

		fmt.Fscanf(r,"%s\n",&s);
		
		if m[s] == 0 {
			fmt.Printf("OK\n")
		} else {
			fmt.Printf("%s%d\n", s, m[s])
		}
		m[s] ++;
		
		
	}


}
