/**
 * ------------------------------------------------------------------------
 * Go's runes are used to represent single characters.
 * Written with single quotation marks.
 * Uses unicode standard for storing runes.
 * ------------------------------------------------------------------------
 */
package main

import "fmt"

func main() {
	fmt.Println('A')
	fmt.Println('B') 
	fmt.Println('ᅴ') // Got this hangul when searching for unicode 1174
	fmt.Println('Җ') // Zhje or Zhe with descender is a letter of the Cyrillic script.
	fmt.Println('\t')
	fmt.Println('\n')
	fmt.Println('\\')
}