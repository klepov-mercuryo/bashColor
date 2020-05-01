package main

import (
	"fmt"
	"github.com/evgeny-klyopov/bashColor"
	"strconv"
	"strings"
)

func main() {
	color := bashColor.NewColor()

	printHeader("Default color:", false)
	printColorExample(color)
	printHeader("Modified or replace color color code:", true)
	codes := color.GetColors()
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			code := strconv.Itoa(i * 16 + j)
			codes["Code_" + code] = "\u001b[38;5;" + code + "m"
		}
	}
	color.SetColors(codes)

	for i := 0; i < 256; i++ {
		code := strconv.Itoa(i)
		message := color.GetColor("Code_" +  code) +  fmt.Sprintf("%4v", code)
		if (i+1)%16 == 0 && i > 0 {
			fmt.Println(message)
		} else {
			fmt.Print(message)
		}
	}
}

func printColorExample(color bashColor.Colorer)  {
	fmt.Println("Service messages:")
	color.Print(color.Info, "\tInfo message")
	color.Print(color.Warning, "\tWarning message")
	color.Print(color.Fatal, "\tFatal message")
	color.Print(color.Success, "\tSuccess message")

	fmt.Println("")
	fmt.Println("Color messages:")
	color.Print(color.Black, "\tBlack message")
	color.Print(color.Red, "\tRed message")
	color.Print(color.Green, "\tGreen message")
	color.Print(color.Yellow, "\tYellow message")
	color.Print(color.Purple, "\tPurple message")
	color.Print(color.Magenta, "\tMagenta message")
	color.Print(color.Teal, "\tTeal message")
	color.Print(color.White, "\tWhite message")

	fmt.Println("")
	fmt.Println("Get color code:")
	codes := []string{
		color.GetColor(bashColor.Default) + "Default",
		color.GetColor(bashColor.Black) + "Black",
		color.GetColor(bashColor.Red) + "Red",
		color.GetColor(bashColor.Green) + "Green",
		color.GetColor(bashColor.Yellow) + "Yellow",
		color.GetColor(bashColor.Purple) + "Purple",
		color.GetColor(bashColor.Magenta) + "Magenta",
		color.GetColor(bashColor.Teal) + "Teal",
		color.GetColor(bashColor.White) + "White",
	}
	fmt.Println(strings.Join(codes, " "))
}

func printHeader(message string, separate bool) {
	if separate == true {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("========================================================")
		fmt.Println("")
	}
	fmt.Println(message)
	fmt.Println("")
}