package main

import (
	"fmt"
	"strings"
)

// IDEAS:
// - Wrap to 80 characters
// - Play music
// - Use a GUI

func main() {
	intro := "On the {{day}} day of Christmas my true love gave to me"
	days := []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	gifts := []string{"a partridge in a pear tree", "two turtle doves", "three french hens", "four calling birds", "five gold rings", "six geese a-laying", "seven swans a-swimming", "eight maids a-milking", "nine ladies dancing", "ten lords a-leaping", "eleven pipers piping", "twelve drummers drumming"}

	for i, day := range days {
		fmt.Println(strings.Replace(intro, "{{day}}", day, -1))
		giftsToday := gifts[:i+1]
		for j := len(giftsToday) - 1; j >= 0; j-- {
			fmt.Print(giftsToday[j])
			if j > 1 {
				fmt.Print(", ")
			} else if j == 1 {
				fmt.Print(" and ")
			} else { // j == 0
				fmt.Print("!")
			}
		}
		fmt.Println()
		fmt.Println()
	}

    
}