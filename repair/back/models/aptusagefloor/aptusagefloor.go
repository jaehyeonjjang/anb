package aptusagefloor

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnFloor
    ColumnPurpose
    ColumnArea
    ColumnRemark
    ColumnOrder
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




