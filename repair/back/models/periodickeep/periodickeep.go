package periodickeep

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnStatus1
    ColumnStatus2
    ColumnStatus3
    ColumnStatus4
    ColumnStatus5
    ColumnStatus6
    ColumnContent1
    ColumnContent2
    ColumnContent3
    ColumnContent4
    ColumnContent5
    ColumnContent6
    ColumnRemark1
    ColumnRemark2
    ColumnRemark3
    ColumnRemark4
    ColumnRemark5
    ColumnRemark6
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




