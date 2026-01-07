package advice

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnContent
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




