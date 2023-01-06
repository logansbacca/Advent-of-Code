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
	mp := make(map[string]int)
	mp["A"] = 1
	mp["X"] = 1
	mp["B"] = 2
	mp["Y"] = 2
	mp["C"] = 3
	mp["Z"] = 3

	for i := 0; i < (j / 2); i++ {
		mine := mp[myscore[i]]
		their := mp[theyscore[i]]

		if their == 1 && mine == 3 || their == 3 && mine == 2 || their == 2 && mine == 1 {
			total += mine
		} else if mine == their {
			total += mine
			total += 3
		} else {
			total += 6
			total += mine
		}

	}

	part2(myscore, theyscore, j)
}

func part2(me []string, them []string, j int) {

	mp := make(map[string]int)
	mp["A"] = 1
	mp["X"] = 1
	mp["B"] = 2
	mp["Y"] = 2
	mp["C"] = 3
	mp["Z"] = 3

	var total int

	for i := 0; i < j/2; i++ {
		mine := mp[me[i]]
		their := mp[them[i]]

		//draw
		if mine == 2 {
			total += their
			total += 3
		}

		//Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
		//A for Rock, B for Paper, and C for Scissors.

		/* In the first round, your opponent will choose Rock (A), and you need the round to end in a draw (Y), so you also choose Rock. This gives you a score of 1 + 3 = 4.
		In the second round, your opponent will choose Paper (B), and you choose Rock so you lose (X) with a score of 1 + 0 = 1.
		In the third round, you will defeat your opponent's Scissors with Rock for a score of 1 + 6 = 7. */

		//loose
		if mine == 1 {
			if them[i] == "A" {
				total += 3
			}
			if them[i] == "B" {
				total += 1
			}
			if them[i] == "C" {
				total += 2
			}
		}

		//win

		//Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
		//A for Rock, B for Paper, and C for Scissors.
		if mine == 3 {
			if them[i] == "A" {
				total += 2
				total += 6
			}
			if them[i] == "B" {
				total += 3
				total += 6
			}
			if them[i] == "C" {
				total += 1
				total += 6
			}
		}

	}
	fmt.Println(total)
}
