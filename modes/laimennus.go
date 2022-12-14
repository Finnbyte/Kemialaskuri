package modes

import (
	"fmt"
	"kemialaskuri/utils"
	"strconv"
)

var msgs = []string{
	"\n\nAnna c1 (konsentraatio, josta laimennat) (mol/l): ",
	"\nAnna c2 (Päämäärä, miten pienen konsetraation haluat) (mol/l):  ",
	"\nAnna v2 (Kuinka paljon liuosta haluat) (ml/l):  ",
}

func Laimennus_mode() {
	var (
		c1 float64 // Liuoksen konsetraatio, josta laimennat. Alku luku
		c2 float64 // Päämäärä, miten pienen konsetraation haluat.
		v2 float64 // Kuinka paljon lopussa haluat liuosta

		v1 string // Kuinka paljon liuosta tarvitaan että voit tehdä laimennuksen
	)

	utils.AskWithMsgs(&c1, &c2, &v2, msgs)
	// utils.EmptyArgs(c1, c2, v2);

	v1 = strconv.FormatFloat((v2 * c2 / c1), 'f', 1, 64)

	formatted_answer := "V1 = " +
		strconv.FormatFloat(v2, 'f', 1, 64) +
		" ml * " +
		strconv.FormatFloat(c2, 'f', 1, 64) +
		"mol/l / " +
		strconv.FormatFloat(c1, 'f', 1, 64) +
		"mol/l = " +
		v1 + " ml"

	fmt.Printf("\n%s\n", formatted_answer)
}
