package technician

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnGrade
    ColumnStamp
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




