package manager

type Page struct {
	Title string
	Layers []LayerInterface
}

func (p *Page) Load() {
	for _, layer := range p.Layers {
		layer.Load()
	}
}
