package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type BreakdownhistoryController struct {
	controllers.Controller
}



func (c *BreakdownhistoryController) CountByApt(apt int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownhistoryManager(conn)
    
    item := _manager.CountByApt(apt)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *BreakdownhistoryController) FindByApt(apt int64) []models.Breakdownhistory {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownhistoryManager(conn)
    
    item := _manager.FindByApt(apt)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Delete()
func (c *BreakdownhistoryController) DeleteByApt(apt int64) {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownhistoryManager(conn)
    
    err := _manager.DeleteByApt(apt)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *BreakdownhistoryController) CountByBreakdown(breakdown int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownhistoryManager(conn)
    
    item := _manager.CountByBreakdown(breakdown)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *BreakdownhistoryController) GetByBreakdown(breakdown int64) *models.Breakdownhistory {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownhistoryManager(conn)
    
    item := _manager.GetByBreakdown(breakdown)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *BreakdownhistoryController) FindByCategory(category int64) []models.Breakdownhistory {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownhistoryManager(conn)
    
    item := _manager.FindByCategory(category)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Delete()
func (c *BreakdownhistoryController) DeleteByCategory(category int64) {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownhistoryManager(conn)
    
    err := _manager.DeleteByCategory(category)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *BreakdownhistoryController) FindByMethod(method int64) []models.Breakdownhistory {
    
    conn := c.NewConnection()

	_manager := models.NewBreakdownhistoryManager(conn)
    
    item := _manager.FindByMethod(method)
    
    
    c.Set("items", item)
    
    
    return item
    
}


func (c *BreakdownhistoryController) Insert(item *models.Breakdownhistory) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewBreakdownhistoryManager(conn)
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

func (c *BreakdownhistoryController) Insertbatch(item *[]models.Breakdownhistory) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewBreakdownhistoryManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *BreakdownhistoryController) Update(item *models.Breakdownhistory) {
    
    
	conn := c.NewConnection()

	manager := models.NewBreakdownhistoryManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *BreakdownhistoryController) Delete(item *models.Breakdownhistory) {
    
    
    conn := c.NewConnection()

	manager := models.NewBreakdownhistoryManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *BreakdownhistoryController) Deletebatch(item *[]models.Breakdownhistory) {
    
    
    conn := c.NewConnection()

	manager := models.NewBreakdownhistoryManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *BreakdownhistoryController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewBreakdownhistoryManager(conn)

    var args []interface{}
    
    _topcategory := c.Geti64("topcategory")
    if _topcategory != 0 {
        args = append(args, models.Where{Column:"topcategory", Value:_topcategory, Compare:"="})    
    }
    _subcategory := c.Geti64("subcategory")
    if _subcategory != 0 {
        args = append(args, models.Where{Column:"subcategory", Value:_subcategory, Compare:"="})    
    }
    _category := c.Geti64("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _method := c.Geti64("method")
    if _method != 0 {
        args = append(args, models.Where{Column:"method", Value:_method, Compare:"="})    
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _lastdate := c.Geti("lastdate")
    if _lastdate != 0 {
        args = append(args, models.Where{Column:"lastdate", Value:_lastdate, Compare:"="})    
    }
    _duedate := c.Geti("duedate")
    if _duedate != 0 {
        args = append(args, models.Where{Column:"duedate", Value:_duedate, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _elevator := c.Geti("elevator")
    if _elevator != 0 {
        args = append(args, models.Where{Column:"elevator", Value:_elevator, Compare:"="})    
    }
    _percent := c.Geti("percent")
    if _percent != 0 {
        args = append(args, models.Where{Column:"percent", Value:_percent, Compare:"="})    
    }
    _rate := c.Geti("rate")
    if _rate != 0 {
        args = append(args, models.Where{Column:"rate", Value:_rate, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _originalcount := c.Geti("originalcount")
    if _originalcount != 0 {
        args = append(args, models.Where{Column:"originalcount", Value:_originalcount, Compare:"="})    
    }
    _originalprice := c.Geti("originalprice")
    if _originalprice != 0 {
        args = append(args, models.Where{Column:"originalprice", Value:_originalprice, Compare:"="})    
    }
    _originalduedate := c.Geti("originalduedate")
    if _originalduedate != 0 {
        args = append(args, models.Where{Column:"originalduedate", Value:_originalduedate, Compare:"="})    
    }
    _totalcount := c.Geti("totalcount")
    if _totalcount != 0 {
        args = append(args, models.Where{Column:"totalcount", Value:_totalcount, Compare:"="})    
    }
    _totalprice := c.Geti64("totalprice")
    if _totalprice != 0 {
        args = append(args, models.Where{Column:"totalprice", Value:_totalprice, Compare:"="})    
    }
    _dong := c.Geti64("dong")
    if _dong != 0 {
        args = append(args, models.Where{Column:"dong", Value:_dong, Compare:"="})    
    }
    _standard := c.Geti64("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
    }
    _breakdown := c.Geti64("breakdown")
    if _breakdown != 0 {
        args = append(args, models.Where{Column:"breakdown", Value:_breakdown, Compare:"="})    
    }
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
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


func (c *BreakdownhistoryController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewBreakdownhistoryManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *BreakdownhistoryController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewBreakdownhistoryManager(conn)

    var args []interface{}
    
    _topcategory := c.Geti64("topcategory")
    if _topcategory != 0 {
        args = append(args, models.Where{Column:"topcategory", Value:_topcategory, Compare:"="})    
    }
    _subcategory := c.Geti64("subcategory")
    if _subcategory != 0 {
        args = append(args, models.Where{Column:"subcategory", Value:_subcategory, Compare:"="})    
    }
    _category := c.Geti64("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _method := c.Geti64("method")
    if _method != 0 {
        args = append(args, models.Where{Column:"method", Value:_method, Compare:"="})    
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _lastdate := c.Geti("lastdate")
    if _lastdate != 0 {
        args = append(args, models.Where{Column:"lastdate", Value:_lastdate, Compare:"="})    
    }
    _duedate := c.Geti("duedate")
    if _duedate != 0 {
        args = append(args, models.Where{Column:"duedate", Value:_duedate, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"="})
    }
    _elevator := c.Geti("elevator")
    if _elevator != 0 {
        args = append(args, models.Where{Column:"elevator", Value:_elevator, Compare:"="})    
    }
    _percent := c.Geti("percent")
    if _percent != 0 {
        args = append(args, models.Where{Column:"percent", Value:_percent, Compare:"="})    
    }
    _rate := c.Geti("rate")
    if _rate != 0 {
        args = append(args, models.Where{Column:"rate", Value:_rate, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _originalcount := c.Geti("originalcount")
    if _originalcount != 0 {
        args = append(args, models.Where{Column:"originalcount", Value:_originalcount, Compare:"="})    
    }
    _originalprice := c.Geti("originalprice")
    if _originalprice != 0 {
        args = append(args, models.Where{Column:"originalprice", Value:_originalprice, Compare:"="})    
    }
    _originalduedate := c.Geti("originalduedate")
    if _originalduedate != 0 {
        args = append(args, models.Where{Column:"originalduedate", Value:_originalduedate, Compare:"="})    
    }
    _totalcount := c.Geti("totalcount")
    if _totalcount != 0 {
        args = append(args, models.Where{Column:"totalcount", Value:_totalcount, Compare:"="})    
    }
    _totalprice := c.Geti64("totalprice")
    if _totalprice != 0 {
        args = append(args, models.Where{Column:"totalprice", Value:_totalprice, Compare:"="})    
    }
    _dong := c.Geti64("dong")
    if _dong != 0 {
        args = append(args, models.Where{Column:"dong", Value:_dong, Compare:"="})    
    }
    _standard := c.Geti64("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
    }
    _breakdown := c.Geti64("breakdown")
    if _breakdown != 0 {
        args = append(args, models.Where{Column:"breakdown", Value:_breakdown, Compare:"="})    
    }
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
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
                    str += ", bh_" + strings.Trim(v, " ")                
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





func (c *BreakdownhistoryController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewBreakdownhistoryManager(conn)

    var args []interface{}
    
    _topcategory := c.Geti64("topcategory")
    if _topcategory != 0 {
        args = append(args, models.Where{Column:"topcategory", Value:_topcategory, Compare:"="})    
    }
    _subcategory := c.Geti64("subcategory")
    if _subcategory != 0 {
        args = append(args, models.Where{Column:"subcategory", Value:_subcategory, Compare:"="})    
    }
    _category := c.Geti64("category")
    if _category != 0 {
        args = append(args, models.Where{Column:"category", Value:_category, Compare:"="})    
    }
    _method := c.Geti64("method")
    if _method != 0 {
        args = append(args, models.Where{Column:"method", Value:_method, Compare:"="})    
    }
    _count := c.Geti("count")
    if _count != 0 {
        args = append(args, models.Where{Column:"count", Value:_count, Compare:"="})    
    }
    _lastdate := c.Geti("lastdate")
    if _lastdate != 0 {
        args = append(args, models.Where{Column:"lastdate", Value:_lastdate, Compare:"="})    
    }
    _duedate := c.Geti("duedate")
    if _duedate != 0 {
        args = append(args, models.Where{Column:"duedate", Value:_duedate, Compare:"="})    
    }
    _remark := c.Get("remark")
    if _remark != "" {
        args = append(args, models.Where{Column:"remark", Value:_remark, Compare:"like"})
    }
    _elevator := c.Geti("elevator")
    if _elevator != 0 {
        args = append(args, models.Where{Column:"elevator", Value:_elevator, Compare:"="})    
    }
    _percent := c.Geti("percent")
    if _percent != 0 {
        args = append(args, models.Where{Column:"percent", Value:_percent, Compare:"="})    
    }
    _rate := c.Geti("rate")
    if _rate != 0 {
        args = append(args, models.Where{Column:"rate", Value:_rate, Compare:"="})    
    }
    _type := c.Geti("type")
    if _type != 0 {
        args = append(args, models.Where{Column:"type", Value:_type, Compare:"="})    
    }
    _originalcount := c.Geti("originalcount")
    if _originalcount != 0 {
        args = append(args, models.Where{Column:"originalcount", Value:_originalcount, Compare:"="})    
    }
    _originalprice := c.Geti("originalprice")
    if _originalprice != 0 {
        args = append(args, models.Where{Column:"originalprice", Value:_originalprice, Compare:"="})    
    }
    _originalduedate := c.Geti("originalduedate")
    if _originalduedate != 0 {
        args = append(args, models.Where{Column:"originalduedate", Value:_originalduedate, Compare:"="})    
    }
    _totalcount := c.Geti("totalcount")
    if _totalcount != 0 {
        args = append(args, models.Where{Column:"totalcount", Value:_totalcount, Compare:"="})    
    }
    _totalprice := c.Geti64("totalprice")
    if _totalprice != 0 {
        args = append(args, models.Where{Column:"totalprice", Value:_totalprice, Compare:"="})    
    }
    _dong := c.Geti64("dong")
    if _dong != 0 {
        args = append(args, models.Where{Column:"dong", Value:_dong, Compare:"="})    
    }
    _standard := c.Geti64("standard")
    if _standard != 0 {
        args = append(args, models.Where{Column:"standard", Value:_standard, Compare:"="})    
    }
    _breakdown := c.Geti64("breakdown")
    if _breakdown != 0 {
        args = append(args, models.Where{Column:"breakdown", Value:_breakdown, Compare:"="})    
    }
    _apt := c.Geti64("apt")
    if _apt != 0 {
        args = append(args, models.Where{Column:"apt", Value:_apt, Compare:"="})    
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
    

    
    
    item := manager.Sum(args)
	c.Set("item", item)
}

