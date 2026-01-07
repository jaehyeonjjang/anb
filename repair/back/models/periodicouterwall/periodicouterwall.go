package periodicouterwall

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnResult1
    ColumnResult2
    ColumnResult3
    ColumnResult4
    ColumnResult5
    ColumnResult6
    ColumnStatus1
    ColumnStatus2
    ColumnStatus3
    ColumnStatus4
    ColumnStatus5
    ColumnStatus6
    ColumnPosition1
    ColumnPosition2
    ColumnPosition3
    ColumnPosition4
    ColumnPosition5
    ColumnPosition6
    ColumnContent
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




