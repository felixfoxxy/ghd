package main

//Code by FelixFoxxy
//Website: https://felixfoxxy.dev/
//Repo: https://github.com/felixfoxxy/ghd

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		run(os.Args[1], 10)
	} else if len(os.Args) == 3 {
		i, err := strconv.Atoi(os.Args[2])
		if err == nil {
			run(os.Args[1], i)
		} else {
			fmt.Println("Invalid Column Size!")
		}
	} else {
		fmt.Println("Usage: ghd [FILE] <COLUMN SIZE>")
	}
}

func run(fn string, lineSz int) {
	bytes, err := os.ReadFile(fn)
	if err == nil {
		printSpc("OFFSET", " ", lineSz)
		fmt.Print("   ")
		printSpc("BYTES", " ", lineSz*3)
		fmt.Print("  ")
		printSpc("DECODED", " ", lineSz)
		fmt.Println()
		fmt.Println()
		i := 0
		for i = 0; i < len(bytes)/lineSz; i++ {
			newRow(bytes, i, lineSz, 0)
			fmt.Println()
		}
		m := len(bytes) % lineSz
		newRow(bytes, i, lineSz, lineSz*3-m*3)
		fmt.Println()
	} else {
		fmt.Println("Error reading File!")
		fmt.Println(err.Error())
	}
}

func newRow(bytes []byte, i int, lineSz int, spacer int) {
	off := strings.ToUpper(fmt.Sprintf("%x", (i * lineSz)))
	spc("0", lineSz-len(off))
	fmt.Print(off)
	fmt.Print(" - ")
	for b := 0; b < lineSz; b++ {
		if b+(i*lineSz) < len(bytes) {
			pnt := strings.ToUpper(hex.EncodeToString([]byte{bytes[b+(i*lineSz)]}))
			fmt.Print(pnt + " ")
		}
	}
	spc(" ", spacer)
	fmt.Print("- ")
	for c := 0; c < lineSz; c++ {
		if c+(i*lineSz) < len(bytes) {
			fmt.Print(strings.Replace(strings.Replace(string(bytes[c+(i*lineSz)]), "\r", "\\r", -1), "\n", "\\n", -1) + " ")
		}
	}
}

func spc(txt string, sz int) {
	for i := 0; i < sz; i++ {
		fmt.Print(txt)
	}
}

func printSpc(txt string, spacer string, lineSz int) {
	fmt.Print(txt)
	for s := len(txt); s < lineSz; s++ {
		fmt.Print(spacer)
	}
}
