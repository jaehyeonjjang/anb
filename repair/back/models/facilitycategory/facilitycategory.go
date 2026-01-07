package facilitycategory

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnOrder
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




