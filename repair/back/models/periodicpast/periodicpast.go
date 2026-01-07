package periodicpast

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnType
    ColumnCompany
    ColumnName
    ColumnRepairstartdate
    ColumnRepairenddate
    ColumnContent
    ColumnGrade
    ColumnOrder
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




