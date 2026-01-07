package global

type Hwp struct {
}

type Margin struct {
	Top    int
	Bottom int
	Left   int
	Right  int
}

type Paragraph struct {
	Id   int
	Name string
}

type Border struct {
	Id   int
	Name string
}

type Cell struct {
	Width     int
	Height    int
	Text      string
	Margin    Margin
	Border    Border
	Paragraph Paragraph
}

type Row struct {
	Cells []Cell
}

type Table struct {
	Margin Margin
	Border Border
	Rows   []Row
}

func (p *Row) Add(item Cell) {
	p.Cells = append(p.Cells, item)
}

func (p *Table) Add(item Row) {
	p.Rows = append(p.Rows, item)
}
