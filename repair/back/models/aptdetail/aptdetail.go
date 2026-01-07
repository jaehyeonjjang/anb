package aptdetail

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnCurrent1
    ColumnCurrent2
    ColumnCurrent3
    ColumnCurrent4
    ColumnCurrent5
    ColumnCurrent6
    ColumnCurrent7
    ColumnCurrent8
    ColumnCurrent9
    ColumnCurrent10
    ColumnCurrent11
    ColumnCurrent12
    ColumnCurrent13
    ColumnCurrent14
    ColumnCurrent15
    ColumnCurrent16
    ColumnCurrent17
    ColumnCurrent18
    ColumnCurrent19
    ColumnCurrent20
    ColumnCurrent21
    ColumnCurrent22
    ColumnCurrent23
    ColumnOutline1
    ColumnOutline2
    ColumnOutline3
    ColumnOutline4
    ColumnOutline5
    ColumnOutline6
    ColumnOutline7
    ColumnOutline8
    ColumnOutline9
    ColumnRecord1
    ColumnRecord2
    ColumnRecord3
    ColumnRecord4
    ColumnRecord5
    ColumnDeligate
    ColumnFacilitydivision
    ColumnFacilitytype
    ColumnFacilitycategory
    ColumnStruct1
    ColumnStruct2
    ColumnStruct3
    ColumnStruct4
    ColumnStruct5
    ColumnStruct6
    ColumnStruct7
    ColumnStruct8
    ColumnStruct9
    ColumnStruct10
    ColumnStruct11
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




