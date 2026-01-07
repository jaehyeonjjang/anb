package patrolimage

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnFilename
    ColumnPatrol
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




