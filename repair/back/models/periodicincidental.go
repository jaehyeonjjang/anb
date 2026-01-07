package models

import (
    "repair/models/periodicincidental"
    "database/sql"
    "errors"
    "fmt"
    "strings"
    "time"

    "repair/global/config"
    log "repair/global/log"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"

    
)

type Periodicincidental struct {
            
    Id                int64 `json:"id"`         
    Result1                string `json:"result1"`         
    Result2                string `json:"result2"`         
    Result3                string `json:"result3"`         
    Result4                string `json:"result4"`         
    Result5                string `json:"result5"`         
    Result6                string `json:"result6"`         
    Result7                string `json:"result7"`         
    Result8                string `json:"result8"`         
    Result9                string `json:"result9"`         
    Result10                string `json:"result10"`         
    Result11                string `json:"result11"`         
    Result12                string `json:"result12"`         
    Result13                string `json:"result13"`         
    Result14                string `json:"result14"`         
    Result15                string `json:"result15"`         
    Result16                string `json:"result16"`         
    Result17                string `json:"result17"`         
    Result18                string `json:"result18"`         
    Result19                string `json:"result19"`         
    Result20                string `json:"result20"`         
    Result21                string `json:"result21"`         
    Status1                string `json:"status1"`         
    Status2                string `json:"status2"`         
    Status3                string `json:"status3"`         
    Status4                string `json:"status4"`         
    Status5                string `json:"status5"`         
    Status6                string `json:"status6"`         
    Status7                string `json:"status7"`         
    Status8                string `json:"status8"`         
    Status9                string `json:"status9"`         
    Status10                string `json:"status10"`         
    Status11                string `json:"status11"`         
    Status12                string `json:"status12"`         
    Status13                string `json:"status13"`         
    Status14                string `json:"status14"`         
    Status15                string `json:"status15"`         
    Status16                string `json:"status16"`         
    Status17                string `json:"status17"`         
    Status18                string `json:"status18"`         
    Status19                string `json:"status19"`         
    Status20                string `json:"status20"`         
    Status21                string `json:"status21"`         
    Position1                string `json:"position1"`         
    Position2                string `json:"position2"`         
    Position3                string `json:"position3"`         
    Position4                string `json:"position4"`         
    Position5                string `json:"position5"`         
    Position6                string `json:"position6"`         
    Position7                string `json:"position7"`         
    Position8                string `json:"position8"`         
    Position9                string `json:"position9"`         
    Position10                string `json:"position10"`         
    Position11                string `json:"position11"`         
    Position12                string `json:"position12"`         
    Position13                string `json:"position13"`         
    Position14                string `json:"position14"`         
    Position15                string `json:"position15"`         
    Position16                string `json:"position16"`         
    Position17                string `json:"position17"`         
    Position18                string `json:"position18"`         
    Position19                string `json:"position19"`         
    Position20                string `json:"position20"`         
    Position21                string `json:"position21"`         
    Periodic                int64 `json:"periodic"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type PeriodicincidentalManager struct {
    Conn    *Connection
    Result  *sql.Result
    Index   string
    Isolation   bool
    SelectQuery  string
    JoinQuery string
    CountQuery   string
    GroupQuery string
    SelectLog bool
    Log bool
}



func (c *Periodicincidental) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewPeriodicincidentalManager(conn *Connection) *PeriodicincidentalManager {
    var item PeriodicincidentalManager


    if conn == nil {
        item.Conn = NewConnection()
        item.Isolation = false
    } else {
        item.Conn = conn 
        item.Isolation = conn.Isolation
    }

    item.Index = ""
    item.SelectLog = config.Log.Database
    item.Log = config.Log.Database

    return &item
}

func (p *PeriodicincidentalManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *PeriodicincidentalManager) SetIndex(index string) {
    p.Index = index
}

func (p *PeriodicincidentalManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *PeriodicincidentalManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *PeriodicincidentalManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *PeriodicincidentalManager) Query(query string, params ...any) (*sql.Rows, error) {
    if p.Isolation == true {
        query += " for update"
    }

    if p.SelectLog {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Query(query, params...)
}

func (p *PeriodicincidentalManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select pi_id, pi_result1, pi_result2, pi_result3, pi_result4, pi_result5, pi_result6, pi_result7, pi_result8, pi_result9, pi_result10, pi_result11, pi_result12, pi_result13, pi_result14, pi_result15, pi_result16, pi_result17, pi_result18, pi_result19, pi_result20, pi_result21, pi_status1, pi_status2, pi_status3, pi_status4, pi_status5, pi_status6, pi_status7, pi_status8, pi_status9, pi_status10, pi_status11, pi_status12, pi_status13, pi_status14, pi_status15, pi_status16, pi_status17, pi_status18, pi_status19, pi_status20, pi_status21, pi_position1, pi_position2, pi_position3, pi_position4, pi_position5, pi_position6, pi_position7, pi_position8, pi_position9, pi_position10, pi_position11, pi_position12, pi_position13, pi_position14, pi_position15, pi_position16, pi_position17, pi_position18, pi_position19, pi_position20, pi_position21, pi_periodic, pi_date from periodicincidental_tb")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    if p.JoinQuery != "" {
        ret.WriteString(", ")
        ret.WriteString(p.JoinQuery)
    }

    ret.WriteString(" where 1=1 ")
    

    return ret.String()
}

func (p *PeriodicincidentalManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from periodicincidental_tb")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    if p.JoinQuery != "" {
        ret.WriteString(", ")
        ret.WriteString(p.JoinQuery)
    }

    ret.WriteString(" where 1=1 ")
    

    return ret.String()
}

func (p *PeriodicincidentalManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select pi_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from periodicincidental_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    


    return ret.String()
}

func (p *PeriodicincidentalManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate periodicincidental_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *PeriodicincidentalManager) Insert(item *Periodicincidental) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now().UTC().Add(time.Hour * 9)
        //t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        if config.Database.Type == config.Postgresql {
          query = "insert into periodicincidental_tb (pi_id, pi_result1, pi_result2, pi_result3, pi_result4, pi_result5, pi_result6, pi_result7, pi_result8, pi_result9, pi_result10, pi_result11, pi_result12, pi_result13, pi_result14, pi_result15, pi_result16, pi_result17, pi_result18, pi_result19, pi_result20, pi_result21, pi_status1, pi_status2, pi_status3, pi_status4, pi_status5, pi_status6, pi_status7, pi_status8, pi_status9, pi_status10, pi_status11, pi_status12, pi_status13, pi_status14, pi_status15, pi_status16, pi_status17, pi_status18, pi_status19, pi_status20, pi_status21, pi_position1, pi_position2, pi_position3, pi_position4, pi_position5, pi_position6, pi_position7, pi_position8, pi_position9, pi_position10, pi_position11, pi_position12, pi_position13, pi_position14, pi_position15, pi_position16, pi_position17, pi_position18, pi_position19, pi_position20, pi_position21, pi_periodic, pi_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66)"
        } else {
          query = "insert into periodicincidental_tb (pi_id, pi_result1, pi_result2, pi_result3, pi_result4, pi_result5, pi_result6, pi_result7, pi_result8, pi_result9, pi_result10, pi_result11, pi_result12, pi_result13, pi_result14, pi_result15, pi_result16, pi_result17, pi_result18, pi_result19, pi_result20, pi_result21, pi_status1, pi_status2, pi_status3, pi_status4, pi_status5, pi_status6, pi_status7, pi_status8, pi_status9, pi_status10, pi_status11, pi_status12, pi_status13, pi_status14, pi_status15, pi_status16, pi_status17, pi_status18, pi_status19, pi_status20, pi_status21, pi_position1, pi_position2, pi_position3, pi_position4, pi_position5, pi_position6, pi_position7, pi_position8, pi_position9, pi_position10, pi_position11, pi_position12, pi_position13, pi_position14, pi_position15, pi_position16, pi_position17, pi_position18, pi_position19, pi_position20, pi_position21, pi_periodic, pi_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Result1, item.Result2, item.Result3, item.Result4, item.Result5, item.Result6, item.Result7, item.Result8, item.Result9, item.Result10, item.Result11, item.Result12, item.Result13, item.Result14, item.Result15, item.Result16, item.Result17, item.Result18, item.Result19, item.Result20, item.Result21, item.Status1, item.Status2, item.Status3, item.Status4, item.Status5, item.Status6, item.Status7, item.Status8, item.Status9, item.Status10, item.Status11, item.Status12, item.Status13, item.Status14, item.Status15, item.Status16, item.Status17, item.Status18, item.Status19, item.Status20, item.Status21, item.Position1, item.Position2, item.Position3, item.Position4, item.Position5, item.Position6, item.Position7, item.Position8, item.Position9, item.Position10, item.Position11, item.Position12, item.Position13, item.Position14, item.Position15, item.Position16, item.Position17, item.Position18, item.Position19, item.Position20, item.Position21, item.Periodic, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into periodicincidental_tb (pi_result1, pi_result2, pi_result3, pi_result4, pi_result5, pi_result6, pi_result7, pi_result8, pi_result9, pi_result10, pi_result11, pi_result12, pi_result13, pi_result14, pi_result15, pi_result16, pi_result17, pi_result18, pi_result19, pi_result20, pi_result21, pi_status1, pi_status2, pi_status3, pi_status4, pi_status5, pi_status6, pi_status7, pi_status8, pi_status9, pi_status10, pi_status11, pi_status12, pi_status13, pi_status14, pi_status15, pi_status16, pi_status17, pi_status18, pi_status19, pi_status20, pi_status21, pi_position1, pi_position2, pi_position3, pi_position4, pi_position5, pi_position6, pi_position7, pi_position8, pi_position9, pi_position10, pi_position11, pi_position12, pi_position13, pi_position14, pi_position15, pi_position16, pi_position17, pi_position18, pi_position19, pi_position20, pi_position21, pi_periodic, pi_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65)"
        } else {
          query = "insert into periodicincidental_tb (pi_result1, pi_result2, pi_result3, pi_result4, pi_result5, pi_result6, pi_result7, pi_result8, pi_result9, pi_result10, pi_result11, pi_result12, pi_result13, pi_result14, pi_result15, pi_result16, pi_result17, pi_result18, pi_result19, pi_result20, pi_result21, pi_status1, pi_status2, pi_status3, pi_status4, pi_status5, pi_status6, pi_status7, pi_status8, pi_status9, pi_status10, pi_status11, pi_status12, pi_status13, pi_status14, pi_status15, pi_status16, pi_status17, pi_status18, pi_status19, pi_status20, pi_status21, pi_position1, pi_position2, pi_position3, pi_position4, pi_position5, pi_position6, pi_position7, pi_position8, pi_position9, pi_position10, pi_position11, pi_position12, pi_position13, pi_position14, pi_position15, pi_position16, pi_position17, pi_position18, pi_position19, pi_position20, pi_position21, pi_periodic, pi_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Result1, item.Result2, item.Result3, item.Result4, item.Result5, item.Result6, item.Result7, item.Result8, item.Result9, item.Result10, item.Result11, item.Result12, item.Result13, item.Result14, item.Result15, item.Result16, item.Result17, item.Result18, item.Result19, item.Result20, item.Result21, item.Status1, item.Status2, item.Status3, item.Status4, item.Status5, item.Status6, item.Status7, item.Status8, item.Status9, item.Status10, item.Status11, item.Status12, item.Status13, item.Status14, item.Status15, item.Status16, item.Status17, item.Status18, item.Status19, item.Status20, item.Status21, item.Position1, item.Position2, item.Position3, item.Position4, item.Position5, item.Position6, item.Position7, item.Position8, item.Position9, item.Position10, item.Position11, item.Position12, item.Position13, item.Position14, item.Position15, item.Position16, item.Position17, item.Position18, item.Position19, item.Position20, item.Position21, item.Periodic, item.Date)
    }
    
    if err == nil {
        p.Result = &res
        
    } else {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
        p.Result = nil
    }

    return err
}

func (p *PeriodicincidentalManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from periodicincidental_tb where pi_id = ")
    if config.Database.Type == config.Postgresql {
        query.WriteString("$1")
    } else {
        query.WriteString("?")
    }
    _, err := p.Exec(query.String(), id)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *PeriodicincidentalManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from periodicincidental_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *PeriodicincidentalManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
    var params []any
    if initParams != nil {
        params = append(params, initParams...)
    }

    pos := 1

    var query strings.Builder
	query.WriteString(initQuery)

    for _, arg := range args {
        switch v := arg.(type) {        
        case Where:
            item := v

            if strings.Contains(item.Column, "_") {
                query.WriteString(" and ")
            } else {
                query.WriteString(" and pi_")
            }
            query.WriteString(item.Column)

            if item.Compare == "in" {
                query.WriteString(" in (")
                query.WriteString(strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]"))
                query.WriteString(")")
            } else if item.Compare == "not in" {
                query.WriteString(" not in (")
                query.WriteString(strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]"))
                query.WriteString(")")
            } else if item.Compare == "between" {
                if config.Database.Type == config.Postgresql {
                    query.WriteString(fmt.Sprintf(" between $%v and $%v", pos, pos + 1))
                    pos += 2
                } else {
                    query.WriteString(" between ? and ?")
                }

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                if config.Database.Type == config.Postgresql {
                    query.WriteString(" ")
                    query.WriteString(item.Compare)
                    query.WriteString(" $")
                    query.WriteString(fmt.Sprintf("%v", pos))
                    pos++
                } else {
                    query.WriteString(" ")
                    query.WriteString(item.Compare)
                    query.WriteString(" ?")
                }
                if item.Compare == "like" {
                    params = append(params, "%" + item.Value.(string) + "%")
                } else {
                    params = append(params, item.Value)                
                }
            }
        case Custom:
             item := v

            query.WriteString(" and ")
            query.WriteString(item.Query)
        }        
    }

	query.WriteString(postQuery)

    return query.String(), params
}

func (p *PeriodicincidentalManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from periodicincidental_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *PeriodicincidentalManager) Update(item *Periodicincidental) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update periodicincidental_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" pi_result1 = $1, pi_result2 = $2, pi_result3 = $3, pi_result4 = $4, pi_result5 = $5, pi_result6 = $6, pi_result7 = $7, pi_result8 = $8, pi_result9 = $9, pi_result10 = $10, pi_result11 = $11, pi_result12 = $12, pi_result13 = $13, pi_result14 = $14, pi_result15 = $15, pi_result16 = $16, pi_result17 = $17, pi_result18 = $18, pi_result19 = $19, pi_result20 = $20, pi_result21 = $21, pi_status1 = $22, pi_status2 = $23, pi_status3 = $24, pi_status4 = $25, pi_status5 = $26, pi_status6 = $27, pi_status7 = $28, pi_status8 = $29, pi_status9 = $30, pi_status10 = $31, pi_status11 = $32, pi_status12 = $33, pi_status13 = $34, pi_status14 = $35, pi_status15 = $36, pi_status16 = $37, pi_status17 = $38, pi_status18 = $39, pi_status19 = $40, pi_status20 = $41, pi_status21 = $42, pi_position1 = $43, pi_position2 = $44, pi_position3 = $45, pi_position4 = $46, pi_position5 = $47, pi_position6 = $48, pi_position7 = $49, pi_position8 = $50, pi_position9 = $51, pi_position10 = $52, pi_position11 = $53, pi_position12 = $54, pi_position13 = $55, pi_position14 = $56, pi_position15 = $57, pi_position16 = $58, pi_position17 = $59, pi_position18 = $60, pi_position19 = $61, pi_position20 = $62, pi_position21 = $63, pi_periodic = $64, pi_date = $65 where pi_id = $66")
    } else {
        query.WriteString(" pi_result1 = ?, pi_result2 = ?, pi_result3 = ?, pi_result4 = ?, pi_result5 = ?, pi_result6 = ?, pi_result7 = ?, pi_result8 = ?, pi_result9 = ?, pi_result10 = ?, pi_result11 = ?, pi_result12 = ?, pi_result13 = ?, pi_result14 = ?, pi_result15 = ?, pi_result16 = ?, pi_result17 = ?, pi_result18 = ?, pi_result19 = ?, pi_result20 = ?, pi_result21 = ?, pi_status1 = ?, pi_status2 = ?, pi_status3 = ?, pi_status4 = ?, pi_status5 = ?, pi_status6 = ?, pi_status7 = ?, pi_status8 = ?, pi_status9 = ?, pi_status10 = ?, pi_status11 = ?, pi_status12 = ?, pi_status13 = ?, pi_status14 = ?, pi_status15 = ?, pi_status16 = ?, pi_status17 = ?, pi_status18 = ?, pi_status19 = ?, pi_status20 = ?, pi_status21 = ?, pi_position1 = ?, pi_position2 = ?, pi_position3 = ?, pi_position4 = ?, pi_position5 = ?, pi_position6 = ?, pi_position7 = ?, pi_position8 = ?, pi_position9 = ?, pi_position10 = ?, pi_position11 = ?, pi_position12 = ?, pi_position13 = ?, pi_position14 = ?, pi_position15 = ?, pi_position16 = ?, pi_position17 = ?, pi_position18 = ?, pi_position19 = ?, pi_position20 = ?, pi_position21 = ?, pi_periodic = ?, pi_date = ? where pi_id = ?")
    }

	_, err := p.Exec(query.String() , item.Result1, item.Result2, item.Result3, item.Result4, item.Result5, item.Result6, item.Result7, item.Result8, item.Result9, item.Result10, item.Result11, item.Result12, item.Result13, item.Result14, item.Result15, item.Result16, item.Result17, item.Result18, item.Result19, item.Result20, item.Result21, item.Status1, item.Status2, item.Status3, item.Status4, item.Status5, item.Status6, item.Status7, item.Status8, item.Status9, item.Status10, item.Status11, item.Status12, item.Status13, item.Status14, item.Status15, item.Status16, item.Status17, item.Status18, item.Status19, item.Status20, item.Status21, item.Position1, item.Position2, item.Position3, item.Position4, item.Position5, item.Position6, item.Position7, item.Position8, item.Position9, item.Position10, item.Position11, item.Position12, item.Position13, item.Position14, item.Position15, item.Position16, item.Position17, item.Position18, item.Position19, item.Position20, item.Position21, item.Periodic, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *PeriodicincidentalManager) UpdateWhere(columns []periodicincidental.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update periodicincidental_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == periodicincidental.ColumnId {
        initQuery.WriteString("pi_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult1 {
        initQuery.WriteString("pi_result1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult2 {
        initQuery.WriteString("pi_result2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult3 {
        initQuery.WriteString("pi_result3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult4 {
        initQuery.WriteString("pi_result4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult5 {
        initQuery.WriteString("pi_result5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult6 {
        initQuery.WriteString("pi_result6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult7 {
        initQuery.WriteString("pi_result7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult8 {
        initQuery.WriteString("pi_result8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult9 {
        initQuery.WriteString("pi_result9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult10 {
        initQuery.WriteString("pi_result10 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult11 {
        initQuery.WriteString("pi_result11 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult12 {
        initQuery.WriteString("pi_result12 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult13 {
        initQuery.WriteString("pi_result13 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult14 {
        initQuery.WriteString("pi_result14 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult15 {
        initQuery.WriteString("pi_result15 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult16 {
        initQuery.WriteString("pi_result16 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult17 {
        initQuery.WriteString("pi_result17 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult18 {
        initQuery.WriteString("pi_result18 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult19 {
        initQuery.WriteString("pi_result19 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult20 {
        initQuery.WriteString("pi_result20 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnResult21 {
        initQuery.WriteString("pi_result21 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus1 {
        initQuery.WriteString("pi_status1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus2 {
        initQuery.WriteString("pi_status2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus3 {
        initQuery.WriteString("pi_status3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus4 {
        initQuery.WriteString("pi_status4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus5 {
        initQuery.WriteString("pi_status5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus6 {
        initQuery.WriteString("pi_status6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus7 {
        initQuery.WriteString("pi_status7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus8 {
        initQuery.WriteString("pi_status8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus9 {
        initQuery.WriteString("pi_status9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus10 {
        initQuery.WriteString("pi_status10 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus11 {
        initQuery.WriteString("pi_status11 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus12 {
        initQuery.WriteString("pi_status12 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus13 {
        initQuery.WriteString("pi_status13 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus14 {
        initQuery.WriteString("pi_status14 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus15 {
        initQuery.WriteString("pi_status15 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus16 {
        initQuery.WriteString("pi_status16 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus17 {
        initQuery.WriteString("pi_status17 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus18 {
        initQuery.WriteString("pi_status18 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus19 {
        initQuery.WriteString("pi_status19 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus20 {
        initQuery.WriteString("pi_status20 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnStatus21 {
        initQuery.WriteString("pi_status21 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition1 {
        initQuery.WriteString("pi_position1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition2 {
        initQuery.WriteString("pi_position2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition3 {
        initQuery.WriteString("pi_position3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition4 {
        initQuery.WriteString("pi_position4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition5 {
        initQuery.WriteString("pi_position5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition6 {
        initQuery.WriteString("pi_position6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition7 {
        initQuery.WriteString("pi_position7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition8 {
        initQuery.WriteString("pi_position8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition9 {
        initQuery.WriteString("pi_position9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition10 {
        initQuery.WriteString("pi_position10 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition11 {
        initQuery.WriteString("pi_position11 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition12 {
        initQuery.WriteString("pi_position12 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition13 {
        initQuery.WriteString("pi_position13 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition14 {
        initQuery.WriteString("pi_position14 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition15 {
        initQuery.WriteString("pi_position15 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition16 {
        initQuery.WriteString("pi_position16 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition17 {
        initQuery.WriteString("pi_position17 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition18 {
        initQuery.WriteString("pi_position18 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition19 {
        initQuery.WriteString("pi_position19 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition20 {
        initQuery.WriteString("pi_position20 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPosition21 {
        initQuery.WriteString("pi_position21 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnPeriodic {
        initQuery.WriteString("pi_periodic = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == periodicincidental.ColumnDate {
        initQuery.WriteString("pi_date = ?")
        initParams = append(initParams, v.Value)
        
        }
    }

    initQuery.WriteString(" where 1=1 ")

    query, params := p.MakeQuery(initQuery.String(), "", initParams, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}


/*

func (p *PeriodicincidentalManager) UpdateResult1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result1 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result2 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result3 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result4 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result5 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result6 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult7(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result7 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult8(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result8 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult9(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result9 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult10(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result10 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult11(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result11 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult12(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result12 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult13(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result13 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult14(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result14 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult15(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result15 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult16(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result16 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult17(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result17 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult18(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result18 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult19(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result19 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult20(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result20 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateResult21(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_result21 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status1 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status2 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status3 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status4 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status5 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status6 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus7(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status7 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus8(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status8 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus9(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status9 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus10(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status10 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus11(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status11 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus12(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status12 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus13(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status13 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus14(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status14 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus15(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status15 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus16(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status16 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus17(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status17 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus18(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status18 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus19(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status19 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus20(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status20 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdateStatus21(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_status21 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position1 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position2 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position3 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position4 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position5 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position6 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition7(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position7 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition8(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position8 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition9(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position9 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition10(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position10 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition11(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position11 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition12(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position12 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition13(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position13 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition14(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position14 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition15(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position15 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition16(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position16 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition17(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position17 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition18(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position18 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition19(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position19 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition20(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position20 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePosition21(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_position21 = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PeriodicincidentalManager) UpdatePeriodic(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_periodic = ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

/*

func (p *PeriodicincidentalManager) IncreasePeriodic(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update periodicincidental_tb set pi_periodic = pi_periodic + ? where pi_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *PeriodicincidentalManager) GetIdentity() int64 {
    if !p.Conn.IsConnect() {
        return 0
    }

    id, err := (*p.Result).LastInsertId()

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
        return 0
    } else {
        return id
    }
}

func (p *Periodicincidental) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *PeriodicincidentalManager) ReadRow(rows *sql.Rows) *Periodicincidental {
    var item Periodicincidental
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Result1, &item.Result2, &item.Result3, &item.Result4, &item.Result5, &item.Result6, &item.Result7, &item.Result8, &item.Result9, &item.Result10, &item.Result11, &item.Result12, &item.Result13, &item.Result14, &item.Result15, &item.Result16, &item.Result17, &item.Result18, &item.Result19, &item.Result20, &item.Result21, &item.Status1, &item.Status2, &item.Status3, &item.Status4, &item.Status5, &item.Status6, &item.Status7, &item.Status8, &item.Status9, &item.Status10, &item.Status11, &item.Status12, &item.Status13, &item.Status14, &item.Status15, &item.Status16, &item.Status17, &item.Status18, &item.Status19, &item.Status20, &item.Status21, &item.Position1, &item.Position2, &item.Position3, &item.Position4, &item.Position5, &item.Position6, &item.Position7, &item.Position8, &item.Position9, &item.Position10, &item.Position11, &item.Position12, &item.Position13, &item.Position14, &item.Position15, &item.Position16, &item.Position17, &item.Position18, &item.Position19, &item.Position20, &item.Position21, &item.Periodic, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" || item.Date == "9999-01-01 00:00:00" {
            item.Date = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Date = strings.ReplaceAll(strings.ReplaceAll(item.Date, "T", " "), "Z", "")
        }
		
        

    } else {
        return nil
    }

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
        return nil
    } else {

        item.InitExtra()
        
        return &item
    }
}

func (p *PeriodicincidentalManager) ReadRows(rows *sql.Rows) []Periodicincidental {
    items := make([]Periodicincidental, 0)

    for rows.Next() {
        var item Periodicincidental
        
    
        err := rows.Scan(&item.Id, &item.Result1, &item.Result2, &item.Result3, &item.Result4, &item.Result5, &item.Result6, &item.Result7, &item.Result8, &item.Result9, &item.Result10, &item.Result11, &item.Result12, &item.Result13, &item.Result14, &item.Result15, &item.Result16, &item.Result17, &item.Result18, &item.Result19, &item.Result20, &item.Result21, &item.Status1, &item.Status2, &item.Status3, &item.Status4, &item.Status5, &item.Status6, &item.Status7, &item.Status8, &item.Status9, &item.Status10, &item.Status11, &item.Status12, &item.Status13, &item.Status14, &item.Status15, &item.Status16, &item.Status17, &item.Status18, &item.Status19, &item.Status20, &item.Status21, &item.Position1, &item.Position2, &item.Position3, &item.Position4, &item.Position5, &item.Position6, &item.Position7, &item.Position8, &item.Position9, &item.Position10, &item.Position11, &item.Position12, &item.Position13, &item.Position14, &item.Position15, &item.Position16, &item.Position17, &item.Position18, &item.Position19, &item.Position20, &item.Position21, &item.Periodic, &item.Date)
        if err != nil {
           if p.Log {
             log.Error().Str("error", err.Error()).Msg("SQL")
           }
           break
        }

        
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" || item.Date == "9999-01-01 00:00:00" {
            item.Date = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Date = strings.ReplaceAll(strings.ReplaceAll(item.Date, "T", " "), "Z", "")
        }
		
		
        
        item.InitExtra()        
        
        items = append(items, item)
    }


     return items
}

func (p *PeriodicincidentalManager) Get(id int64) *Periodicincidental {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and pi_id = ?")

    
    
    rows, err := p.Query(query.String(), id)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
       return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}

func (p *PeriodicincidentalManager) GetWhere(args []any) *Periodicincidental {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *PeriodicincidentalManager) Count(args []any) int {
    if !p.Conn.IsConnect() {
        return 0
    }

    query, params := p.MakeQuery(p.GetQuerySelect(), p.GroupQuery, nil, args)
    rows, err := p.Query(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
       return 0
    }

    defer rows.Close()

    if !rows.Next() {
        return 0
    }

    cnt := 0
    err = rows.Scan(&cnt)

    if err != nil {
        return 0
    } else {
        return cnt
    }
}

func (p *PeriodicincidentalManager) FindAll() []Periodicincidental {
    return p.Find(nil)
}

func (p *PeriodicincidentalManager) Find(args []any) []Periodicincidental {
    if !p.Conn.IsConnect() {
        items := make([]Periodicincidental, 0)
        return items
    }

    var params []any
    baseQuery := p.GetQuery()

    var query strings.Builder

    page := 0
    pagesize := 0
    orderby := ""

    pos := 1
    
    for _, arg := range args {
        switch v := arg.(type) {
        case PagingType:
            item := v
            page = item.Page
            pagesize = item.Pagesize            
        case OrderingType:
            item := v
            orderby = item.Order
        case LimitType:
            item := v
            page = 1
            pagesize = item.Limit
        case OptionType:
            item := v
            if item.Limit > 0 {
                page = 1
                pagesize = item.Limit
            } else {
                page = item.Page
                pagesize = item.Pagesize                
            }
            orderby = item.Order
        case Where:
            item := v

            if strings.Contains(item.Column, "_") {
                query.WriteString(" and ")
            } else {
                query.WriteString(" and pi_")
            }
            query.WriteString(item.Column)
            
            if item.Compare == "in" {
                query.WriteString(" in (")
                query.WriteString(strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]"))
                query.WriteString(")")
            } else if item.Compare == "not in" {
                query.WriteString(" not in (")
                query.WriteString(strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]"))
                query.WriteString(")")
            } else if item.Compare == "between" {
                if config.Database.Type == config.Postgresql {
                    query.WriteString(fmt.Sprintf(" between $%v and $%v", pos, pos + 1))
                    pos += 2
                } else {
                    query.WriteString(" between ? and ?")
                }

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                if config.Database.Type == config.Postgresql {
                    query.WriteString(" ")
                    query.WriteString(item.Compare)
                    query.WriteString(" $")
                    query.WriteString(fmt.Sprintf("%v", pos))
                    pos++
                } else {
                    query.WriteString(" ")
                    query.WriteString(item.Compare)
                    query.WriteString(" ?")
                }
                if item.Compare == "like" {
                    params = append(params, "%" + item.Value.(string) + "%")
                } else {
                    params = append(params, item.Value)                
                }
            }
        case Custom:
             item := v

            query.WriteString(" and ")
            query.WriteString(item.Query)
        case Base:
             item := v

             baseQuery = item.Query
        }
    }

    query.WriteString(p.GroupQuery)
    
    startpage := (page - 1) * pagesize
    
    if page > 0 && pagesize > 0 {
        if orderby == "" {
            orderby = "pi_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "pi_" + orderby
                }
            }
            
        }
        query.WriteString(" order by ")
        query.WriteString(orderby)
        if config.Database.Type == config.Postgresql {
            query.WriteString(fmt.Sprintf(" limit $%v offset $%v", pos, pos + 1))
            params = append(params, pagesize)
            params = append(params, startpage)
        } else if config.Database.Type == config.Mysql {
            query.WriteString(" limit ? offset ?")
            params = append(params, pagesize)
            params = append(params, startpage)
        } else if config.Database.Type == config.Sqlserver {
            query.WriteString("OFFSET ? ROWS FETCH NEXT ? ROWS ONLY")
            params = append(params, startpage)
            params = append(params, pagesize)
        }
    } else {
        if orderby == "" {
            orderby = "pi_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "pi_" + orderby
                }
            }
        }
        query.WriteString(" order by ")
        query.WriteString(orderby)
    }

    rows, err := p.Query(baseQuery + query.String(), params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
        items := make([]Periodicincidental, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *PeriodicincidentalManager) GetByPeriodic(periodic int64, args ...any) *Periodicincidental {
    rets := make([]any, 0)
    rets = append(rets, args...)
    if periodic != 0 {
        rets = append(rets, Where{Column:"periodic", Value:periodic, Compare:"="})        
    }
    
    items := p.Find(rets)

    if len(items) > 0 {
        return &items[0]
    } else {
        return nil
    }
}




func (p *PeriodicincidentalManager) GroupBy(name string, args []any) []Groupby {
    if !p.Conn.IsConnect() {
        var items []Groupby
        return items
    }

    var params []any
    baseQuery := p.GetQueryGroup(name)
    var query strings.Builder
    pos := 1

    for _, arg := range args {
        switch v := arg.(type) {
        case Where:
            item := v

            if strings.Contains(item.Column, "_") {
                query.WriteString(" and ")
            } else {
                query.WriteString(" and pi_")
            }
            query.WriteString(item.Column)
            
            if item.Compare == "in" {
                query.WriteString(" in (")
                query.WriteString(strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]"))
                query.WriteString(")")
            } else if item.Compare == "not in" {
                query.WriteString(" not in (")
                query.WriteString(strings.Trim(strings.Replace(fmt.Sprint(item.Value), " ", ", ", -1), "[]"))
                query.WriteString(")")
            } else if item.Compare == "between" {
                if config.Database.Type == config.Postgresql {
                    query.WriteString(fmt.Sprintf(" between $%v and $%v", pos, pos + 1))
                    pos += 2
                } else {
                    query.WriteString(" between ? and ?")
                }

                s := item.Value.([2]string)
                params = append(params, s[0])
                params = append(params, s[1])
            } else {
                if config.Database.Type == config.Postgresql {
                    query.WriteString(" ")
                    query.WriteString(item.Compare)
                    query.WriteString(" $")
                    query.WriteString(fmt.Sprintf("%v", pos))
                    pos++
                } else {
                    query.WriteString(" ")
                    query.WriteString(item.Compare)
                    query.WriteString(" ?")
                }
                if item.Compare == "like" {
                    params = append(params, "%" + item.Value.(string) + "%")
                } else {
                    params = append(params, item.Value)                
                }
            }
        case Custom:
             item := v

            query.WriteString(" and ")
            query.WriteString(item.Query)
        case Base:
             item := v

             baseQuery = item.Query
        }
    }
    
    query.WriteString(" group by pi_")
    query.WriteString(name)

    rows, err := p.Query(baseQuery + query.String(), params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
        var items []Groupby
        return items
    }

    defer rows.Close()

    var items []Groupby

    for rows.Next() {
        var item Groupby
        err := rows.Scan(&item.Value, &item.Count)
        if err != nil {
           if p.Log {
                log.Error().Str("error", err.Error()).Msg("SQL")
           }
           break
        }

        items = append(items, item)
    }

    return items
}



func (p *PeriodicincidentalManager) MakeMap(items []Periodicincidental) map[int64]Periodicincidental {
     ret := make(map[int64]Periodicincidental)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *PeriodicincidentalManager) FindToMap(args []any) map[int64]Periodicincidental {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *PeriodicincidentalManager) FindAllToMap() map[int64]Periodicincidental {
     items := p.Find(nil)
     return p.MakeMap(items)
}


