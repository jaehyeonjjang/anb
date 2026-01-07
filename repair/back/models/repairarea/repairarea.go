package repairarea

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnFilename
    ColumnLength
    ColumnStandard
    ColumnContent
    ColumnOrder
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




