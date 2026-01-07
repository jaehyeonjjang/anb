package area

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnFamilycount
    ColumnSize
    ColumnOrder
    ColumnRemark
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




