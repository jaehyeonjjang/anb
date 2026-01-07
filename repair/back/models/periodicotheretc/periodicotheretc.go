package periodicotheretc

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnContent1
    ColumnContent2
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




