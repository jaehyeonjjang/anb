package comparecompany

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnAddress
    ColumnAddressetc
    ColumnTel
    ColumnFax
    ColumnCeo
    ColumnFormat
    ColumnImage
    ColumnImage2
    ColumnAdjust
    ColumnFinancialprice
    ColumnTechprice
    ColumnDirectprice
    ColumnPrintprice
    ColumnExtraprice
    ColumnTravelprice
    ColumnGasprice
    ColumnDangerprice
    ColumnMachineprice
    ColumnRemark
    ColumnType
    ColumnDefault
    ColumnOrder
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




