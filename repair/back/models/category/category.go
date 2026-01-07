package category

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnLevel
    ColumnParent
    ColumnCycle
    ColumnPercent
    ColumnUnit
    ColumnElevator
    ColumnRemark
    ColumnOrder
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




