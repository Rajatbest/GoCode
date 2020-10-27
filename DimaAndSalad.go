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
const MAXN = 200000 + 10 
const fi = 100000
func max(a , b int) int{

	if(a > b){
		return a
	}

	return b 
}

func main(){
	scanner = bufio.NewScanner(os.Stdin)
  	scanner.Split(bufio.ScanWords)
 	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n := getI()
	k := getI()
	
	var a,b [100+10]int
	var vis [MAXN]bool 
	var dp,x [MAXN]int
	vis[fi] = true
	for i := 0 ; i < n ; i ++{
		a[i] = getI()
	}

	for i := 0 ; i < n ; i ++{
		b[i] = getI()
	}

	ch := false 

	for i := 0 ; i < n ; i ++{
		tmp := a[i] - k*b[i]
		var v []int
		for j := 0 ; j <= 200000 ; j ++{
			x[j] = 0 
			if vis[j]{
				v = append(v,j)
			}	
		}
		l := len(v)

		for j := 0 ; j < l ; j ++{
			idx := v[j] + tmp
			if idx < 0 {
				continue
			}

			x[idx] = max(x[idx], dp[v[j]] + a[i])

			vis[idx] = true
			if idx == fi {
				ch = true
			}
		 }
		 
		 for j := 0 ; j <= 200000 ; j ++ {
			 dp[j] = max(dp[j],x[j])
		 }
	}

	if ch {
		writer.WriteString(fmt.Sprintf("%v\n",dp[fi]))
	}else{
		writer.WriteString(fmt.Sprintf("-1\n"))
	}
}
