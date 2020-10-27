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

func lowerBound(en,k int) int {
	st := 0

	for st <= en{
		mid := (st + en) >> 1
		if DP[mid] < k {
			st = mid + 1
		}else{
			en = mid - 1
		}
	 }
	 
	 return st
}
const MAXN = 100010
var DP [MAXN]int

func main(){

	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n:= getI() 
	
	sol := 0
	
	for i := 0 ; i < n ; i ++ {
		k := getI()
		x := lowerBound(sol,k)
		
		DP[x] = k
		
		if (x > sol){
			sol = x
		}
	}
	writer.WriteString(fmt.Sprintf("%v\n",sol))
}
