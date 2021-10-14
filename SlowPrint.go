package main

import (
	"github.com/01-edu/z01"
	"time"
)

func Slow(s string, v int) {
	if v == 1 {
		for _, letter := range s {
			time.Sleep(35 * time.Millisecond)
			z01.PrintRune(letter)
		}
	} else if v == 2 {
		for _, letter := range s {
			time.Sleep(10 * time.Millisecond)
			z01.PrintRune(letter)
		}
	} else if v == 3 {
		for _, letter := range s {
			time.Sleep(1 * time.Millisecond)
			z01.PrintRune(letter)
		}
	} else if v == 4 {
		for _, letter := range s {
			time.Sleep(100 * time.Millisecond)
			z01.PrintRune(letter)
		}
	} else if v == 5 {
		for _, letter := range s {
			time.Sleep(200 * time.Millisecond)
			z01.PrintRune(letter)
		}
	}
}
