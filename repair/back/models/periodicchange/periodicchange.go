package periodicchange

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnContent1
    ColumnContent2
    ColumnContent3
    ColumnContent4
    ColumnContent5
    ColumnContent6
    ColumnContent7
    ColumnType
    ColumnOrder
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




