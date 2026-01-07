package periodicresult

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnDefect
    ColumnReinforcement
    ColumnRemark
    ColumnPeriodic
    ColumnAptdong
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




