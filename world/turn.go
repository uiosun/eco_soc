package world

func (w *World) Turn() {
	for _, human := range w.Humans {
		human.Work
	}
}
