package aptdongetc

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnFloortype
    ColumnParent
    ColumnOrder
    ColumnAptdong
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




