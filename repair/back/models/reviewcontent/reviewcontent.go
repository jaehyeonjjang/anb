package reviewcontent

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnContent
    ColumnOrder
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




