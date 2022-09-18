package main

// Imports
import (
	"fmt"
	"kemialaskuri/modes"
)

// c1, V1 = väkevämpi

const (
	modesText string = "\n\t(1) Väkevästä laimeaksi" +
		"\n\t(2) Jotain jotain\n\t" +
		"(3) Taas jotain ja jotain"
)

// Variables used in app
var (
	syöte string
)

func main() {
	fmt.Println("Tervetuloa hienoon kemialaskuriin!")
	for i := 0; i < 40; i++ {
		fmt.Print("-")
	} // Cool row of lines

	for {
		fmt.Print("\n\nMinkä laskurin haluat?", modesText, "\n>> ")
		fmt.Scanln(&syöte)

		switch syöte {
		case "1":
			modes.Laimennus_mode()
		default:
			fmt.Printf("\nLaskuria numerolla %s ei löydy!\n\n", syöte)
		}
	}

}
