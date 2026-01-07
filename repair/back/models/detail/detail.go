package detail

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
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
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




