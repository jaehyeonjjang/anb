package models

import (
    "repair/models/oldapt"
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

type Oldapt struct {
            
    Id                int64 `json:"id"`         
    Aptgroup                int64 `json:"aptgroup"`         
    Name                string `json:"name"`         
    Workstartdate                string `json:"workstartdate"`         
    Workenddate                string `json:"workenddate"`         
    Type                int `json:"type"`         
    Master                int64 `json:"master"`         
    Status                int `json:"status"`         
    Company                int64 `json:"company"`         
    Report                int `json:"report"`         
    Report1                int `json:"report1"`         
    Report2                int `json:"report2"`         
    Report3                int `json:"report3"`         
    Report4                int `json:"report4"`         
    Report5                int `json:"report5"`         
    Report6                int `json:"report6"`         
    Summarytype                int `json:"summarytype"`         
    Search                string `json:"search"`         
    User                int64 `json:"user"`         
    Updateuser                int64 `json:"updateuser"`         
    Date                string `json:"date"` 
    
    Extra                    map[string]any `json:"extra"`
}




type OldaptManager struct {
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



func (c *Oldapt) AddExtra(key string, value any) {    
	c.Extra[key] = value     
}

func NewOldaptManager(conn *Connection) *OldaptManager {
    var item OldaptManager


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

func (p *OldaptManager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *OldaptManager) SetIndex(index string) {
    p.Index = index
}

func (p *OldaptManager) SetCountQuery(query string) {
    p.CountQuery = query
}

func (p *OldaptManager) SetSelectQuery(query string) {
    p.SelectQuery = query
}

func (p *OldaptManager) Exec(query string, params ...any) (sql.Result, error) {
    if p.Log {
       if len(params) > 0 {
	       log.Debug().Str("query", query).Any("param", params).Msg("SQL")
       } else {
	       log.Debug().Str("query", query).Msg("SQL")
       }
    }

    return p.Conn.Exec(query, params...)
}

func (p *OldaptManager) Query(query string, params ...any) (*sql.Rows, error) {
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

func (p *OldaptManager) GetQuery() string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder

    ret.WriteString("select a_id, a_aptgroup, a_name, a_workstartdate, a_workenddate, a_type, a_master, a_status, a_company, a_report, a_report1, a_report2, a_report3, a_report4, a_report5, a_report6, a_summarytype, a_search, a_user, a_updateuser, a_date from oldapt_tb")

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

func (p *OldaptManager) GetQuerySelect() string {
    if p.CountQuery != "" {
        return p.CountQuery    
    }

    var ret strings.Builder
    
    ret.WriteString("select count(*) from oldapt_tb")

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

func (p *OldaptManager) GetQueryGroup(name string) string {
    if p.SelectQuery != "" {
        return p.SelectQuery    
    }

    var ret strings.Builder
    ret.WriteString("select a_")
    ret.WriteString(name)
    ret.WriteString(", count(*) from oldapt_tb ")

    if p.Index != "" {
        ret.WriteString(" use index(")
        ret.WriteString(p.Index)
        ret.WriteString(")")
    }

    ret.WriteString(" where 1=1 ")
    


    return ret.String()
}

func (p *OldaptManager) Truncate() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    query := "truncate oldapt_tb "
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return nil
}

func (p *OldaptManager) Insert(item *Oldapt) error {
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
          query = "insert into oldapt_tb (a_id, a_aptgroup, a_name, a_workstartdate, a_workenddate, a_type, a_master, a_status, a_company, a_report, a_report1, a_report2, a_report3, a_report4, a_report5, a_report6, a_summarytype, a_search, a_user, a_updateuser, a_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)"
        } else {
          query = "insert into oldapt_tb (a_id, a_aptgroup, a_name, a_workstartdate, a_workenddate, a_type, a_master, a_status, a_company, a_report, a_report1, a_report2, a_report3, a_report4, a_report5, a_report6, a_summarytype, a_search, a_user, a_updateuser, a_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Id, item.Aptgroup, item.Name, item.Workstartdate, item.Workenddate, item.Type, item.Master, item.Status, item.Company, item.Report, item.Report1, item.Report2, item.Report3, item.Report4, item.Report5, item.Report6, item.Summarytype, item.Search, item.User, item.Updateuser, item.Date)
    } else {
        if config.Database.Type == config.Postgresql {
          query = "insert into oldapt_tb (a_aptgroup, a_name, a_workstartdate, a_workenddate, a_type, a_master, a_status, a_company, a_report, a_report1, a_report2, a_report3, a_report4, a_report5, a_report6, a_summarytype, a_search, a_user, a_updateuser, a_date) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)"
        } else {
          query = "insert into oldapt_tb (a_aptgroup, a_name, a_workstartdate, a_workenddate, a_type, a_master, a_status, a_company, a_report, a_report1, a_report2, a_report3, a_report4, a_report5, a_report6, a_summarytype, a_search, a_user, a_updateuser, a_date) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }
        res, err = p.Exec(query , item.Aptgroup, item.Name, item.Workstartdate, item.Workenddate, item.Type, item.Master, item.Status, item.Company, item.Report, item.Report1, item.Report2, item.Report3, item.Report4, item.Report5, item.Report6, item.Summarytype, item.Search, item.User, item.Updateuser, item.Date)
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

func (p *OldaptManager) Delete(id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var query strings.Builder
    
    query.WriteString("delete from oldapt_tb where a_id = ")
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

func (p *OldaptManager) DeleteAll() error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query := "delete from oldapt_tb"
    _, err := p.Exec(query)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    return err
}

func (p *OldaptManager) MakeQuery(initQuery string , postQuery string, initParams []any, args []any) (string, []any) {
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

func (p *OldaptManager) DeleteWhere(args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    query, params := p.MakeQuery("delete from oldapt_tb where 1=1", "", nil, args)
    _, err := p.Exec(query, params...)

    if err != nil {
       if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
       }
    }

    
    return err
}

func (p *OldaptManager) Update(item *Oldapt) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }
    
    
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    if item.Date == "" {
       item.Date = "1000-01-01 00:00:00"
    }
	

    var query strings.Builder
	query.WriteString("update oldapt_tb set ")
    if config.Database.Type == config.Postgresql {
        query.WriteString(" a_aptgroup = $1, a_name = $2, a_workstartdate = $3, a_workenddate = $4, a_type = $5, a_master = $6, a_status = $7, a_company = $8, a_report = $9, a_report1 = $10, a_report2 = $11, a_report3 = $12, a_report4 = $13, a_report5 = $14, a_report6 = $15, a_summarytype = $16, a_search = $17, a_user = $18, a_updateuser = $19, a_date = $20 where a_id = $21")
    } else {
        query.WriteString(" a_aptgroup = ?, a_name = ?, a_workstartdate = ?, a_workenddate = ?, a_type = ?, a_master = ?, a_status = ?, a_company = ?, a_report = ?, a_report1 = ?, a_report2 = ?, a_report3 = ?, a_report4 = ?, a_report5 = ?, a_report6 = ?, a_summarytype = ?, a_search = ?, a_user = ?, a_updateuser = ?, a_date = ? where a_id = ?")
    }

	_, err := p.Exec(query.String() , item.Aptgroup, item.Name, item.Workstartdate, item.Workenddate, item.Type, item.Master, item.Status, item.Company, item.Report, item.Report1, item.Report2, item.Report3, item.Report4, item.Report5, item.Report6, item.Summarytype, item.Search, item.User, item.Updateuser, item.Date, item.Id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }
    
        
    return err
}


func (p *OldaptManager) UpdateWhere(columns []oldapt.Params, args []any) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

    var initQuery strings.Builder
    var initParams []any

    initQuery.WriteString("update oldapt_tb set ")
    for i, v := range columns {
        if i > 0 {
            initQuery.WriteString(", ")
        }

        if v.Column == oldapt.ColumnId {
        initQuery.WriteString("a_id = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnAptgroup {
        initQuery.WriteString("a_aptgroup = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnName {
        initQuery.WriteString("a_name = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnWorkstartdate {
        initQuery.WriteString("a_workstartdate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnWorkenddate {
        initQuery.WriteString("a_workenddate = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnType {
        initQuery.WriteString("a_type = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnMaster {
        initQuery.WriteString("a_master = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnStatus {
        initQuery.WriteString("a_status = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnCompany {
        initQuery.WriteString("a_company = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnReport {
        initQuery.WriteString("a_report = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnReport1 {
        initQuery.WriteString("a_report1 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnReport2 {
        initQuery.WriteString("a_report2 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnReport3 {
        initQuery.WriteString("a_report3 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnReport4 {
        initQuery.WriteString("a_report4 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnReport5 {
        initQuery.WriteString("a_report5 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnReport6 {
        initQuery.WriteString("a_report6 = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnSummarytype {
        initQuery.WriteString("a_summarytype = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnSearch {
        initQuery.WriteString("a_search = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnUser {
        initQuery.WriteString("a_user = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnUpdateuser {
        initQuery.WriteString("a_updateuser = ?")
        initParams = append(initParams, v.Value)
         } else if v.Column == oldapt.ColumnDate {
        initQuery.WriteString("a_date = ?")
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

func (p *OldaptManager) UpdateAptgroup(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_aptgroup = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateName(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_name = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateWorkstartdate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_workstartdate = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateWorkenddate(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_workenddate = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_type = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateMaster(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_master = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateStatus(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_status = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateCompany(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_company = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateReport(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateReport1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report1 = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateReport2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report2 = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateReport3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report3 = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateReport4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report4 = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateReport5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report5 = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateReport6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report6 = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateSummarytype(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_summarytype = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateSearch(value string, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_search = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateUser(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_user = ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) UpdateUpdateuser(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_updateuser = ? where a_id = ?"
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

func (p *OldaptManager) IncreaseAptgroup(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_aptgroup = a_aptgroup + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseType(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_type = a_type + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseMaster(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_master = a_master + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseStatus(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_status = a_status + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseCompany(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_company = a_company + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseReport(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report = a_report + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseReport1(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report1 = a_report1 + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseReport2(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report2 = a_report2 + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseReport3(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report3 = a_report3 + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseReport4(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report4 = a_report4 + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseReport5(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report5 = a_report5 + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseReport6(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_report6 = a_report6 + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseSummarytype(value int, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_summarytype = a_summarytype + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseUser(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_user = a_user + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

func (p *OldaptManager) IncreaseUpdateuser(value int64, id int64) error {
    if !p.Conn.IsConnect() {
        return errors.New("Connection Error")
    }

	query := "update oldapt_tb set a_updateuser = a_updateuser + ? where a_id = ?"
	_, err := p.Exec(query, value, id)

    if err != nil {
        if p.Log {
          log.Error().Str("error", err.Error()).Msg("SQL")
        }
    }

    return err
}

*/

func (p *OldaptManager) GetIdentity() int64 {
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

func (p *Oldapt) InitExtra() {
    p.Extra = map[string]any{

    }
}

func (p *OldaptManager) ReadRow(rows *sql.Rows) *Oldapt {
    var item Oldapt
    var err error

    

    if rows.Next() {
        err = rows.Scan(&item.Id, &item.Aptgroup, &item.Name, &item.Workstartdate, &item.Workenddate, &item.Type, &item.Master, &item.Status, &item.Company, &item.Report, &item.Report1, &item.Report2, &item.Report3, &item.Report4, &item.Report5, &item.Report6, &item.Summarytype, &item.Search, &item.User, &item.Updateuser, &item.Date)
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
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

func (p *OldaptManager) ReadRows(rows *sql.Rows) []Oldapt {
    items := make([]Oldapt, 0)

    for rows.Next() {
        var item Oldapt
        
    
        err := rows.Scan(&item.Id, &item.Aptgroup, &item.Name, &item.Workstartdate, &item.Workenddate, &item.Type, &item.Master, &item.Status, &item.Company, &item.Report, &item.Report1, &item.Report2, &item.Report3, &item.Report4, &item.Report5, &item.Report6, &item.Summarytype, &item.Search, &item.User, &item.Updateuser, &item.Date)
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

func (p *OldaptManager) Get(id int64) *Oldapt {
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

func (p *OldaptManager) GetWhere(args []any) *Oldapt {
    items := p.Find(args)
    if len(items) == 0 {
        return nil
    }

    return &items[0]
}

func (p *OldaptManager) Count(args []any) int {
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

func (p *OldaptManager) FindAll() []Oldapt {
    return p.Find(nil)
}

func (p *OldaptManager) Find(args []any) []Oldapt {
    if !p.Conn.IsConnect() {
        items := make([]Oldapt, 0)
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
        items := make([]Oldapt, 0)
        return items
    }

    defer rows.Close()

    return p.ReadRows(rows)
}





func (p *OldaptManager) GroupBy(name string, args []any) []Groupby {
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



func (p *OldaptManager) MakeMap(items []Oldapt) map[int64]Oldapt {
     ret := make(map[int64]Oldapt)
     for _, v := range items {
        ret[v.Id] = v
     }

     return ret
}

func (p *OldaptManager) FindToMap(args []any) map[int64]Oldapt {
     items := p.Find(args)
     return p.MakeMap(items)
}

func (p *OldaptManager) FindAllToMap() map[int64]Oldapt {
     items := p.Find(nil)
     return p.MakeMap(items)
}

func (p *OldaptManager) MakeNameMap(items []Oldapt) map[string]Oldapt {
     ret := make(map[string]Oldapt)
     for _, v := range items {
        ret[v.Name] = v
     }

     return ret
}

func (p *OldaptManager) FindToNameMap(args []any) map[string]Oldapt {
     items := p.Find(args)
     return p.MakeNameMap(items)
}

func (p *OldaptManager) FindAllToNameMap() map[string]Oldapt {
     items := p.Find(nil)
     return p.MakeNameMap(items)
}
