package world

import (
	"eco_soc/gp"
	"fmt"
	"math/rand"
)

func (w *World) Turn() {
	for _, human := range w.Humans {
		info := ""

		// 每回合 +1 岁，到 100 岁有概率会死亡
		dieAge := 100
		human.Age++
		if human.Age > dieAge {
			if rand.Intn(dieAge) < 10 {
				info += fmt.Sprintf("%d 岁了，死了。", human.Id)
				delete(w.Humans, human.Id)
				continue
			}
		}

		if len(human.Actions) == 0 {
			info = human.Think()
			human.History = append(human.History, info)
			continue
		}

		action := human.Actions[0]
		human.Actions = human.Actions[1:]
		switch a := action.(type) {
		case *gp.MyAction:
			info += fmt.Sprintf("%d 正在 %s\n", human.Id, a.Name)
			switch a.Name {
			case gp.ActionEnumEat:
				human.Eat()
			case gp.ActionEnumSleep:
				human.Sleep()
			case gp.ActionEnumForage:
				human.Forage()
			}
		}
		human.History = append(human.History, info)
	}
}
