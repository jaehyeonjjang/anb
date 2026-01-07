package periodicincidental

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnResult1
    ColumnResult2
    ColumnResult3
    ColumnResult4
    ColumnResult5
    ColumnResult6
    ColumnResult7
    ColumnResult8
    ColumnResult9
    ColumnResult10
    ColumnResult11
    ColumnResult12
    ColumnResult13
    ColumnResult14
    ColumnResult15
    ColumnResult16
    ColumnResult17
    ColumnResult18
    ColumnResult19
    ColumnResult20
    ColumnResult21
    ColumnStatus1
    ColumnStatus2
    ColumnStatus3
    ColumnStatus4
    ColumnStatus5
    ColumnStatus6
    ColumnStatus7
    ColumnStatus8
    ColumnStatus9
    ColumnStatus10
    ColumnStatus11
    ColumnStatus12
    ColumnStatus13
    ColumnStatus14
    ColumnStatus15
    ColumnStatus16
    ColumnStatus17
    ColumnStatus18
    ColumnStatus19
    ColumnStatus20
    ColumnStatus21
    ColumnPosition1
    ColumnPosition2
    ColumnPosition3
    ColumnPosition4
    ColumnPosition5
    ColumnPosition6
    ColumnPosition7
    ColumnPosition8
    ColumnPosition9
    ColumnPosition10
    ColumnPosition11
    ColumnPosition12
    ColumnPosition13
    ColumnPosition14
    ColumnPosition15
    ColumnPosition16
    ColumnPosition17
    ColumnPosition18
    ColumnPosition19
    ColumnPosition20
    ColumnPosition21
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




