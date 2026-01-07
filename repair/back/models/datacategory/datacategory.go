package datacategory

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnCategory
    ColumnType
    ColumnRemark
    ColumnOrder
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




