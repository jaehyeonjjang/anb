package olddata

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnApt
    ColumnImage
    ColumnImagetype
    ColumnUser
    ColumnType
    ColumnX
    ColumnY
    ColumnPoint
    ColumnNumber
    ColumnGroup
    ColumnName
    ColumnFault
    ColumnContent
    ColumnWidth
    ColumnLength
    ColumnCount
    ColumnProgress
    ColumnRemark
    ColumnImagename
    ColumnFilename
    ColumnMemo
    ColumnReport
    ColumnUsermemo
    ColumnAptmemo
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




