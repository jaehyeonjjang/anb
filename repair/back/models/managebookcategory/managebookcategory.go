package managebookcategory

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnOrder
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




