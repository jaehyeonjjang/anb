package apt

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnCompleteyear
    ColumnFlatcount
    ColumnType
    ColumnFloor
    ColumnFamilycount
    ColumnFamilycount1
    ColumnFamilycount2
    ColumnFamilycount3
    ColumnTel
    ColumnFax
    ColumnEmail
    ColumnPersonalemail
    ColumnPersonalname
    ColumnPersonalhp
    ColumnZip
    ColumnAddress
    ColumnAddress2
    ColumnContracttype
    ColumnContractprice
    ColumnTestdate
    ColumnNexttestdate
    ColumnRepair
    ColumnSafety
    ColumnFault
    ColumnContractdate
    ColumnContractduration
    ColumnInvoice
    ColumnDepositdate
    ColumnFmsloginid
    ColumnFmspasswd
    ColumnFacilitydivision
    ColumnFacilitycategory
    ColumnPosition
    ColumnArea
    ColumnGroundfloor
    ColumnUndergroundfloor
    ColumnUseapproval
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




