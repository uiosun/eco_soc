package main

import (
	"eco_soc/world"
	"fmt"
	"github.com/pterm/pterm"
	"strconv"
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
		specialPanel := []pterm.Panel{}

		switch result {
		case "human", "H":
			result, _ = pterm.DefaultInteractiveTextInput.WithDefaultValue("输入居民的ID，获得详细信息。不存在则直接返回").Show()
			atoi, err := strconv.Atoi(result)
			if err == nil {
				if human, ok := w.Humans[atoi]; ok {
					specialPanel = append(specialPanel, pterm.Panel{Data: fmt.Sprintf("居民：%d（年龄：%d；饥饿：%d；疲惫：%d；现金：%d）", human.Id, human.Age, human.Hunger, human.Tired, human.Cash)})
				}
			} else {
				pterm.Red(err)
			}
		case "turn", "T":
			w.Turn()
		case "exit", "E":
			pterm.Println("Bey~")
			return
		}

		w.RefreshUI(specialPanel)
	}
}
