package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func rgb(i int) (int, int, int) {
	var f = 0.1
	//to be not to lame, changed the calc from sin to cosine
	//to violet as a color starts first instead of neon yellow
	return int(math.Cos(f*float64(i)+0)*127 + 128),
		int(math.Cos(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Cos(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func print(output []rune) {
	for j := range output {
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
}

func main() {
	info, _ := os.Stdin.Stat()
	var output []rune

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gorainbow")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	print(output)
}
