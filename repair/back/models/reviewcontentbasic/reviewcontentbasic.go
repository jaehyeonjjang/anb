package reviewcontentbasic

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnContent
    ColumnOrder
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




