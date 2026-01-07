package saving

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnYear
    ColumnForward
    ColumnInterest
    ColumnSurplus
    ColumnSaving
    ColumnEtc
    ColumnUse
    ColumnRemark
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




