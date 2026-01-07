package patrol

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnLocation
    ColumnContent
    ColumnProcess
    ColumnOpinion
    ColumnStatus
    ColumnUser
    ColumnApt
    ColumnStartdate
    ColumnEnddate
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}


type Status int

const (
    _ Status  = iota

    StatusNormal
    StatusComplete
)

var Statuss = []string{ "", "순찰중", "순찰완료" }



func GetStatus(value Status) string {
    i := int(value)
    if i <= 0 || i >= len(Statuss) {
        return ""
    }
     
    return Statuss[i]
}

func FindStatus(value string) Status {
    for i := 1; i < len(Statuss); i++ {
        if Statuss[i] == value {
            return Status(i)
        }
    }
     
    return 0
}

func ConvertStatus(value []int) []Status {
     items := make([]Status, 0)

     for item := range value {
         items = append(items, Status(item))
     }
     
     return items
}

