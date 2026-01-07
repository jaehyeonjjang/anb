package breakdown

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnTopcategory
    ColumnSubcategory
    ColumnCategory
    ColumnMethod
    ColumnCount
    ColumnLastdate
    ColumnDuedate
    ColumnRemark
    ColumnElevator
    ColumnPercent
    ColumnRate
    ColumnType
    ColumnDong
    ColumnStandard
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




