package repair

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnType
    ColumnStatus
    ColumnCalculatetype
    ColumnProvision
    ColumnComplex1
    ColumnComplex2
    ColumnCompletionyear
    ColumnCompletionmonth
    ColumnCompletionday
    ColumnParcelrate
    ColumnPlanyears
    ColumnInfo1
    ColumnInfo2
    ColumnInfo3
    ColumnInfo4
    ColumnInfo5
    ColumnInfo6
    ColumnInfo7
    ColumnInfo8
    ColumnInfo9
    ColumnInfo10
    ColumnInfo11
    ColumnStructure1
    ColumnStructure2
    ColumnStructure3
    ColumnStructure4
    ColumnStructure5
    ColumnStructure6
    ColumnStructure7
    ColumnStructure8
    ColumnStructure9
    ColumnStructure10
    ColumnStructure11
    ColumnStructure12
    ColumnStructure13
    ColumnStructure14
    ColumnReviewcontent1
    ColumnReviewcontent2
    ColumnReviewcontent3
    ColumnReviewcontent4
    ColumnReviewcontent5
    ColumnReviewcontent6
    ColumnReviewcontent7
    ColumnSavingprice
    ColumnPrice1
    ColumnPrice2
    ColumnPrice3
    ColumnPrice4
    ColumnPrice5
    ColumnReportdate
    ColumnContent1
    ColumnContent2
    ColumnContent3
    ColumnContent4
    ColumnContent5
    ColumnContent6
    ColumnContent7
    ColumnContent8
    ColumnContent9
    ColumnContent10
    ColumnContent11
    ColumnPeriodtype
    ColumnRemark
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}


type Periodtype int

const (
    _ Periodtype  = iota

    PeriodtypePeriodic
    PeriodtypeSometime
)

var Periodtypes = []string{ "", "정기", "수시" }



func GetPeriodtype(value Periodtype) string {
    i := int(value)
    if i <= 0 || i >= len(Periodtypes) {
        return ""
    }
     
    return Periodtypes[i]
}

func FindPeriodtype(value string) Periodtype {
    for i := 1; i < len(Periodtypes); i++ {
        if Periodtypes[i] == value {
            return Periodtype(i)
        }
    }
     
    return 0
}

func ConvertPeriodtype(value []int) []Periodtype {
     items := make([]Periodtype, 0)

     for item := range value {
         items = append(items, Periodtype(item))
     }
     
     return items
}

