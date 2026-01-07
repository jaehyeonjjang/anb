package periodicopinion

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnGrade
    ColumnContent1
    ColumnContent2
    ColumnContent3
    ColumnContent4
    ColumnCause1
    ColumnCause2
    ColumnCause3
    ColumnCause4
    ColumnCause5
    ColumnCause6
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




