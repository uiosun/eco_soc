package gp

import (
	"github.com/kelindar/goap"
	"strings"
)

// MyAction represents a single action that can be performed by the agent.
type MyAction struct {
	Name     ActionEnum
	ObjectId int
	cost     float32
	require  *goap.State
	outcome  *goap.State
}

type ActionEnum int

func (ae ActionEnum) String() string {
	return ActionEnums[ae]
}

var (
	ActionEnumSleep    = ActionEnum(1)
	ActionEnumEat      = ActionEnum(2)
	ActionEnumForage   = ActionEnum(3)
	ActionEnumSellFood = ActionEnum(4)
	ActionEnumBuyFood  = ActionEnum(5)
	ActionEnums        = map[ActionEnum]string{
		ActionEnumEat:      "小麦",
		ActionEnumSleep:    "休息",
		ActionEnumForage:   "收集食物",
		ActionEnumSellFood: "出售食物",
		ActionEnumBuyFood:  "购买食物",
	}
)

// NewAction creates a new action from the given name, require and outcome.
func NewAction(id int, cost float32, name ActionEnum, require, outcome string) *MyAction {
	return &MyAction{
		Name:     name,
		ObjectId: id,
		cost:     cost,
		require:  goap.StateOf(strings.Split(require, ",")...),
		outcome:  goap.StateOf(strings.Split(outcome, ",")...),
	}
}

// Simulate simulates the action and returns the required and outcome states.
func (a *MyAction) Simulate(current *goap.State) (*goap.State, *goap.State) {
	return a.require, a.outcome
}

// Cost returns the cost of the action.
func (a *MyAction) Cost() float32 {
	return a.cost
}
