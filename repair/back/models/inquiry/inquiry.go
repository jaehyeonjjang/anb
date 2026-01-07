package inquiry

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnType
    ColumnContent
    ColumnStatus
    ColumnApt
    ColumnUser
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




