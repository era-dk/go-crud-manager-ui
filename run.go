package manager

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	IndexPage = Page{
		Title: "List records",
		Layers: []LayerInterface{
			Listing,
			Create,
			View,
			Edit,
			Delete,
		},
	}

	CreatePage = Page{
		Title: "Create record",
		Layers: []LayerInterface{
			Form,
		},
	}

	ViewPage = Page{
		Title: "View record",
		Layers: []LayerInterface{
			View,
		},
	}

	EditPage = Page{
		Title: "Edit record",
		Layers: []LayerInterface{
			Form,
		},
	}

	CurrentPage *Page = &IndexPage
	SetCurrentPage = func (p *Page) {
		CurrentPage = p
		CurrentPage.Load()
	}
)

func Run() error {
	if CurrentPage == nil {
		return errors.New("page not found")
	}

	CurrentPage.Load()
	_, err := tea.NewProgram(&model{}).Run()
	return err
}