package world

import "github.com/pterm/pterm"

// RefreshUI 刷新 UI
func (w *World) RefreshUI() {
	// panel
	panels := pterm.Panels{
		{
			{Data: "This is the first panel"},
		},
		{
			{Data: pterm.Red("This is another\npanel line")},
			{Data: "This is the second panel\nwith a new line"},
		},
	}
	_ = pterm.DefaultPanel.WithPanels(panels).WithPadding(5).Render()
}
