package outline

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnStartyear
    ColumnEndyear
    ColumnStartmonth
    ColumnEndmonth
    ColumnRate
    ColumnPrice
    ColumnRemark
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




