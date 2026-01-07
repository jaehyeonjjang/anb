package periodictechnician

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnType
    ColumnPart
    ColumnSignupstartdate
    ColumnSignupenddate
    ColumnRemark
    ColumnOrder
    ColumnTechnician
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




