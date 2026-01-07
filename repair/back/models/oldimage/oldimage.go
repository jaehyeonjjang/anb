package oldimage

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnApt
    ColumnName
    ColumnLevel
    ColumnParent
    ColumnLast
    ColumnTitle
    ColumnType
    ColumnFloortype
    ColumnFilename
    ColumnOrder
    ColumnDate
    ColumnStandard

)

type Params struct {
    Column Column
    Value interface{}
}




