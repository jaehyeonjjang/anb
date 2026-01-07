package aptrepairlist

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnTel
    ColumnFax
    ColumnTestdate
    ColumnEmail
    ColumnPersonalemail
    ColumnZip
    ColumnAddress
    ColumnAddress2
    ColumnCompleteyear
    ColumnType
    ColumnFlatcount
    ColumnFamilycount
    ColumnFloor
    ColumnFmsloginid
    ColumnFmspasswd
    ColumnReportdate

)

type Params struct {
    Column Column
    Value interface{}
}




