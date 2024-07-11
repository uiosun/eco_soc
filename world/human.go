package world

import (
	"eco_soc/gp"
	"fmt"
	"github.com/kelindar/goap"
	"math/rand"
)

// Human 自然人
type Human struct {
	Id      int
	Work    Work
	Age     int
	Cash    int
	Hunger  int
	Tired   int
	State   uint // 状态[能采集食物|存在进行中的订单]
	Bag     map[GoodsType]int
	Actions []goap.Action
	History []string
}

func (h *Human) CanCollectFood(demo bool) string {
	if demo || h.State&1 == 1 {
		return "can_collect_food=1"
	}
	return "!can_collect_food"
}

func (h *Human) HasMarketOrder(demo bool) string {
	if demo || h.State&2 == 2 {
		return "has_market_order=1"
	}
	return "!has_market_order"
}

func (h *Human) Think() string {
	goals := []*goap.State{
		goap.StateOf("food>80", h.CanCollectFood(true)),
		goap.StateOf("cash>50"),
	}
	info := ""

	info += fmt.Sprintf("%d 在思考做些什么……", h.Id)
	init := goap.StateOf("hunger=80", "!food", "!tired", h.CanCollectFood(false))
	goal := goals[rand.Intn(len(goals))]
	actions := []goap.Action{
		gp.NewAction(h.Id, 1, gp.ActionEnumEat, "food>0", "hunger-50,food-5"),
		gp.NewAction(h.Id, 1, gp.ActionEnumForage, fmt.Sprintf("tired<50,%s", h.CanCollectFood(true)), "tired+20,food+10,hunger+5"),
		gp.NewAction(h.Id, 2, gp.ActionEnumSleep, "tired>45", "tired-30"),
	}
	plan, err := goap.Plan(init, goal, actions)
	if err != nil {
		if err.Error() == "no plan could be found to reach the goal" {
			info += fmt.Sprintf("%d 不知道怎么完成 %s，放弃了", h.Id, goal)
			return info
		} else {
			panic(err)
		}
	} else {
		info += fmt.Sprintf("%d 决定去 %s", h.Id, goal)
	}
	h.Actions = plan

	return info
}

func (h *Human) Eat() string {
	info := ""

	for goodsType, num := range h.Bag {
		if goodsType == GoodsTypeWheat {
			h.Bag[goodsType] = num - 5
			h.Hunger -= 50
			info += fmt.Sprintf("%d 吃掉了 %s", h.Id, goodsType)
			fmt.Println(h.Id, "吃掉了", goodsType)
			break
		}
	}
	info += fmt.Sprintf("%d 没找到吃的……很可怜", h.Id)

	return info
}

func (h *Human) Sleep() string {
	info := ""

	if h.Tired > 30 {
		h.Tired -= 30
		info += fmt.Sprintf("%d 休息了一下，精神恢复了", h.Id)
	} else {
		info += fmt.Sprintf("%d 躺了一会儿，但并不累，便站了起来", h.Id)
	}

	return info
}

func (h *Human) Forage() string {
	info := ""

	if h.Tired < 50 {
		h.Tired += 20
		h.Hunger += 5
		findCount := 10
		h.Bag[GoodsTypeWheat] += findCount
		info += fmt.Sprintf("%d 找到了 %d 个 %s", h.Id, findCount, GoodsTypeWheat)
	} else {
		info += fmt.Sprintf("%d 想去采集一些食物，但有些累了", h.Id)
	}

	return info
}

func (h *Human) SetMarketOrder() string {
	info := ""

	return info
}
