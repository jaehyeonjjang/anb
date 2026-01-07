package repairlist

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
    ColumnRepairid
    ColumnRepairtype
    ColumnReportdate
    ColumnRepairdate
    ColumnInfo1
    ColumnStatus

)

type Params struct {
    Column Column
    Value interface{}
}


type Repairtype int

const (
    _ Repairtype  = iota

    RepairtypeEstablishment
    RepairtypeReview
    RepairtypeOther
)

var Repairtypes = []string{ "", "재수립", "검토조정", "타업체작업" }



func GetRepairtype(value Repairtype) string {
    i := int(value)
    if i <= 0 || i >= len(Repairtypes) {
        return ""
    }
     
    return Repairtypes[i]
}

func FindRepairtype(value string) Repairtype {
    for i := 1; i < len(Repairtypes); i++ {
        if Repairtypes[i] == value {
            return Repairtype(i)
        }
    }
     
    return 0
}

func ConvertRepairtype(value []int) []Repairtype {
     items := make([]Repairtype, 0)

     for item := range value {
         items = append(items, Repairtype(item))
     }
     
     return items
}

