package history

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnYear
    ColumnMonth
    ColumnTopcategory
    ColumnSubcategory
    ColumnCategory
    ColumnContent
    ColumnPrice
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




