package periodicother

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnType
    ColumnResult
    ColumnStatus
    ColumnPosition
    ColumnFilename
    ColumnOfflinefilename
    ColumnChange
    ColumnCategory
    ColumnOrder
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




