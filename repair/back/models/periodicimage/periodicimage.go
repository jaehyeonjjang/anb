package periodicimage

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnType
    ColumnFilename
    ColumnOfflinefilename
    ColumnName
    ColumnUse
    ColumnOrder
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




