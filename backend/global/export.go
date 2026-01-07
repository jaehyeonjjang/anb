package global

type Export interface {
	Save() (string, error)
	Cell(string) string
	CellInt(int)
	CellPrice(int)
	CellImage(string)
	SetHeight(float64)
}
