package compareestimate

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnType
    ColumnSubtype
    ColumnOriginalprice
    ColumnSaleprice
    ColumnPrice
    ColumnFinancialprice
    ColumnTechprice
    ColumnDirectprice
    ColumnPrintprice
    ColumnExtraprice
    ColumnTravelprice
    ColumnLossprice
    ColumnGasprice
    ColumnEtcprice
    ColumnDangerprice
    ColumnMachineprice
    ColumnCarprice
    ColumnDiscount
    ColumnPerson1
    ColumnPerson2
    ColumnPerson3
    ColumnPerson4
    ColumnPerson5
    ColumnPerson6
    ColumnPerson7
    ColumnPerson8
    ColumnPerson9
    ColumnPerson10
    ColumnPersonprice1
    ColumnPersonprice2
    ColumnPersonprice3
    ColumnPersonprice4
    ColumnPersonprice5
    ColumnPersonprice6
    ColumnPersonprice7
    ColumnPersonprice8
    ColumnPersonprice9
    ColumnPersonprice10
    ColumnTravel
    ColumnLoss
    ColumnGas
    ColumnEtc
    ColumnDanger
    ColumnMachine
    ColumnCar
    ColumnPrint
    ColumnWritedate
    ColumnStart
    ColumnRemark
    ColumnAdjust
    ColumnUser
    ColumnComparecompany
    ColumnEstimate
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




