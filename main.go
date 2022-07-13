package main

import (
	"strings"
	"time"

	"github.com/buger/goterm"
	"github.com/jbweber/jefftris/internal/terminal"
)

//var jays = [][][]string{
//	{{"", "", terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT}, {terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT, terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT, terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT}},
//	{{terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT, terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT}, {"", terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT}, {"", terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT}},
//	{{terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT, "", ""}, {terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT, terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT, terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT}},
//	{{"", terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT}, {"", terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT}, {terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT, terminal.BACKGROUND_BLUE + "@" + terminal.BACKGROUND_DEFAULT}},
//}

var tees = [][][]string{
	{{"", terminal.MAGENTA_BLOCK}, {terminal.MAGENTA_BLOCK, terminal.MAGENTA_BLOCK}, {"", terminal.MAGENTA_BLOCK}},
	{{"", terminal.MAGENTA_BLOCK, ""}, {terminal.MAGENTA_BLOCK, terminal.MAGENTA_BLOCK, terminal.MAGENTA_BLOCK}},
}

var test = [][][]string{
	//{{"X", "X", "X"}, {"X", "O", "O"}, {"Z", "O", "O"}, {"Z", "Z", "Z"}},
	{{terminal.RED_BLOCK, terminal.YELLOW_BLOCK, terminal.YELLOW_BLOCK, terminal.BLUE_BLOCK}, {terminal.RED_BLOCK, terminal.YELLOW_BLOCK, terminal.YELLOW_BLOCK, terminal.BLUE_BLOCK}, {terminal.RED_BLOCK, terminal.RED_BLOCK, terminal.BLUE_BLOCK, terminal.BLUE_BLOCK}},
}

func main() {
	width := goterm.Width()
	height := goterm.Height()

	j := 0
	for i := 1; i < height-1; i++ {
		drawBox(width, height, i, j, tees)
		j++
		if j >= len(tees) {
			j = 0
		}
		time.Sleep(1 * time.Second)
	}

	time.Sleep(100 * time.Second)
}

func drawBox(width, height, loc, j int, block [][][]string) {
	goterm.MoveCursor(1, 1)
	goterm.Clear()
	goterm.Flush()

	for i := 1; i < height; i++ {
		if i == height-1 {
			goterm.MoveCursor(2, i)
			goterm.Print(strings.Repeat("=", width-2))

			continue
		}

		goterm.MoveCursor(2, i)
		goterm.Printf("%s||%d%s", terminal.FOREGROUND_CYAN, i, terminal.FOREGROUND_DEFAULT)
		if i == loc {
			goterm.MoveCursor(width/2, i)
			//goterm.Print(jay) // strings.Repeat(goterm.Color("X", goterm.BLUE), 6)
			drawLetter(width/2, i, j, block)
		}
		goterm.MoveCursor(width-1, i)
		goterm.Print("||")
	}

	goterm.Flush()
}

func drawLetter(width, height, j int, block [][][]string) {
	for x := 0; x < len(block[j]); x++ {
		for y := 0; y < len(block[j][x]); y++ {
			goterm.MoveCursor(width+x, height+y)
			goterm.Print(block[j][x][y])
		}
	}
}
