package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type OldaptController struct {
	controllers.Controller
}



func (c *OldaptController) Insert(item *models.Oldapt) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewOldaptManager(conn)
	err := manager.Insert(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }

    id := manager.GetIdentity()
    c.Result["id"] = id
    item.Id = id
}

func (c *OldaptController) Insertbatch(item *[]models.Oldapt) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewOldaptManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *OldaptController) Update(item *models.Oldapt) {
    
    
	conn := c.NewConnection()

	manager := models.NewOldaptManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *OldaptController) Delete(item *models.Oldapt) {
    
    
    conn := c.NewConnection()

	manager := models.NewOldaptManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *OldaptController) Deletebatch(item *[]models.Oldapt) {
    
    
    conn := c.NewConnection()

	manager := models.NewOldaptManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *OldaptController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewOldaptManager(conn)

    var args []interface{}
    
    _aptgroup := c.Geti64("aptgroup")
    if _aptgroup != 0 {
        args = append(args, models.Where{Column:"aptgroup", Value:_aptgroup, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _workstartdate := c.Get("workstartdate")
    if _workstartdate != "" {
        args = append(args, models.Where{Column:"workstartdate", Value:_workstartdate, Compare:"="})
    }
    _workenddate := c.Get("workenddate")
    if _workenddate != "" {
        args = append(args, models.Where{Column:"workenddate", Value:_workenddate, Compare:"="})
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _master := c.Geti64("master")
    if _master != 0 {
        args = append(args, models.Where{Column:"master", Value:_master, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _report := c.Geti("report")
    if _report != 0 {
        args = append(args, models.Where{Column:"report", Value:_report, Compare:"="})    
    }
    _report1 := c.Geti("report1")
    if _report1 != 0 {
        args = append(args, models.Where{Column:"report1", Value:_report1, Compare:"="})    
    }
    _report2 := c.Geti("report2")
    if _report2 != 0 {
        args = append(args, models.Where{Column:"report2", Value:_report2, Compare:"="})    
    }
    _report3 := c.Geti("report3")
    if _report3 != 0 {
        args = append(args, models.Where{Column:"report3", Value:_report3, Compare:"="})    
    }
    _report4 := c.Geti("report4")
    if _report4 != 0 {
        args = append(args, models.Where{Column:"report4", Value:_report4, Compare:"="})    
    }
    _report5 := c.Geti("report5")
    if _report5 != 0 {
        args = append(args, models.Where{Column:"report5", Value:_report5, Compare:"="})    
    }
    _report6 := c.Geti("report6")
    if _report6 != 0 {
        args = append(args, models.Where{Column:"report6", Value:_report6, Compare:"="})    
    }
    _summarytype := c.Geti("summarytype")
    if _summarytype != 0 {
        args = append(args, models.Where{Column:"summarytype", Value:_summarytype, Compare:"="})    
    }
    _search := c.Get("search")
    if _search != "" {
        args = append(args, models.Where{Column:"search", Value:_search, Compare:"like"})
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _updateuser := c.Geti64("updateuser")
    if _updateuser != 0 {
        args = append(args, models.Where{Column:"updateuser", Value:_updateuser, Compare:"="})    
    }
    _startdate := c.Get("startdate")
    _enddate := c.Get("enddate")

    if _startdate != "" && _enddate != "" {        
        var v [2]string
        v[0] = _startdate
        v[1] = _enddate  
        args = append(args, models.Where{Column:"date", Value:v, Compare:"between"})    
    } else if  _startdate != "" {          
        args = append(args, models.Where{Column:"date", Value:_startdate, Compare:">="})
    } else if  _enddate != "" {          
        args = append(args, models.Where{Column:"date", Value:_enddate, Compare:"<="})            
    }
    

    
    
    total := manager.Count(args)
	c.Set("total", total)
}


func (c *OldaptController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewOldaptManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *OldaptController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewOldaptManager(conn)

    var args []interface{}
    
    _aptgroup := c.Geti64("aptgroup")
    if _aptgroup != 0 {
        args = append(args, models.Where{Column:"aptgroup", Value:_aptgroup, Compare:"="})    
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _workstartdate := c.Get("workstartdate")
    if _workstartdate != "" {
        args = append(args, models.Where{Column:"workstartdate", Value:_workstartdate, Compare:"="})
    }
    _workenddate := c.Get("workenddate")
    if _workenddate != "" {
        args = append(args, models.Where{Column:"workenddate", Value:_workenddate, Compare:"="})
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _master := c.Geti64("master")
    if _master != 0 {
        args = append(args, models.Where{Column:"master", Value:_master, Compare:"="})    
    }
    _status := c.Geti("status")
    if _status != 0 {
        args = append(args, models.Where{Column:"status", Value:_status, Compare:"="})    
    }
    _company := c.Geti64("company")
    if _company != 0 {
        args = append(args, models.Where{Column:"company", Value:_company, Compare:"="})    
    }
    _report := c.Geti("report")
    if _report != 0 {
        args = append(args, models.Where{Column:"report", Value:_report, Compare:"="})    
    }
    _report1 := c.Geti("report1")
    if _report1 != 0 {
        args = append(args, models.Where{Column:"report1", Value:_report1, Compare:"="})    
    }
    _report2 := c.Geti("report2")
    if _report2 != 0 {
        args = append(args, models.Where{Column:"report2", Value:_report2, Compare:"="})    
    }
    _report3 := c.Geti("report3")
    if _report3 != 0 {
        args = append(args, models.Where{Column:"report3", Value:_report3, Compare:"="})    
    }
    _report4 := c.Geti("report4")
    if _report4 != 0 {
        args = append(args, models.Where{Column:"report4", Value:_report4, Compare:"="})    
    }
    _report5 := c.Geti("report5")
    if _report5 != 0 {
        args = append(args, models.Where{Column:"report5", Value:_report5, Compare:"="})    
    }
    _report6 := c.Geti("report6")
    if _report6 != 0 {
        args = append(args, models.Where{Column:"report6", Value:_report6, Compare:"="})    
    }
    _summarytype := c.Geti("summarytype")
    if _summarytype != 0 {
        args = append(args, models.Where{Column:"summarytype", Value:_summarytype, Compare:"="})    
    }
    _search := c.Get("search")
    if _search != "" {
        args = append(args, models.Where{Column:"search", Value:_search, Compare:"like"})
    }
    _user := c.Geti64("user")
    if _user != 0 {
        args = append(args, models.Where{Column:"user", Value:_user, Compare:"="})    
    }
    _updateuser := c.Geti64("updateuser")
    if _updateuser != 0 {
        args = append(args, models.Where{Column:"updateuser", Value:_updateuser, Compare:"="})    
    }
    _startdate := c.Get("startdate")
    _enddate := c.Get("enddate")
    if _startdate != "" && _enddate != "" {        
        var v [2]string
        v[0] = _startdate
        v[1] = _enddate  
        args = append(args, models.Where{Column:"date", Value:v, Compare:"between"})    
    } else if  _startdate != "" {          
        args = append(args, models.Where{Column:"date", Value:_startdate, Compare:">="})
    } else if  _enddate != "" {          
        args = append(args, models.Where{Column:"date", Value:_enddate, Compare:"<="})            
    }
    

    
    
    if page != 0 && pagesize != 0 {
        args = append(args, models.Paging(page, pagesize))
    }
    
    orderby := c.Get("orderby")
    if orderby == "" {
        if page != 0 && pagesize != 0 {
            orderby = "id desc"
            args = append(args, models.Ordering(orderby))
        }
    } else {
        orderbys := strings.Split(orderby, ",")

        str := ""
        for i, v := range orderbys {
            if i == 0 {
                str += v
            } else {
                if strings.Contains(v, "_") {                   
                    str += ", " + strings.Trim(v, " ")
                } else {
                    str += ", a_" + strings.Trim(v, " ")                
                }
            }
        }
        
        args = append(args, models.Ordering(str))
    }
    
	items := manager.Find(args)
	c.Set("items", items)

    if page == 1 {
       total := manager.Count(args)
	   c.Set("total", total)
    }
}





