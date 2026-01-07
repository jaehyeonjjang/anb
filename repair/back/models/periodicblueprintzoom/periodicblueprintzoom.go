package periodicblueprintzoom

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnIconzoom
    ColumnNumberzoom
    ColumnCrackzoom
    ColumnZoom
    ColumnStatus
    ColumnBlueprint
    ColumnPeriodic
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




