package periodicdata

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnGroup
    ColumnType
    ColumnPart
    ColumnMember
    ColumnShape
    ColumnWidth
    ColumnLength
    ColumnCount
    ColumnProgress
    ColumnRemark
    ColumnOrder
    ColumnContent
    ColumnStatus
    ColumnFilename
    ColumnOfflinefilename
    ColumnUser
    ColumnBlueprint
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




