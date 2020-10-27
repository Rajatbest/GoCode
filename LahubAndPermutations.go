package main

import(
	"os"
	"bufio"
	"strconv"
	"fmt"
)

var scanner *bufio.Scanner

func getI64() int64 {
	scanner.Scan()
	x, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return x
}
func getI() int {
	return int(getI64())
}
func getF() float64 {
	scanner.Scan()
	x, _ := strconv.ParseFloat(scanner.Text(), 64)
	return x
}
func getS() string {
	scanner.Scan()
	return scanner.Text()
}

const MOD = int64(1000000007)
const MAXN = 2000 + 10 


func main(){

	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	n := getI()

	var sol int64
	var a [MAXN]int
	var cant [MAXN]bool 
	for i := 1 ; i <= n ; i ++ {

		a[i] = getI()
	}
	wang := 0
	for i := 1 ; i <= n ; i ++{

		if a[i] != -1 {
			cant[a[i]] = true 
			cant[i] = true
		}else{
			wang ++
		}
		
		if a[i] == i {
			fmt.Println(0)
			return 
		}

	}
	
	tmp := 0
	for i := 1 ; i <= n ; i ++ {
		if !cant[i] {
			tmp ++
		}
	}

	var C [MAXN][MAXN]int64
	C[0][0] = int64(1)
	for i := 1 ; i < MAXN ; i ++ {
	//	fmt.Println("-----",i)
		for j := 0 ; j <= i ; j ++ {
		//	fmt.Println("!!",j)
			if j == 0 || j == i {
				C[i][j] = int64(1)
			} else {
				C[i][j] = (C[i-1][j-1] + C[i-1][j] )%MOD
			}
		}
	} 

	var fac [MAXN]int64

	fac[0] = int64(1) 

	for i := 1 ; i < MAXN ; i ++{
		fac[i] = fac[i-1]* int64(i)
		fac[i] %= MOD 
	}

	for i := 0 ; i <= tmp ; i ++{
		if i%2 == 0 {
			sol += C[tmp][i] * fac[wang-i]
		} else {
			sol -= C[tmp][i] * fac[wang-i]
		}
		sol %= MOD

	}
	sol = ( sol+MOD ) %MOD

	fmt.Println(sol)
}
