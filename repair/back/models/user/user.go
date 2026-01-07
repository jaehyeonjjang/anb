package user

type Column int

const (
    _ Column = iota
    ColumnId
    ColumnLoginid
    ColumnPasswd
    ColumnName
    ColumnEmail
    ColumnLevel
    ColumnApt
    ColumnDate

)

type Params struct {
    Column Column
    Value interface{}
}


type Level int

const (
    _ Level  = iota

    LevelNormal
    LevelManager
    LevelAdmin
    LevelRootadmin
)

var Levels = []string{ "", "일반", "매니저", "관리자", "총관리자" }



func GetLevel(value Level) string {
    i := int(value)
    if i <= 0 || i >= len(Levels) {
        return ""
    }
     
    return Levels[i]
}

func FindLevel(value string) Level {
    for i := 1; i < len(Levels); i++ {
        if Levels[i] == value {
            return Level(i)
        }
    }
     
    return 0
}

func ConvertLevel(value []int) []Level {
     items := make([]Level, 0)

     for item := range value {
         items = append(items, Level(item))
     }
     
     return items
}

