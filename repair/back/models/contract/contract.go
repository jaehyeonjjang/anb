package contract

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnType
    ColumnContractdate
    ColumnContractstartdate
    ColumnContractenddate
    ColumnPrice
    ColumnVat
    ColumnInvoice
    ColumnDepositdate
    ColumnRemark
    ColumnUser
    ColumnEstimate
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




