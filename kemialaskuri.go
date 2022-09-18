package main

// Imports
import (
	"fmt"
	"kemialaskuri/modes"
	"kemialaskuri/utils"
	"strings"
)

// c1, V1 = väkevämpi

const (
	modesText string = "\n\t(1) Väkevästä laimeaksi" +
		"\n\t(2) - Jotain jotain" +
		"\n\t(3) - Taas jotain ja jotain" +
		"\n\n\tEXIT - Sammuttaa sovelluksen" // Exit command
)

// Variables used in app
var (
	syöte string
	exit  bool = false
)

func main() {
	fmt.Println("Tervetuloa hienoon kemialaskuriin!")
	utils.DashRow(40, 'y')

	for exit == false {
		fmt.Print("\nMinkä laskurin haluat?", modesText, "\n>> ")
		fmt.Scanln(&syöte)

		switch strings.ToLower(syöte) {
		case "1":
			modes.Laimennus_mode()
		case "exit":
			exit = true
		default:
			fmt.Printf("\nLaskuria numerolla %s ei löydy!\n", syöte)
			utils.DashRow(40, 'y')
			fmt.Printf("\n\n")
		}
	}

}
