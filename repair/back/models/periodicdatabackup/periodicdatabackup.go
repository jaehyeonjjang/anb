package periodicdatabackup

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnDate
    ColumnBlueprint
    ColumnCount

)

type Params struct {
    Column Column
    Value interface{}
}




