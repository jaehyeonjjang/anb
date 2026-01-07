package aptdong

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnDong
    ColumnType
    ColumnGroundcount
    ColumnUndergroundcount
    ColumnParkingcount
    ColumnTopcount
    ColumnRoofcount
    ColumnFamilycount
    ColumnArea
    ColumnRemark
    ColumnPrivate
    ColumnOrder
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




