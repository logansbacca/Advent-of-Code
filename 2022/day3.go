
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var l rune
	var inThrees []string
	var tot = 0
	alfabeto := "abcdefghijklmnopqrstuvwxyz"
	alfabeto2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(inThrees) < 3 {
			inThrees = append(inThrees, line)
		}

		if len(inThrees) == 3 {
			l = threeLines(inThrees)
			inThrees = nil 
			if unicode.IsLower(l) {
				for i, v := range alfabeto {
					if l == v {
						tot += i + 1
						break
					}
				}
			} else {
				for i, v := range alfabeto2 {
					if l == v {
						tot += 26 + i
						tot++
						break
					}
				}
			}
			
		}
		
	}
	fmt.Println(tot)
}

func spuria(s string) rune {
	var shared rune
	metà := len(s) / 2
	prima := s[:metà]
	seconda := s[metà:]
	for i := 0; i < len(prima); i++ {
		letter := prima[i]
		for j := 0; j < len(seconda); j++ {
			if letter == seconda[j] {
				shared = rune(letter)
				break
			}
		}
	}
	return shared

}

func threeLines(lines []string) (secondRes rune) {
	for _, v := range lines[0] {
		if strings.Contains(lines[1], string(v)) && strings.Contains(lines[2], string(v)) {
			secondRes = v
		}
	}
	return secondRes
}
