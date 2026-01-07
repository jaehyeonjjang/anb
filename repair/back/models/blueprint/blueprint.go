package blueprint

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnName
    ColumnLevel
    ColumnParent
    ColumnFloortype
    ColumnFilename
    ColumnUpload
    ColumnParentorder
    ColumnOrder
    ColumnOfflinefilename
    ColumnCategory
    ColumnAptdong
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}


type Floortype int

const (
    _ Floortype  = iota

    FloortypeParking
    FloortypeUnderground
    FloortypeGround
    FloortypeTop
    FloortypeRoof
)

var Floortypes = []string{ "", "주차장", "지하", "지상", "옥탑", "지붕" }



func GetFloortype(value Floortype) string {
    i := int(value)
    if i <= 0 || i >= len(Floortypes) {
        return ""
    }
     
    return Floortypes[i]
}

func FindFloortype(value string) Floortype {
    for i := 1; i < len(Floortypes); i++ {
        if Floortypes[i] == value {
            return Floortype(i)
        }
    }
     
    return 0
}

func ConvertFloortype(value []int) []Floortype {
     items := make([]Floortype, 0)

     for item := range value {
         items = append(items, Floortype(item))
     }
     
     return items
}

