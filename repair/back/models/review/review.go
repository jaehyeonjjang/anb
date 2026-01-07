package review

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnTopcategory
    ColumnSubcategory
    ColumnCategory
    ColumnStandard
    ColumnMethod
    ColumnCycle
    ColumnPercent
    ColumnCount
    ColumnPrice
    ColumnContent
    ColumnAdjust
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




