package main

import (
	"eco_soc/world"
	"github.com/pterm/pterm"
)

/*
PTerm: https://github.com/pterm/pterm?tab=readme-ov-file
*/
func main() {
	w := world.World{}
	w.Init()

	for {
		// input with single line
		result, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue("输入你的选项，'turn' 可用于进入下一回合").Show()
		pterm.Println()

		switch result {
		case "turn":
			w.Turn()
		case "exit":
			fallthrough
		case "E":
			pterm.Println("Bey~")
			return
		}

		w.RefreshUI()
	}
}
