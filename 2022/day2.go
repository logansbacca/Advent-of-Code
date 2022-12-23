package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var myscore []string
	var theyscore []string
	mp := make(map[string]int) //4, 6, 2, 6, 2
	mp["A"] = 1
	mp["X"] = 1
	mp["B"] = 2
	mp["Y"] = 2
	mp["C"] = 3
	mp["Z"] = 3

	j := 0
	for scanner.Scan() {
		curr := scanner.Text()
		if j%2 != 0 {
			myscore = append(myscore, curr)
		} else {
			theyscore = append(theyscore, curr)
		}
		j++
	}
	var total int
	var added int
	for i := 0; i < (j / 2); i++ {
		mine := mp[myscore[i]]
		their := mp[theyscore[i]]
		if mine > their {
			added+= 6
			added += mp[myscore[i]]
		} else if mine == their {
			added += 3
			added += mp[myscore[i]]
		} else {
			added += mp[myscore[i]]
		}
		fmt.Println(added)
		total += added
		added = 0
	}
	fmt.Print(total)
}