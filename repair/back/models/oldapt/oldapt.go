package oldapt

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnAptgroup
    ColumnName
    ColumnWorkstartdate
    ColumnWorkenddate
    ColumnType
    ColumnMaster
    ColumnStatus
    ColumnCompany
    ColumnReport
    ColumnReport1
    ColumnReport2
    ColumnReport3
    ColumnReport4
    ColumnReport5
    ColumnReport6
    ColumnSummarytype
    ColumnSearch
    ColumnUser
    ColumnUpdateuser
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




