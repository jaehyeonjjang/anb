package detailtechnician

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
    ColumnDetail
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




