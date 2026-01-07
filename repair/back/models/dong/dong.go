package dong

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnGround
    ColumnUnderground
    ColumnFamilycount
    ColumnParking
    ColumnElevator
    ColumnBasic
    ColumnRemark
    ColumnOrder
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




