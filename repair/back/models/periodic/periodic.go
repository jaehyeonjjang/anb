package periodic

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnAptname
    ColumnTaskrange
    ColumnReportdate
    ColumnStartdate
    ColumnEnddate
    ColumnSupply
    ColumnContract
    ColumnPrice
    ColumnSafetygrade
    ColumnStatus
    ColumnPrestartdate
    ColumnPreenddate
    ColumnResearchstartdate
    ColumnResearchenddate
    ColumnAnalyzestartdate
    ColumnAnalyzeenddate
    ColumnRatingstartdate
    ColumnRatingenddate
    ColumnWritestartdate
    ColumnWriteenddate
    ColumnPrintstartdate
    ColumnPrintenddate
    ColumnBlueprint1
    ColumnBlueprint2
    ColumnBlueprint3
    ColumnBlueprint4
    ColumnBlueprint5
    ColumnBlueprint6
    ColumnBlueprint7
    ColumnBlueprint8
    ColumnBlueprint9
    ColumnBlueprint10
    ColumnBlueprint11
    ColumnBlueprint1save
    ColumnOwner
    ColumnManager
    ColumnAgent
    ColumnResult1
    ColumnResult2
    ColumnResult3
    ColumnResult4
    ColumnResult5
    ColumnResulttext1
    ColumnResulttext2
    ColumnResulttext3
    ColumnResulttext4
    ColumnResulttext5
    ColumnPast
    ColumnCategory
    ColumnUser
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




