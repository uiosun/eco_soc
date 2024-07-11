package world

import (
	"fmt"
	"github.com/pterm/pterm"
)

// RefreshUI 刷新 UI
func (w *World) RefreshUI(specialPanel []pterm.Panel) {
	marketCount := 0
	for _, sells := range w.Market {
		marketCount += len(sells)
	}

	//展示 human.History 最新的信息
	humanInfo := ""
	first := true
	for _, human := range w.Humans {
		if !first {
			humanInfo += fmt.Sprintf("\n")
		}
		first = false
		humanInfo += fmt.Sprintf("居民：%s", human.History[len(human.History)-1])
	}

	// panel
	panels := pterm.Panels{
		{{Data: fmt.Sprintf("居民数：%d；市场订单数：%d", len(w.Humans), marketCount)}},
		specialPanel,
		{{Data: humanInfo}},
		//{
		//	{Data: pterm.Red("This is another\npanel line")},
		//	{Data: "humanInfo"},
		//},
	}

	_ = pterm.DefaultPanel.WithPanels(panels).WithPadding(5).Render()
}
