package reviewbasic

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnContent
    ColumnAdjust
    ColumnOrder
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




