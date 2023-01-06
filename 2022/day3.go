/* input: lista di stringhe di lunghezza pari, una per riga, del tipo:
MVWpzTTrTFNNLtssjV
Per ogni riga di input, determinare l'unica lettera "spuria", cioè presente sia nella prima che nella seconda metà della stringa  (in ogni riga di input ce n'è sempre e solo una).
A ciascuna lettera spuria è associato un "valore":

se minuscola, dalla 'a' alla 'z', un valore da 1 ('a') a 26 ('z')
se maiuscola, dalla 'A' alla 'Z', un valore da 27 ('A') a 52 ('Z').

D 3-1. Qual è la somma dei valori delle lettere "spurie" trovate?
Ora, per qualche motivo, le stringhe in input sono da considerare in gruppi di 3 (la 1, 2, 3; la 4, 5, 6; ecc.).
Ogni gruppo di tre stringhe ha una lettera "speciale" che è l'unica presente in tutte e tre le stringe. Ogni altra lettera sarà presente al massimo in due delle tre stringhe di uno stesso gruppo.
D 3-2. Qual è la somma dei valori delle lettere "speciali" trovate?
Vale la stessa attribuzione di valori della domanda 3-1. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var tot int
	alfabeto := "abcdefghijklmnopqrstuvwxyz"
	alfabeto2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		l := spuria(line)
		if unicode.IsLower(l) {
			for i, v := range alfabeto {
				if l == v {
					tot += i + 1
					break
				}
			}
		}else {
			for i, v := range alfabeto2 {
				if l == v {
					tot += 26 + i
					tot++
					break
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
