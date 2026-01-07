package standardwage

type Column int

const (
    _ Column = iota
    ColumnId
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
    ColumnTechprice1
    ColumnTechprice2
    ColumnTechprice3
    ColumnTechprice4
    ColumnFinancialprice1
    ColumnFinancialprice2
    ColumnFinancialprice3
    ColumnFinancialprice4
    ColumnDirectprice
    ColumnPrintprice1
    ColumnPrintprice2
    ColumnLossprice
    ColumnGasprice
    ColumnTravelprice
    ColumnTravel
    ColumnLoss
    ColumnGas
    ColumnEtc
    ColumnDanger
    ColumnMachine
    ColumnPrint
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




