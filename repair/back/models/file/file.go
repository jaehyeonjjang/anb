package file

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnTitle
    ColumnFilename
    ColumnOriginalfilename
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}




