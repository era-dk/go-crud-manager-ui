package manager

import (
	"bytes"
	"fmt"
	"math"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	Grid = &GridLayer{
		Limit: 10,
		FetchFn: nil,
		Columns: []table.Column{},
		currPage: 1,
		lastPage: 1,
		totalSize: 0,
		tableModel: createTableModel(),
	}
)

type GridLayer struct {
	Limit int
	FetchFn func (limit int, offset int) (RecordCollection, int, error)
	Columns []table.Column
	currPage int
	lastPage int
	totalSize int
	ids map[int]int
	tableModel table.Model
	err error
}

func (l *GridLayer) AddColumn(title string, width int) *GridLayer {
	l.Columns = append(l.Columns, table.Column{Title: title, Width: width})
	return l
}

func (l GridLayer) HasRows() bool {
	return len(l.tableModel.Rows()) > 0
}

func (l GridLayer) FocusedId() int {
	cursor := l.tableModel.Cursor()
	return l.ids[cursor]
}

/* LayerInterface */
func (l *GridLayer) Load() {
	l.load()
}

/* LayerInterface */
func (l *GridLayer) Watch(msg tea.Msg) tea.Cmd {
	rowsExists := l.HasRows()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyLeft:
			if rowsExists && l.currPage > 1 {
				l.currPage--
				l.load()
			}
			return BreakCmd
		case tea.KeyRight:
			if rowsExists && l.currPage < l.lastPage {
				l.currPage++
				l.load()
			}
			return BreakCmd
		}
	}

	l.tableModel, _ = l.tableModel.Update(msg)
	return nil
}

/* LayerInterface */
func (l *GridLayer) RenderBody() string {
	if len(l.Columns) < 1 {
		return ErrorStyle.Render("columns not found")
	}

	var buffer bytes.Buffer

	buffer.WriteString(l.renderToolbar())
	buffer.WriteString(TableBaseStyle.Render(l.tableModel.View()))
	buffer.WriteString("\n")

	if l.err != nil {
		buffer.WriteString(ErrorStyle.Render(l.err.Error()))
		buffer.WriteString("\n")
	}

	return buffer.String()
}

/* LayerInterface */
func (l GridLayer) Help() []HelpCmd {
	if l.totalSize > l.Limit {
		return []HelpCmd{
			{Label: "Pager", Cmd: "left / right"},
		}
	}
	return []HelpCmd{}
}

func (l GridLayer) renderToolbar() string {
	leftPagerSymbol := "<"
	rightPagerSymbol := ">"

	total := PagerActiveStyle.Render(fmt.Sprintf("Total: %d", l.totalSize))
	if !l.HasRows() {
		return fmt.Sprintf(
			"\n%s\n",
			total,
		)
	}

	if l.currPage > 1 {
		leftPagerSymbol = PagerActiveStyle.Render(leftPagerSymbol)
	} else {
		leftPagerSymbol = PagerDisabledStyle.Render(leftPagerSymbol)
	}

	pages := PagerActiveStyle.Render(fmt.Sprintf("Page: %d / %d", l.currPage, l.lastPage))

	if l.currPage < l.lastPage {
		rightPagerSymbol = PagerActiveStyle.Render(rightPagerSymbol)
	} else {
		rightPagerSymbol = PagerDisabledStyle.Render(rightPagerSymbol)
	}

	return fmt.Sprintf(
		"\n%s %s %s %s\n",
		total,
		leftPagerSymbol,
		pages,
		rightPagerSymbol,
	)
}

func (l *GridLayer) load() {
	var (
		records RecordCollection
		totalItems int
	)

	if l.FetchFn != nil {
		offset := (l.currPage - 1) * l.Limit
		records, totalItems, l.err = l.FetchFn(l.Limit, offset)
		if l.err == nil && l.currPage > 1 && records.Len() == 0 {
			l.currPage--
			l.load()
			return
		}
	}
	l.totalSize = totalItems
	l.lastPage = int(math.Ceil(float64(totalItems) / float64(l.Limit)))

	rowId := 0
	l.ids = map[int]int{}
	tableRows := []table.Row{}
	for _, record := range records {
		row := []string{}
		for _, column := range l.Columns {
			row = append(row, record.GetString(column.Title))
		}
		tableRows = append(tableRows, row)
		l.ids[rowId] = record.ID
		rowId++
	}

	l.tableModel.SetColumns(l.Columns)
	l.tableModel.SetHeight(len(tableRows) + 2)
	l.tableModel.SetRows(tableRows)
	l.tableModel.SetCursor(0)
}

func createTableModel() table.Model {
	t := table.New(
		table.WithFocused(true),
		table.WithHeight(2),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}