package standardhistory

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
    ColumnStandard
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




