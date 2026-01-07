package reviewdate

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnYear
    ColumnMonth
    ColumnApt
    ColumnOrder
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




