package periodicdataimage

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnFilename
    ColumnOfflinefilename
    ColumnOrder
    ColumnPeriodicdata
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




