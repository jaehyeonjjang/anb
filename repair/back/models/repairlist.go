package models

import (
    "repair/models/repairlist"
    "database/sql"
    "errors"
    "fmt"
    "strings"
    
    "repair/global/config"
    log "repair/global/log"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"

    
)

type Repairlist struct {
            
    Id                int64 `json:"id"`         
    Name                string `json:"name"`         
    Completeyear                string `json:"completeyear"`         
    Flatcount                string `json:"flatcount"`         
    Type                string `json:"type"`         
    Floor                string `json:"floor"`         
    Familycount                string `json:"familycount"`         
    Familycount1                int `json:"familycount1"`         
    Familycount2                int `json:"familycount2"`         
    Familycount3                int `json:"familycount3"`         
    Tel                string `json:"tel"`         
    Fax                string `json:"fax"`         
    Email                string `json:"email"`         
    Personalemail                string `json:"personalemail"`         
    Personalname                string `json:"personalname"`         
    Personalhp                string `json:"personalhp"`         
    Zip                string `json:"zip"`         
    Address                string `json:"address"`         
    Address2                string `json:"address2"`         
    Contracttype                int `json:"contracttype"`         
    Contractprice                string `json:"contractprice"`         
    Testdate                string `json:"testdate"`         
    Nexttestdate                string `json:"nexttestdate"`         
    Repair                string `json:"repair"`         
    Safety                string `json:"safety"`         
    Fault                string `json:"fault"`         
    Contractdate                string `json:"contractdate"`         
    Contractduration                string `json:"contractduration"`         
    Invoice                string `json:"invoice"`         
    Depositdate                string `json:"depositdate"`         
    Fmsloginid                string `json:"fmsloginid"`         
    Fmspasswd                string `json:"fmspasswd"`         
    Facilitydivision                int `json:"facilitydivision"`         
    Facilitycategory                int `json:"facilitycategory"`         
    Position                string `json:"position"`         
    Area                string `json:"area"`         
    Groundfloor                int `json:"groundfloor"`         
    Undergroundfloor                int `json:"undergroundfloor"`         
    Useapproval                string `json:"useapproval"`         
    Date                string `json:"date"`         
    Repairid                int64 `json:"repairid"`         
    Repairtype                repairlist.Repairtype `json:"repairtype"`         
    Reportdate                string `json:"reportdate"`         
    Repairdate                string `json:"repairdate"`         
    Info1                string `json:"info1"`         
    Status                int `json:"status"` 
    
    Extra                    map[string]any `json:"extra"`
}




type RepairlistManager struct {
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



func (c *Repairlist) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewRepairlistManager(conn *Connection) *RepairlistManager {
    var item RepairlistManager


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

func (p *RepairlistManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *RepairlistManager) SetIndex(index string) {
    p.Index = index
}

func (p *RepairlistManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *RepairlistManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *RepairlistManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *RepairlistManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *RepairlistManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select a_id, a_name, a_completeyear, a_flatcount, a_type, a_floor, a_familycount, a_familycount1, a_familycount2, a_familycount3, a_tel, a_fax, a_email, a_personalemail, a_personalname, a_personalhp, a_zip, a_address, a_address2, a_contracttype, a_contractprice, a_testdate, a_nexttestdate, a_repair, a_safety, a_fault, a_contractdate, a_contractduration, a_invoice, a_depositdate, a_fmsloginid, a_fmspasswd, a_facilitydivision, a_facilitycategory, a_position, a_area, a_groundfloor, a_undergroundfloor, a_useapproval, a_date, a_repairid, a_repairtype, a_reportdate, a_repairdate, a_info1, a_status from repairlist_vw")

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

func (p *RepairlistManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from repairlist_vw")

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

func (p *RepairlistManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select a_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from repairlist_vw ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    


    return ret.String()
}

func (p *RepairlistManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate repairlist_vw "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}



func (p *RepairlistManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from repairlist_vw where a_id = ")
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

func (p *RepairlistManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from repairlist_vw"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *RepairlistManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and a_")
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

func (p *RepairlistManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from repairlist_vw where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}


/*

func (p *RepairlistManager) IncreaseFamilycount1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairlist_vw set a_familycount1 = a_familycount1 + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairlistManager) IncreaseFamilycount2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairlist_vw set a_familycount2 = a_familycount2 + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairlistManager) IncreaseFamilycount3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairlist_vw set a_familycount3 = a_familycount3 + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairlistManager) IncreaseContracttype(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairlist_vw set a_contracttype = a_contracttype + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairlistManager) IncreaseFacilitydivision(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairlist_vw set a_facilitydivision = a_facilitydivision + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairlistManager) IncreaseFacilitycategory(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairlist_vw set a_facilitycategory = a_facilitycategory + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairlistManager) IncreaseGroundfloor(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairlist_vw set a_groundfloor = a_groundfloor + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairlistManager) IncreaseUndergroundfloor(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairlist_vw set a_undergroundfloor = a_undergroundfloor + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairlistManager) IncreaseRepairid(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairlist_vw set a_repairid = a_repairid + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *RepairlistManager) IncreaseStatus(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update repairlist_vw set a_status = a_status + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *RepairlistManager) GetIdentity() int64 {
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

func (p *Repairlist) InitExtra() {
    p.Extra = map[string]any{
            "repairtype":     repairlist.GetRepairtype(p.Repairtype),

    }
}

func (p *RepairlistManager) ReadRow(rows *sql.Rows) *Repairlist {
    var item Repairlist
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Name, &item.Completeyear, &item.Flatcount, &item.Type, &item.Floor, &item.Familycount, &item.Familycount1, &item.Familycount2, &item.Familycount3, &item.Tel, &item.Fax, &item.Email, &item.Personalemail, &item.Personalname, &item.Personalhp, &item.Zip, &item.Address, &item.Address2, &item.Contracttype, &item.Contractprice, &item.Testdate, &item.Nexttestdate, &item.Repair, &item.Safety, &item.Fault, &item.Contractdate, &item.Contractduration, &item.Invoice, &item.Depositdate, &item.Fmsloginid, &item.Fmspasswd, &item.Facilitydivision, &item.Facilitycategory, &item.Position, &item.Area, &item.Groundfloor, &item.Undergroundfloor, &item.Useapproval, &item.Date, &item.Repairid, &item.Repairtype, &item.Reportdate, &item.Repairdate, &item.Info1, &item.Status)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" || item.Date == "9999-01-01 00:00:00" {
            item.Date = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Date = strings.ReplaceAll(strings.ReplaceAll(item.Date, "T", " "), "Z", "")
        }
		
        
        
        
        
        
        
        
        if item.Repairdate == "0000-00-00 00:00:00" || item.Repairdate == "1000-01-01 00:00:00" || item.Repairdate == "9999-01-01 00:00:00" {
            item.Repairdate = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Repairdate = strings.ReplaceAll(strings.ReplaceAll(item.Repairdate, "T", " "), "Z", "")
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

func (p *RepairlistManager) ReadRows(rows *sql.Rows) []Repairlist {
    items := make([]Repairlist, 0)

    for rows.Next() {
        var item Repairlist
        
    
        err := rows.Scan(&item.Id, &item.Name, &item.Completeyear, &item.Flatcount, &item.Type, &item.Floor, &item.Familycount, &item.Familycount1, &item.Familycount2, &item.Familycount3, &item.Tel, &item.Fax, &item.Email, &item.Personalemail, &item.Personalname, &item.Personalhp, &item.Zip, &item.Address, &item.Address2, &item.Contracttype, &item.Contractprice, &item.Testdate, &item.Nexttestdate, &item.Repair, &item.Safety, &item.Fault, &item.Contractdate, &item.Contractduration, &item.Invoice, &item.Depositdate, &item.Fmsloginid, &item.Fmspasswd, &item.Facilitydivision, &item.Facilitycategory, &item.Position, &item.Area, &item.Groundfloor, &item.Undergroundfloor, &item.Useapproval, &item.Date, &item.Repairid, &item.Repairtype, &item.Reportdate, &item.Repairdate, &item.Info1, &item.Status)
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
		
		
        
		
        
		
        
		
        if item.Repairdate == "0000-00-00 00:00:00" || item.Repairdate == "1000-01-01 00:00:00" || item.Repairdate == "9999-01-01 00:00:00" {
            item.Repairdate = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Repairdate = strings.ReplaceAll(strings.ReplaceAll(item.Repairdate, "T", " "), "Z", "")
        }
		
		
        
		
        
		
        
        item.InitExtra()        
        
        items = append(items, item)
    }


     return items
}

func (p *RepairlistManager) Get(id int64) *Repairlist {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and a_id = ?")

    
    
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

func (p *RepairlistManager) GetWhere(args []any) *Repairlist {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *RepairlistManager) Count(args []any) int {
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

func (p *RepairlistManager) FindAll() []Repairlist {
    return p.Find(nil)
}

func (p *RepairlistManager) Find(args []any) []Repairlist {
    if !p.Conn.IsConnect() {
        items := make([]Repairlist, 0)
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
                query.WriteString(" and a_")
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
            orderby = "a_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "a_" + orderby
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
            orderby = "a_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "a_" + orderby
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
        items := make([]Repairlist, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *RepairlistManager) UpdateStatusById(status int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "update repairlist_vw set a_status = ? where 1=1 and a_id = ?"
	_, err := p.Exec(query, status, id)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err    
}

func (p *RepairlistManager) CountByRepairtypes(repairtype []repairlist.Repairtype, args ...any) int {
    rets := make([]any, 0)
    rets = append(rets, args...)
    
    
        rets = append(rets, Where{Column:"repairtype", Value:repairtype, Compare:"in"})
    
    
    return p.Count(rets)
}

func (p *RepairlistManager) FindByRepairtypes(repairtype []repairlist.Repairtype, args ...any) []Repairlist {
    rets := make([]any, 0)
    rets = append(rets, args...)

    
        rets = append(rets, Where{Column:"repairtype", Value:repairtype, Compare:"in"})
    
    
    
    return p.Find(rets)
}




func (p *RepairlistManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and a_")
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
    
    query.WriteString(" group by a_")
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



func (p *RepairlistManager) MakeMap(items []Repairlist) map[int64]Repairlist {
     ret := make(map[int64]Repairlist)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *RepairlistManager) FindToMap(args []any) map[int64]Repairlist {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *RepairlistManager) FindAllToMap() map[int64]Repairlist {
     items := p.Find(nil)
     return p.MakeMap(items)
}

func (p *RepairlistManager) MakeNameMap(items []Repairlist) map[string]Repairlist {
     ret := make(map[string]Repairlist)
     for _, v := range items {
        ret[v.Name] = v
     }

     return ret
}

func (p *RepairlistManager) FindToNameMap(args []any) map[string]Repairlist {
     items := p.Find(args)
     return p.MakeNameMap(items)
}

func (p *RepairlistManager) FindAllToNameMap() map[string]Repairlist {
     items := p.Find(nil)
     return p.MakeNameMap(items)
}
