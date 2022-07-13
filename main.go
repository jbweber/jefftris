package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/buger/goterm"
	"golang.org/x/term"
)

var jay = [2][3]string{{"", "", "\033[31mX\033[39m"}, {"\033[31mX\033[39m", "\033[31mX\033[39m", "\033[31mX\033[39m"}}

var jays = [][][]string{
	{{"", "", "\033[31mX\033[39m"}, {"\033[31mX\033[39m", "\033[31mX\033[39m", "\033[31mX\033[39m"}},
	{{"\u001B[31mX\u001B[39m", "\u001B[31mX\u001B[39m"}, {"", "\033[31mX\033[39m"}, {"", "\033[31mX\033[39m"}},
	{{"\033[31mX\033[39m", "", ""}, {"\033[31mX\033[39m", "\033[31mX\033[39m", "\033[31mX\033[39m"}},
	{{"", "\u001B[31mX\u001B[39m"}, {"", "\u001B[31mX\u001B[39m"}, {"\u001B[31mX\u001B[39m", "\u001B[31mX\u001B[39m"}},
}

func main() {
	if term.IsTerminal(int(os.Stdout.Fd())) {
		width, height, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			log.Fatalf("%s", err)
		}

		log.Printf("%d x %d", width, height)
	} else {
		log.Fatal("cannot determine terminal size")
	}

	width := goterm.Width()
	height := goterm.Height()

	for i := 1; i < height-1; i++ {
		drawBox(width, height, i)
		time.Sleep(1 * time.Second)
	}

	time.Sleep(100 * time.Second)
}

func drawBox(width, height, loc int) {
	goterm.MoveCursor(1, 1)
	goterm.Clear()
	goterm.Flush()

	j := 0
	for i := 1; i < height; i++ {
		if i == height-1 {
			goterm.MoveCursor(2, i)
			goterm.Print(strings.Repeat("-", width-2))
		} else {
			goterm.MoveCursor(2, i)
			goterm.Printf("|%d", i)
			if i == loc {
				goterm.MoveCursor(width/2, i)
				//goterm.Print(jay) // strings.Repeat(goterm.Color("X", goterm.BLUE), 6)
				drawLetter(width/2, i, j)
			}
			goterm.MoveCursor(width-1, i)
			goterm.Print("|")
		}
	}

	goterm.Flush()
}

func drawLetter(width, height, j int) {
	goterm.Printf("XXX%dX", j)

	for x := 0; x < len(jays[j]); x++ {
		for y := 0; y < len(jays[j][x]); y++ {
			goterm.MoveCursor(width+x, height+y)
			goterm.Print(jays[j][x][y])
		}
	}
}
