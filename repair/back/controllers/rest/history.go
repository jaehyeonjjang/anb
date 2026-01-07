package rest


import (
	"repair/controllers"
	"repair/models"

    "strings"
)

type HistoryController struct {
	controllers.Controller
}



func (c *HistoryController) CountByApt(apt int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewHistoryManager(conn)
    
    item := _manager.CountByApt(apt)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *HistoryController) FindByApt(apt int64) []models.History {
    
    conn := c.NewConnection()

	_manager := models.NewHistoryManager(conn)
    
    item := _manager.FindByApt(apt)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Delete()
func (c *HistoryController) DeleteByApt(apt int64) {
    
    conn := c.NewConnection()

	_manager := models.NewHistoryManager(conn)
    
    err := _manager.DeleteByApt(apt)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *HistoryController) FindByCategory(category int64) []models.History {
    
    conn := c.NewConnection()

	_manager := models.NewHistoryManager(conn)
    
    item := _manager.FindByCategory(category)
    
    
    c.Set("items", item)
    
    
    return item
    
}

// @Delete()
func (c *HistoryController) DeleteByCategory(category int64) {
    
    conn := c.NewConnection()

	_manager := models.NewHistoryManager(conn)
    
    err := _manager.DeleteByCategory(category)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
    
}


func (c *HistoryController) Insert(item *models.History) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewHistoryManager(conn)
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

func (c *HistoryController) Insertbatch(item *[]models.History) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewHistoryManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *HistoryController) Update(item *models.History) {
    
    
	conn := c.NewConnection()

	manager := models.NewHistoryManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *HistoryController) Delete(item *models.History) {
    
    
    conn := c.NewConnection()

	manager := models.NewHistoryManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *HistoryController) Deletebatch(item *[]models.History) {
    
    
    conn := c.NewConnection()

	manager := models.NewHistoryManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *HistoryController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewHistoryManager(conn)

    var args []interface{}
    
    _year := c.Geti("year")
    if _year != 0 {
        args = append(args, models.Where{Column:"year", Value:_year, Compare:"="})    
    }
    _month := c.Geti("month")
    if _month != 0 {
        args = append(args, models.Where{Column:"month", Value:_month, Compare:"="})    
    }
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
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"like"})
        
    }
    _price := c.Geti64("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
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


func (c *HistoryController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewHistoryManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *HistoryController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewHistoryManager(conn)

    var args []interface{}
    
    _year := c.Geti("year")
    if _year != 0 {
        args = append(args, models.Where{Column:"year", Value:_year, Compare:"="})    
    }
    _month := c.Geti("month")
    if _month != 0 {
        args = append(args, models.Where{Column:"month", Value:_month, Compare:"="})    
    }
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
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"like"})
        
    }
    _price := c.Geti64("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
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
                    str += ", h_" + strings.Trim(v, " ")                
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





func (c *HistoryController) Sum() {
    
    
	conn := c.NewConnection()

	manager := models.NewHistoryManager(conn)

    var args []interface{}
    
    _year := c.Geti("year")
    if _year != 0 {
        args = append(args, models.Where{Column:"year", Value:_year, Compare:"="})    
    }
    _month := c.Geti("month")
    if _month != 0 {
        args = append(args, models.Where{Column:"month", Value:_month, Compare:"="})    
    }
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
    _content := c.Get("content")
    if _content != "" {
        args = append(args, models.Where{Column:"content", Value:_content, Compare:"like"})
        
    }
    _price := c.Geti64("price")
    if _price != 0 {
        args = append(args, models.Where{Column:"price", Value:_price, Compare:"="})    
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

