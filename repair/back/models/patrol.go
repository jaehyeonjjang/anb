package models

import (
    "repair/models/patrol"
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

type Patrol struct {
            
    Id                int64 `json:"id"`         
    Location                string `json:"location"`         
    Content                string `json:"content"`         
    Process                string `json:"process"`         
    Opinion                string `json:"opinion"`         
    Status                patrol.Status `json:"status"`         
    User                int64 `json:"user"`         
    Apt                int64 `json:"apt"`         
    Startdate                string `json:"startdate"`         
    Enddate                string `json:"enddate"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type PatrolManager struct {
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



func (c *Patrol) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewPatrolManager(conn *Connection) *PatrolManager {
    var item PatrolManager


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

func (p *PatrolManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *PatrolManager) SetIndex(index string) {
    p.Index = index
}

func (p *PatrolManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *PatrolManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *PatrolManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *PatrolManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *PatrolManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select p_id, p_location, p_content, p_process, p_opinion, p_status, p_user, p_apt, p_startdate, p_enddate, p_date, a_id, a_name, a_completeyear, a_flatcount, a_type, a_floor, a_familycount, a_familycount1, a_familycount2, a_familycount3, a_tel, a_fax, a_email, a_personalemail, a_personalname, a_personalhp, a_zip, a_address, a_address2, a_contracttype, a_contractprice, a_testdate, a_nexttestdate, a_repair, a_safety, a_fault, a_contractdate, a_contractduration, a_invoice, a_depositdate, a_fmsloginid, a_fmspasswd, a_facilitydivision, a_facilitycategory, a_position, a_area, a_groundfloor, a_undergroundfloor, a_useapproval, a_date, u_id, u_loginid, u_passwd, u_name, u_email, u_level, u_apt, u_date from patrol_tb, apt_tb, user_tb")

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
    
    ret.WriteString(" and p_apt = a_id ")
    
    ret.WriteString(" and p_user = u_id ")
    

    return ret.String()
}

func (p *PatrolManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from patrol_tb, apt_tb, user_tb")

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
    
    ret.WriteString(" and p_apt = a_id ")
    
    ret.WriteString(" and p_user = u_id ")
    

    return ret.String()
}

func (p *PatrolManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select p_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from patrol_tb, apt_tb, user_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    
    ret.WriteString(" and p_apt = a_id ")
    
    ret.WriteString(" and p_user = u_id ")
    


    return ret.String()
}

func (p *PatrolManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate patrol_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *PatrolManager) Insert(item *Patrol) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    if item.Date == "" {
        t := time.Now().UTC().Add(time.Hour * 9)
        //t := time.Now()
        item.Date = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    
	
	
	
	
	
	
	
	
    if item.Startdate == "" {
       item.Startdate = "1000-01-01 00:00:00"
    }
	
    if item.Enddate == "" {
       item.Enddate = "1000-01-01 00:00:00"
    }
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    query := ""
    var res sql.Result
    var err error
    if item.Id > 0 {
        if config.Database.Type == config.Postgresql {
          query = "insert into patrol_tb (p_id, p_location, p_content, p_process, p_opinion, p_status, p_user, p_apt, p_startdate, p_enddate, p_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
        } else {
          query = "insert into patrol_tb (p_id, p_location, p_content, p_process, p_opinion, p_status, p_user, p_apt, p_startdate, p_enddate, p_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Location, item.Content, item.Process, item.Opinion, item.Status, item.User, item.Apt, item.Startdate, item.Enddate, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into patrol_tb (p_location, p_content, p_process, p_opinion, p_status, p_user, p_apt, p_startdate, p_enddate, p_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
        } else {
          query = "insert into patrol_tb (p_location, p_content, p_process, p_opinion, p_status, p_user, p_apt, p_startdate, p_enddate, p_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Location, item.Content, item.Process, item.Opinion, item.Status, item.User, item.Apt, item.Startdate, item.Enddate, item.Date)
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

func (p *PatrolManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from patrol_tb where p_id = ")
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

func (p *PatrolManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from patrol_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *PatrolManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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
                query.WriteString(" and p_")
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

func (p *PatrolManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from patrol_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *PatrolManager) Update(item *Patrol) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
    if item.Startdate == "" {
       item.Startdate = "1000-01-01 00:00:00"
    }
	
    if item.Enddate == "" {
       item.Enddate = "1000-01-01 00:00:00"
    }
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update patrol_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" p_location = $1, p_content = $2, p_process = $3, p_opinion = $4, p_status = $5, p_user = $6, p_apt = $7, p_startdate = $8, p_enddate = $9, p_date = $10 where p_id = $11")
    } else {
        query.WriteString(" p_location = ?, p_content = ?, p_process = ?, p_opinion = ?, p_status = ?, p_user = ?, p_apt = ?, p_startdate = ?, p_enddate = ?, p_date = ? where p_id = ?")
    }

	_, err := p.Exec(query.String() , item.Location, item.Content, item.Process, item.Opinion, item.Status, item.User, item.Apt, item.Startdate, item.Enddate, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *PatrolManager) UpdateWhere(columns []patrol.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update patrol_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == patrol.ColumnId {
        initQuery.WriteString("p_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == patrol.ColumnLocation {
        initQuery.WriteString("p_location = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == patrol.ColumnContent {
        initQuery.WriteString("p_content = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == patrol.ColumnProcess {
        initQuery.WriteString("p_process = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == patrol.ColumnOpinion {
        initQuery.WriteString("p_opinion = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == patrol.ColumnStatus {
        initQuery.WriteString("p_status = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == patrol.ColumnUser {
        initQuery.WriteString("p_user = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == patrol.ColumnApt {
        initQuery.WriteString("p_apt = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == patrol.ColumnStartdate {
        initQuery.WriteString("p_startdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == patrol.ColumnEnddate {
        initQuery.WriteString("p_enddate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == patrol.ColumnDate {
        initQuery.WriteString("p_date = ?")
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

func (p *PatrolManager) UpdateLocation(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update patrol_tb set p_location = ? where p_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PatrolManager) UpdateContent(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update patrol_tb set p_content = ? where p_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PatrolManager) UpdateProcess(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update patrol_tb set p_process = ? where p_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PatrolManager) UpdateOpinion(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update patrol_tb set p_opinion = ? where p_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PatrolManager) UpdateStatus(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update patrol_tb set p_status = ? where p_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PatrolManager) UpdateUser(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update patrol_tb set p_user = ? where p_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PatrolManager) UpdateApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update patrol_tb set p_apt = ? where p_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PatrolManager) UpdateStartdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update patrol_tb set p_startdate = ? where p_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PatrolManager) UpdateEnddate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update patrol_tb set p_enddate = ? where p_id = ?"
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

func (p *PatrolManager) IncreaseUser(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update patrol_tb set p_user = p_user + ? where p_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *PatrolManager) IncreaseApt(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update patrol_tb set p_apt = p_apt + ? where p_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *PatrolManager) GetIdentity() int64 {
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

func (p *Patrol) InitExtra() {
    p.Extra = map[string]any{
            "status":     patrol.GetStatus(p.Status),

    }
}

func (p *PatrolManager) ReadRow(rows *sql.Rows) *Patrol {
    var item Patrol
    var err error

    var _apt Apt
    var _user User
    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Location, &item.Content, &item.Process, &item.Opinion, &item.Status, &item.User, &item.Apt, &item.Startdate, &item.Enddate, &item.Date, &_apt.Id, &_apt.Name, &_apt.Completeyear, &_apt.Flatcount, &_apt.Type, &_apt.Floor, &_apt.Familycount, &_apt.Familycount1, &_apt.Familycount2, &_apt.Familycount3, &_apt.Tel, &_apt.Fax, &_apt.Email, &_apt.Personalemail, &_apt.Personalname, &_apt.Personalhp, &_apt.Zip, &_apt.Address, &_apt.Address2, &_apt.Contracttype, &_apt.Contractprice, &_apt.Testdate, &_apt.Nexttestdate, &_apt.Repair, &_apt.Safety, &_apt.Fault, &_apt.Contractdate, &_apt.Contractduration, &_apt.Invoice, &_apt.Depositdate, &_apt.Fmsloginid, &_apt.Fmspasswd, &_apt.Facilitydivision, &_apt.Facilitycategory, &_apt.Position, &_apt.Area, &_apt.Groundfloor, &_apt.Undergroundfloor, &_apt.Useapproval, &_apt.Date, &_user.Id, &_user.Loginid, &_user.Passwd, &_user.Name, &_user.Email, &_user.Level, &_user.Apt, &_user.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        if item.Startdate == "0000-00-00 00:00:00" || item.Startdate == "1000-01-01 00:00:00" || item.Startdate == "9999-01-01 00:00:00" {
            item.Startdate = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Startdate = strings.ReplaceAll(strings.ReplaceAll(item.Startdate, "T", " "), "Z", "")
        }
		
        
        if item.Enddate == "0000-00-00 00:00:00" || item.Enddate == "1000-01-01 00:00:00" || item.Enddate == "9999-01-01 00:00:00" {
            item.Enddate = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Enddate = strings.ReplaceAll(strings.ReplaceAll(item.Enddate, "T", " "), "Z", "")
        }
		
        
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
        _apt.InitExtra()
        item.AddExtra("apt",  _apt)
_user.InitExtra()
        item.AddExtra("user",  _user)

        return &item
    }
}

func (p *PatrolManager) ReadRows(rows *sql.Rows) []Patrol {
    items := make([]Patrol, 0)

    for rows.Next() {
        var item Patrol
        var _apt Apt
            var _user User
            
    
        err := rows.Scan(&item.Id, &item.Location, &item.Content, &item.Process, &item.Opinion, &item.Status, &item.User, &item.Apt, &item.Startdate, &item.Enddate, &item.Date, &_apt.Id, &_apt.Name, &_apt.Completeyear, &_apt.Flatcount, &_apt.Type, &_apt.Floor, &_apt.Familycount, &_apt.Familycount1, &_apt.Familycount2, &_apt.Familycount3, &_apt.Tel, &_apt.Fax, &_apt.Email, &_apt.Personalemail, &_apt.Personalname, &_apt.Personalhp, &_apt.Zip, &_apt.Address, &_apt.Address2, &_apt.Contracttype, &_apt.Contractprice, &_apt.Testdate, &_apt.Nexttestdate, &_apt.Repair, &_apt.Safety, &_apt.Fault, &_apt.Contractdate, &_apt.Contractduration, &_apt.Invoice, &_apt.Depositdate, &_apt.Fmsloginid, &_apt.Fmspasswd, &_apt.Facilitydivision, &_apt.Facilitycategory, &_apt.Position, &_apt.Area, &_apt.Groundfloor, &_apt.Undergroundfloor, &_apt.Useapproval, &_apt.Date, &_user.Id, &_user.Loginid, &_user.Passwd, &_user.Name, &_user.Email, &_user.Level, &_user.Apt, &_user.Date)
        if err != nil {
           if p.Log {
             log.Error().Str("error", err.Error()).Msg("SQL")
           }
           break
        }

        
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        
		
        if item.Startdate == "0000-00-00 00:00:00" || item.Startdate == "1000-01-01 00:00:00" || item.Startdate == "9999-01-01 00:00:00" {
            item.Startdate = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Startdate = strings.ReplaceAll(strings.ReplaceAll(item.Startdate, "T", " "), "Z", "")
        }
		
		
        if item.Enddate == "0000-00-00 00:00:00" || item.Enddate == "1000-01-01 00:00:00" || item.Enddate == "9999-01-01 00:00:00" {
            item.Enddate = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Enddate = strings.ReplaceAll(strings.ReplaceAll(item.Enddate, "T", " "), "Z", "")
        }
		
		
        if item.Date == "0000-00-00 00:00:00" || item.Date == "1000-01-01 00:00:00" || item.Date == "9999-01-01 00:00:00" {
            item.Date = ""
        }

        if config.Database.Type == config.Postgresql {
            item.Date = strings.ReplaceAll(strings.ReplaceAll(item.Date, "T", " "), "Z", "")
        }
		
		
        
        item.InitExtra()        
        _apt.InitExtra()
        item.AddExtra("apt",  _apt)
_user.InitExtra()
        item.AddExtra("user",  _user)

        items = append(items, item)
    }


     return items
}

func (p *PatrolManager) Get(id int64) *Patrol {
    if !p.Conn.IsConnect() {
        return nil
    }

    var query strings.Builder
    query.WriteString(p.GetQuery())
    query.WriteString(" and p_id = ?")

    
    query.WriteString(" and p_apt = a_id ")
    
    query.WriteString(" and p_user = u_id ")
    
    
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

func (p *PatrolManager) GetWhere(args []any) *Patrol {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *PatrolManager) Count(args []any) int {
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

func (p *PatrolManager) FindAll() []Patrol {
    return p.Find(nil)
}

func (p *PatrolManager) Find(args []any) []Patrol {
    if !p.Conn.IsConnect() {
        items := make([]Patrol, 0)
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
                query.WriteString(" and p_")
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
            orderby = "p_id desc"
        } else {
            if !strings.Contains(orderby, "_") {                   
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "p_" + orderby
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
            orderby = "p_id"
        } else {
            if !strings.Contains(orderby, "_") {
                if strings.ToUpper(orderby) != "RAND()" {
                  orderby = "p_" + orderby
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
        items := make([]Patrol, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}


func (p *PatrolManager) CountByApt(apt int64, args ...any) int {
    rets := make([]any, 0)
    rets = append(rets, args...)
    
    if apt != 0 { 
        rets = append(rets, Where{Column:"apt", Value:apt, Compare:"="})
     }
    
    return p.Count(rets)
}

func (p *PatrolManager) FindByApt(apt int64, args ...any) []Patrol {
    rets := make([]any, 0)
    rets = append(rets, args...)

    if apt != 0 { 
        rets = append(rets, Where{Column:"apt", Value:apt, Compare:"="})
     }
    
    
    return p.Find(rets)
}




func (p *PatrolManager) GroupBy(name string, args []any) []Groupby {
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
                query.WriteString(" and p_")
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
    
    query.WriteString(" group by p_")
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



func (p *PatrolManager) MakeMap(items []Patrol) map[int64]Patrol {
     ret := make(map[int64]Patrol)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *PatrolManager) FindToMap(args []any) map[int64]Patrol {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *PatrolManager) FindAllToMap() map[int64]Patrol {
     items := p.Find(nil)
     return p.MakeMap(items)
}


