package program

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnType
    ColumnVersion
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




