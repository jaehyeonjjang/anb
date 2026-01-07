package breakdownhistory

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
    ColumnOriginalcount
    ColumnOriginalprice
    ColumnOriginalduedate
    ColumnTotalcount
    ColumnTotalprice
    ColumnDong
    ColumnStandard
    ColumnBreakdown
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




