package totalreport

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnApt
    ColumnTopcategory
    ColumnSubcategory
    ColumnCategory
    ColumnStandard
    ColumnMethod
    ColumnRate
    ColumnLastdate
    ColumnDuedate
    ColumnCount
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




