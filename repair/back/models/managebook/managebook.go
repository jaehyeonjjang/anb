package managebook

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnFilename
    ColumnOrder
    ColumnManagebookcategory
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




