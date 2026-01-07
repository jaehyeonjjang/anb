package approval

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnDuty
    ColumnName
    ColumnOrder
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




