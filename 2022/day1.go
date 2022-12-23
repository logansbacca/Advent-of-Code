package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main () {

scanner := bufio.NewScanner(os.Stdin)
var currCount int
var maxes [] int
var maxCount int

for scanner.Scan() {
    line := scanner.Text()
    
	if line == "" {
		maxes = append(maxes,currCount)
        currCount = 0
    }
	num,_ := strconv.Atoi(line)
	currCount += num
	if maxCount <= currCount {
		maxCount = currCount
		
	}
    
}

fmt.Println(maxCount)

sort.Ints(maxes)

newArr := maxes[len(maxes)-3:]

var tot int
for _,v := range newArr {
	tot += v
}

fmt.Println("totale :",tot)

}