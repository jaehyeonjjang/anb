package periodiccheck

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
    ColumnContent8
    ColumnContent9
    ColumnContent10
    ColumnContent11
    ColumnContent12
    ColumnContent13
    ColumnContent14
    ColumnContent15
    ColumnContent16
    ColumnUse1
    ColumnUse2
    ColumnUse3
    ColumnUse4
    ColumnNeed1
    ColumnNeed2
    ColumnNeed3
    ColumnNeed4
    ColumnAptdong
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




