package standard

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnDirect
    ColumnLabor
    ColumnCost
    ColumnUnit
    ColumnOrder
    ColumnOriginal
    ColumnCategory
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




