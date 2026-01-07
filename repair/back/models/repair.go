package models

import (
    "repair/models/repair"
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

type Repair struct {
            
    Id                int64 `json:"id"`         
    Type                int `json:"type"`         
    Status                int `json:"status"`         
    Calculatetype                int `json:"calculatetype"`         
    Provision                int `json:"provision"`         
    Complex1                string `json:"complex1"`         
    Complex2                string `json:"complex2"`         
    Completionyear                int `json:"completionyear"`         
    Completionmonth                int `json:"completionmonth"`         
    Completionday                int `json:"completionday"`         
    Parcelrate                Double `json:"parcelrate"`         
    Planyears                int `json:"planyears"`         
    Info1                string `json:"info1"`         
    Info2                string `json:"info2"`         
    Info3                string `json:"info3"`         
    Info4                string `json:"info4"`         
    Info5                string `json:"info5"`         
    Info6                string `json:"info6"`         
    Info7                string `json:"info7"`         
    Info8                string `json:"info8"`         
    Info9                string `json:"info9"`         
    Info10                string `json:"info10"`         
    Info11                string `json:"info11"`         
    Structure1                string `json:"structure1"`         
    Structure2                string `json:"structure2"`         
    Structure3                string `json:"structure3"`         
    Structure4                string `json:"structure4"`         
    Structure5                string `json:"structure5"`         
    Structure6                string `json:"structure6"`         
    Structure7                string `json:"structure7"`         
    Structure8                string `json:"structure8"`         
    Structure9                string `json:"structure9"`         
    Structure10                string `json:"structure10"`         
    Structure11                string `json:"structure11"`         
    Structure12                string `json:"structure12"`         
    Structure13                string `json:"structure13"`         
    Structure14                string `json:"structure14"`         
    Reviewcontent1                string `json:"reviewcontent1"`         
    Reviewcontent2                string `json:"reviewcontent2"`         
    Reviewcontent3                string `json:"reviewcontent3"`         
    Reviewcontent4                string `json:"reviewcontent4"`         
    Reviewcontent5                string `json:"reviewcontent5"`         
    Reviewcontent6                string `json:"reviewcontent6"`         
    Reviewcontent7                string `json:"reviewcontent7"`         
    Savingprice                Double `json:"savingprice"`         
    Price1                string `json:"price1"`         
    Price2                string `json:"price2"`         
    Price3                string `json:"price3"`         
    Price4                string `json:"price4"`         
    Price5                string `json:"price5"`         
    Reportdate                string `json:"reportdate"`         
    Content1                string `json:"content1"`         
    Content2                string `json:"content2"`         
    Content3                string `json:"content3"`         
    Content4                string `json:"content4"`         
    Content5                string `json:"content5"`         
    Content6                string `json:"content6"`         
    Content7                string `json:"content7"`         
    Content8                string `json:"content8"`         
    Content9                string `json:"content9"`         
    Content10                string `json:"content10"`         
    Content11                string `json:"content11"`         
    Periodtype                repair.Periodtype `json:"periodtype"`         
    Remark                string `json:"remark"`         
    Apt                int64 `json:"apt"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type RepairManager struct {
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



func (c *Repair) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewRepairManager(conn *Connection) *RepairManager {
    var item RepairManager


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

func (p *RepairManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *RepairManager) SetIndex(index string) {
    p.Index = index
}

func (p *RepairManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *RepairManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *RepairManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *RepairManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *RepairManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select r_id, r_type, r_status, r_calculatetype, r_provision, r_complex1, r_complex2, r_completionyear, r_completionmonth, r_completionday, r_parcelrate, r_planyears, r_info1, r_info2, r_info3, r_info4, r_info5, r_info6, r_info7, r_info8, r_info9, r_info10, r_info11, r_structure1, r_structure2, r_structure3, r_structure4, r_structure5, r_structure6, r_structure7, r_structure8, r_structure9, r_structure10, r_structure11, r_structure12, r_structure13, r_structure14, r_reviewcontent1, r_reviewcontent2, r_reviewcontent3, r_reviewcontent4, r_reviewcontent5, r_reviewcontent6, r_reviewcontent7, r_savingprice, r_price1, r_price2, r_price3, r_price4, r_price5, r_reportdate, r_content1, r_content2, r_content3, r_content4, r_content5, r_content6, r_content7, r_content8, r_content9, r_content10, r_content11, r_periodtype, r_remark, r_apt, r_date from repair_tb")

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

func (p *RepairManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from repair_tb")

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

func (p *RepairManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select r_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from repair_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    


    return ret.String()
}

func (p *RepairManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate repair_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *RepairManager) Insert(item *Repair) error {
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
          query = "insert into repair_tb (r_id, r_type, r_status, r_calculatetype, r_provision, r_complex1, r_complex2, r_completionyear, r_completionmonth, r_completionday, r_parcelrate, r_planyears, r_info1, r_info2, r_info3, r_info4, r_info5, r_info6, r_info7, r_info8, r_info9, r_info10, r_info11, r_structure1, r_structure2, r_structure3, r_structure4, r_structure5, r_structure6, r_structure7, r_structure8, r_structure9, r_structure10, r_structure11, r_structure12, r_structure13, r_structure14, r_reviewcontent1, r_reviewcontent2, r_reviewcontent3, r_reviewcontent4, r_reviewcontent5, r_reviewcontent6, r_reviewcontent7, r_savingprice, r_price1, r_price2, r_price3, r_price4, r_price5, r_reportdate, r_content1, r_content2, r_content3, r_content4, r_content5, r_content6, r_content7, r_content8, r_content9, r_content10, r_content11, r_periodtype, r_remark, r_apt, r_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66)"
        } else {
          query = "insert into repair_tb (r_id, r_type, r_status, r_calculatetype, r_provision, r_complex1, r_complex2, r_completionyear, r_completionmonth, r_completionday, r_parcelrate, r_planyears, r_info1, r_info2, r_info3, r_info4, r_info5, r_info6, r_info7, r_info8, r_info9, r_info10, r_info11, r_structure1, r_structure2, r_structure3, r_structure4, r_structure5, r_structure6, r_structure7, r_structure8, r_structure9, r_structure10, r_structure11, r_structure12, r_structure13, r_structure14, r_reviewcontent1, r_reviewcontent2, r_reviewcontent3, r_reviewcontent4, r_reviewcontent5, r_reviewcontent6, r_reviewcontent7, r_savingprice, r_price1, r_price2, r_price3, r_price4, r_price5, r_reportdate, r_content1, r_content2, r_content3, r_content4, r_content5, r_content6, r_content7, r_content8, r_content9, r_content10, r_content11, r_periodtype, r_remark, r_apt, r_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Type, item.Status, item.Calculatetype, item.Provision, item.Complex1, item.Complex2, item.Completionyear, item.Completionmonth, item.Completionday, item.Parcelrate, item.Planyears, item.Info1, item.Info2, item.Info3, item.Info4, item.Info5, item.Info6, item.Info7, item.Info8, item.Info9, item.Info10, item.Info11, item.Structure1, item.Structure2, item.Structure3, item.Structure4, item.Structure5, item.Structure6, item.Structure7, item.Structure8, item.Structure9, item.Structure10, item.Structure11, item.Structure12, item.Structure13, item.Structure14, item.Reviewcontent1, item.Reviewcontent2, item.Reviewcontent3, item.Reviewcontent4, item.Reviewcontent5, item.Reviewcontent6, item.Reviewcontent7, item.Savingprice, item.Price1, item.Price2, item.Price3, item.Price4, item.Price5, item.Reportdate, item.Content1, item.Content2, item.Content3, item.Content4, item.Content5, item.Content6, item.Content7, item.Content8, item.Content9, item.Content10, item.Content11, item.Periodtype, item.Remark, item.Apt, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into repair_tb (r_type, r_status, r_calculatetype, r_provision, r_complex1, r_complex2, r_completionyear, r_completionmonth, r_completionday, r_parcelrate, r_planyears, r_info1, r_info2, r_info3, r_info4, r_info5, r_info6, r_info7, r_info8, r_info9, r_info10, r_info11, r_structure1, r_structure2, r_structure3, r_structure4, r_structure5, r_structure6, r_structure7, r_structure8, r_structure9, r_structure10, r_structure11, r_structure12, r_structure13, r_structure14, r_reviewcontent1, r_reviewcontent2, r_reviewcontent3, r_reviewcontent4, r_reviewcontent5, r_reviewcontent6, r_reviewcontent7, r_savingprice, r_price1, r_price2, r_price3, r_price4, r_price5, r_reportdate, r_content1, r_content2, r_content3, r_content4, r_content5, r_content6, r_content7, r_content8, r_content9, r_content10, r_content11, r_periodtype, r_remark, r_apt, r_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65)"
        } else {
          query = "insert into repair_tb (r_type, r_status, r_calculatetype, r_provision, r_complex1, r_complex2, r_completionyear, r_completionmonth, r_completionday, r_parcelrate, r_planyears, r_info1, r_info2, r_info3, r_info4, r_info5, r_info6, r_info7, r_info8, r_info9, r_info10, r_info11, r_structure1, r_structure2, r_structure3, r_structure4, r_structure5, r_structure6, r_structure7, r_structure8, r_structure9, r_structure10, r_structure11, r_structure12, r_structure13, r_structure14, r_reviewcontent1, r_reviewcontent2, r_reviewcontent3, r_reviewcontent4, r_reviewcontent5, r_reviewcontent6, r_reviewcontent7, r_savingprice, r_price1, r_price2, r_price3, r_price4, r_price5, r_reportdate, r_content1, r_content2, r_content3, r_content4, r_content5, r_content6, r_content7, r_content8, r_content9, r_content10, r_content11, r_periodtype, r_remark, r_apt, r_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Type, item.Status, item.Calculatetype, item.Provision, item.Complex1, item.Complex2, item.Completionyear, item.Completionmonth, item.Completionday, item.Parcelrate, item.Planyears, item.Info1, item.Info2, item.Info3, item.Info4, item.Info5, item.Info6, item.Info7, item.Info8, item.Info9, item.Info10, item.Info11, item.Structure1, item.Structure2, item.Structure3, item.Structure4, item.Structure5, item.Structure6, item.Structure7, item.Structure8, item.Structure9, item.Structure10, item.Structure11, item.Structure12, item.Structure13, item.Structure14, item.Reviewcontent1, item.Reviewcontent2, item.Reviewcontent3, item.Reviewcontent4, item.Reviewcontent5, item.Reviewcontent6, item.Reviewcontent7, item.Savingprice, item.Price1, item.Price2, item.Price3, item.Price4, item.Price5, item.Reportdate, item.Content1, item.Content2, item.Content3, item.Content4, item.Content5, item.Content6, item.Content7, item.Content8, item.Content9, item.Content10, item.Content11, item.Periodtype, item.Remark, item.Apt, item.Date)
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

func (p *RepairManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from repair_tb where r_id = ")
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

func (p *RepairManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from repair_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *RepairManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and r_")
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

func (p *RepairManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from repair_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *RepairManager) Update(item *Repair) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update repair_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" r_type = $1, r_status = $2, r_calculatetype = $3, r_provision = $4, r_complex1 = $5, r_complex2 = $6, r_completionyear = $7, r_completionmonth = $8, r_completionday = $9, r_parcelrate = $10, r_planyears = $11, r_info1 = $12, r_info2 = $13, r_info3 = $14, r_info4 = $15, r_info5 = $16, r_info6 = $17, r_info7 = $18, r_info8 = $19, r_info9 = $20, r_info10 = $21, r_info11 = $22, r_structure1 = $23, r_structure2 = $24, r_structure3 = $25, r_structure4 = $26, r_structure5 = $27, r_structure6 = $28, r_structure7 = $29, r_structure8 = $30, r_structure9 = $31, r_structure10 = $32, r_structure11 = $33, r_structure12 = $34, r_structure13 = $35, r_structure14 = $36, r_reviewcontent1 = $37, r_reviewcontent2 = $38, r_reviewcontent3 = $39, r_reviewcontent4 = $40, r_reviewcontent5 = $41, r_reviewcontent6 = $42, r_reviewcontent7 = $43, r_savingprice = $44, r_price1 = $45, r_price2 = $46, r_price3 = $47, r_price4 = $48, r_price5 = $49, r_reportdate = $50, r_content1 = $51, r_content2 = $52, r_content3 = $53, r_content4 = $54, r_content5 = $55, r_content6 = $56, r_content7 = $57, r_content8 = $58, r_content9 = $59, r_content10 = $60, r_content11 = $61, r_periodtype = $62, r_remark = $63, r_apt = $64, r_date = $65 where r_id = $66")
    } else {
        query.WriteString(" r_type = ?, r_status = ?, r_calculatetype = ?, r_provision = ?, r_complex1 = ?, r_complex2 = ?, r_completionyear = ?, r_completionmonth = ?, r_completionday = ?, r_parcelrate = ?, r_planyears = ?, r_info1 = ?, r_info2 = ?, r_info3 = ?, r_info4 = ?, r_info5 = ?, r_info6 = ?, r_info7 = ?, r_info8 = ?, r_info9 = ?, r_info10 = ?, r_info11 = ?, r_structure1 = ?, r_structure2 = ?, r_structure3 = ?, r_structure4 = ?, r_structure5 = ?, r_structure6 = ?, r_structure7 = ?, r_structure8 = ?, r_structure9 = ?, r_structure10 = ?, r_structure11 = ?, r_structure12 = ?, r_structure13 = ?, r_structure14 = ?, r_reviewcontent1 = ?, r_reviewcontent2 = ?, r_reviewcontent3 = ?, r_reviewcontent4 = ?, r_reviewcontent5 = ?, r_reviewcontent6 = ?, r_reviewcontent7 = ?, r_savingprice = ?, r_price1 = ?, r_price2 = ?, r_price3 = ?, r_price4 = ?, r_price5 = ?, r_reportdate = ?, r_content1 = ?, r_content2 = ?, r_content3 = ?, r_content4 = ?, r_content5 = ?, r_content6 = ?, r_content7 = ?, r_content8 = ?, r_content9 = ?, r_content10 = ?, r_content11 = ?, r_periodtype = ?, r_remark = ?, r_apt = ?, r_date = ? where r_id = ?")
    }

	_, err := p.Exec(query.String() , item.Type, item.Status, item.Calculatetype, item.Provision, item.Complex1, item.Complex2, item.Completionyear, item.Completionmonth, item.Completionday, item.Parcelrate, item.Planyears, item.Info1, item.Info2, item.Info3, item.Info4, item.Info5, item.Info6, item.Info7, item.Info8, item.Info9, item.Info10, item.Info11, item.Structure1, item.Structure2, item.Structure3, item.Structure4, item.Structure5, item.Structure6, item.Structure7, item.Structure8, item.Structure9, item.Structure10, item.Structure11, item.Structure12, item.Structure13, item.Structure14, item.Reviewcontent1, item.Reviewcontent2, item.Reviewcontent3, item.Reviewcontent4, item.Reviewcontent5, item.Reviewcontent6, item.Reviewcontent7, item.Savingprice, item.Price1, item.Price2, item.Price3, item.Price4, item.Price5, item.Reportdate, item.Content1, item.Content2, item.Content3, item.Content4, item.Content5, item.Content6, item.Content7, item.Content8, item.Content9, item.Content10, item.Content11, item.Periodtype, item.Remark, item.Apt, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *RepairManager) UpdateWhere(columns []repair.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update repair_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == repair.ColumnId {
        initQuery.WriteString("r_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnType {
        initQuery.WriteString("r_type = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStatus {
        initQuery.WriteString("r_status = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnCalculatetype {
        initQuery.WriteString("r_calculatetype = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnProvision {
        initQuery.WriteString("r_provision = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnComplex1 {
        initQuery.WriteString("r_complex1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnComplex2 {
        initQuery.WriteString("r_complex2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnCompletionyear {
        initQuery.WriteString("r_completionyear = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnCompletionmonth {
        initQuery.WriteString("r_completionmonth = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnCompletionday {
        initQuery.WriteString("r_completionday = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnParcelrate {
        initQuery.WriteString("r_parcelrate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnPlanyears {
        initQuery.WriteString("r_planyears = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnInfo1 {
        initQuery.WriteString("r_info1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnInfo2 {
        initQuery.WriteString("r_info2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnInfo3 {
        initQuery.WriteString("r_info3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnInfo4 {
        initQuery.WriteString("r_info4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnInfo5 {
        initQuery.WriteString("r_info5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnInfo6 {
        initQuery.WriteString("r_info6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnInfo7 {
        initQuery.WriteString("r_info7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnInfo8 {
        initQuery.WriteString("r_info8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnInfo9 {
        initQuery.WriteString("r_info9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnInfo10 {
        initQuery.WriteString("r_info10 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnInfo11 {
        initQuery.WriteString("r_info11 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure1 {
        initQuery.WriteString("r_structure1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure2 {
        initQuery.WriteString("r_structure2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure3 {
        initQuery.WriteString("r_structure3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure4 {
        initQuery.WriteString("r_structure4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure5 {
        initQuery.WriteString("r_structure5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure6 {
        initQuery.WriteString("r_structure6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure7 {
        initQuery.WriteString("r_structure7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure8 {
        initQuery.WriteString("r_structure8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure9 {
        initQuery.WriteString("r_structure9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure10 {
        initQuery.WriteString("r_structure10 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure11 {
        initQuery.WriteString("r_structure11 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure12 {
        initQuery.WriteString("r_structure12 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure13 {
        initQuery.WriteString("r_structure13 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnStructure14 {
        initQuery.WriteString("r_structure14 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnReviewcontent1 {
        initQuery.WriteString("r_reviewcontent1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnReviewcontent2 {
        initQuery.WriteString("r_reviewcontent2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnReviewcontent3 {
        initQuery.WriteString("r_reviewcontent3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnReviewcontent4 {
        initQuery.WriteString("r_reviewcontent4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnReviewcontent5 {
        initQuery.WriteString("r_reviewcontent5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnReviewcontent6 {
        initQuery.WriteString("r_reviewcontent6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnReviewcontent7 {
        initQuery.WriteString("r_reviewcontent7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnSavingprice {
        initQuery.WriteString("r_savingprice = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnPrice1 {
        initQuery.WriteString("r_price1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnPrice2 {
        initQuery.WriteString("r_price2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnPrice3 {
        initQuery.WriteString("r_price3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnPrice4 {
        initQuery.WriteString("r_price4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnPrice5 {
        initQuery.WriteString("r_price5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnReportdate {
        initQuery.WriteString("r_reportdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnContent1 {
        initQuery.WriteString("r_content1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnContent2 {
        initQuery.WriteString("r_content2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnContent3 {
        initQuery.WriteString("r_content3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnContent4 {
        initQuery.WriteString("r_content4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnContent5 {
        initQuery.WriteString("r_content5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnContent6 {
        initQuery.WriteString("r_content6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnContent7 {
        initQuery.WriteString("r_content7 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnContent8 {
        initQuery.WriteString("r_content8 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnContent9 {
        initQuery.WriteString("r_content9 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnContent10 {
        initQuery.WriteString("r_content10 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnContent11 {
        initQuery.WriteString("r_content11 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnPeriodtype {
        initQuery.WriteString("r_periodtype = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnRemark {
        initQuery.WriteString("r_remark = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnApt {
        initQuery.WriteString("r_apt = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == repair.ColumnDate {
        initQuery.WriteString("r_date = ?")
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

func (p *RepairManager) UpdateType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_type = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStatus(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_status = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateCalculatetype(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_calculatetype = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateProvision(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_provision = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateComplex1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_complex1 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateComplex2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_complex2 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateCompletionyear(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_completionyear = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateCompletionmonth(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_completionmonth = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateCompletionday(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_completionday = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateParcelrate(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_parcelrate = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdatePlanyears(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_planyears = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateInfo1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_info1 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateInfo2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_info2 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateInfo3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_info3 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateInfo4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_info4 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateInfo5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_info5 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateInfo6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_info6 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateInfo7(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_info7 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateInfo8(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_info8 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateInfo9(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_info9 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateInfo10(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_info10 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateInfo11(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_info11 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure1 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure2 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure3 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure4 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure5 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure6 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure7(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure7 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure8(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure8 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure9(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure9 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure10(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure10 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure11(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure11 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure12(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure12 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure13(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure13 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateStructure14(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_structure14 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateReviewcontent1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_reviewcontent1 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateReviewcontent2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_reviewcontent2 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateReviewcontent3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_reviewcontent3 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateReviewcontent4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_reviewcontent4 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateReviewcontent5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_reviewcontent5 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateReviewcontent6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_reviewcontent6 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateReviewcontent7(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_reviewcontent7 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateSavingprice(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_savingprice = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdatePrice1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_price1 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdatePrice2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_price2 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdatePrice3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_price3 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdatePrice4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_price4 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdatePrice5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_price5 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateReportdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_reportdate = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateContent1(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_content1 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateContent2(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_content2 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateContent3(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_content3 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateContent4(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_content4 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateContent5(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_content5 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateContent6(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_content6 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateContent7(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_content7 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateContent8(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_content8 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateContent9(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_content9 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateContent10(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_content10 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateContent11(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_content11 = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdatePeriodtype(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_periodtype = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateRemark(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_remark = ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) UpdateApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_apt = ? where r_id = ?"
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

func (p *RepairManager) IncreaseType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_type = r_type + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) IncreaseStatus(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_status = r_status + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) IncreaseCalculatetype(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_calculatetype = r_calculatetype + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) IncreaseProvision(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_provision = r_provision + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) IncreaseCompletionyear(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_completionyear = r_completionyear + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) IncreaseCompletionmonth(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_completionmonth = r_completionmonth + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) IncreaseCompletionday(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_completionday = r_completionday + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) IncreaseParcelrate(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_parcelrate = r_parcelrate + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) IncreasePlanyears(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_planyears = r_planyears + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) IncreaseSavingprice(value Double, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_savingprice = r_savingprice + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairManager) IncreaseApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repair_tb set r_apt = r_apt + ? where r_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *RepairManager) GetIdentity() int64 {
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

func (p *Repair) InitExtra() {
    p.Extra = map[string]any{
            "periodtype":     repair.GetPeriodtype(p.Periodtype),

    }
}

func (p *RepairManager) ReadRow(rows *sql.Rows) *Repair {
    var item Repair
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Type, &item.Status, &item.Calculatetype, &item.Provision, &item.Complex1, &item.Complex2, &item.Completionyear, &item.Completionmonth, &item.Completionday, &item.Parcelrate, &item.Planyears, &item.Info1, &item.Info2, &item.Info3, &item.Info4, &item.Info5, &item.Info6, &item.Info7, &item.Info8, &item.Info9, &item.Info10, &item.Info11, &item.Structure1, &item.Structure2, &item.Structure3, &item.Structure4, &item.Structure5, &item.Structure6, &item.Structure7, &item.Structure8, &item.Structure9, &item.Structure10, &item.Structure11, &item.Structure12, &item.Structure13, &item.Structure14, &item.Reviewcontent1, &item.Reviewcontent2, &item.Reviewcontent3, &item.Reviewcontent4, &item.Reviewcontent5, &item.Reviewcontent6, &item.Reviewcontent7, &item.Savingprice, &item.Price1, &item.Price2, &item.Price3, &item.Price4, &item.Price5, &item.Reportdate, &item.Content1, &item.Content2, &item.Content3, &item.Content4, &item.Content5, &item.Content6, &item.Content7, &item.Content8, &item.Content9, &item.Content10, &item.Content11, &item.Periodtype, &item.Remark, &item.Apt, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *RepairManager) ReadRows(rows *sql.Rows) []Repair {
    items := make([]Repair, 0)

    for rows.Next() {
        var item Repair
        
    
        err := rows.Scan(&item.Id, &item.Type, &item.Status, &item.Calculatetype, &item.Provision, &item.Complex1, &item.Complex2, &item.Completionyear, &item.Completionmonth, &item.Completionday, &item.Parcelrate, &item.Planyears, &item.Info1, &item.Info2, &item.Info3, &item.Info4, &item.Info5, &item.Info6, &item.Info7, &item.Info8, &item.Info9, &item.Info10, &item.Info11, &item.Structure1, &item.Structure2, &item.Structure3, &item.Structure4, &item.Structure5, &item.Structure6, &item.Structure7, &item.Structure8, &item.Structure9, &item.Structure10, &item.Structure11, &item.Structure12, &item.Structure13, &item.Structure14, &item.Reviewcontent1, &item.Reviewcontent2, &item.Reviewcontent3, &item.Reviewcontent4, &item.Reviewcontent5, &item.Reviewcontent6, &item.Reviewcontent7, &item.Savingprice, &item.Price1, &item.Price2, &item.Price3, &item.Price4, &item.Price5, &item.Reportdate, &item.Content1, &item.Content2, &item.Content3, &item.Content4, &item.Content5, &item.Content6, &item.Content7, &item.Content8, &item.Content9, &item.Content10, &item.Content11, &item.Periodtype, &item.Remark, &item.Apt, &item.Date)
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

func (p *RepairManager) Get(id int64) *Repair {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and r_id = ?")

    
    
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

func (p *RepairManager) GetWhere(args []any) *Repair {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *RepairManager) Count(args []any) int {
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

func (p *RepairManager) FindAll() []Repair {
    return p.Find(nil)
}

func (p *RepairManager) Find(args []any) []Repair {
    if !p.Conn.IsConnect() {
        items := make([]Repair, 0)
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
                query.WriteString(" and r_")
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
            orderby = "r_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "r_" + orderby
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
            orderby = "r_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "r_" + orderby
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
        items := make([]Repair, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *RepairManager) UpdateStatusById(status int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "update repair_tb set r_status = ? where 1=1 and r_id = ?"
	_, err := p.Exec(query, status, id)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err    
}




func (p *RepairManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and r_")
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
    
    query.WriteString(" group by r_")
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



func (p *RepairManager) MakeMap(items []Repair) map[int64]Repair {
     ret := make(map[int64]Repair)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *RepairManager) FindToMap(args []any) map[int64]Repair {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *RepairManager) FindAllToMap() map[int64]Repair {
     items := p.Find(nil)
     return p.MakeMap(items)
}


