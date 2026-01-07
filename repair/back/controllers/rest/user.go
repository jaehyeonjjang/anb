package rest


import (
	"repair/controllers"
	"repair/models"

	"repair/models/user"

    "strings"
)

type UserController struct {
	controllers.Controller
}



func (c *UserController) CountByLevel(level user.Level) int {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    item := _manager.CountByLevel(level)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *UserController) FindByLevel(level user.Level) []models.User {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    item := _manager.FindByLevel(level)
    
    
    c.Set("items", item)
    
    
    return item
    
}


func (c *UserController) CountByApt(apt int64) int {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    item := _manager.CountByApt(apt)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *UserController) FindByApt(apt int64) []models.User {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    item := _manager.FindByApt(apt)
    
    
    c.Set("items", item)
    
    
    return item
    
}


func (c *UserController) CountByAptLevel(apt int64 ,level user.Level) int {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    item := _manager.CountByAptLevel(apt, level)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *UserController) FindByAptLevel(apt int64 ,level user.Level) []models.User {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    item := _manager.FindByAptLevel(apt, level)
    
    
    c.Set("items", item)
    
    
    return item
    
}


func (c *UserController) GetByEmail(email string) *models.User {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    item := _manager.GetByEmail(email)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *UserController) CountByEmail(email string) int {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    item := _manager.CountByEmail(email)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *UserController) GetByLoginid(loginid string) *models.User {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    item := _manager.GetByLoginid(loginid)
    
    c.Set("item", item)
    
    
    
    return item
    
}


func (c *UserController) CountByLoginid(loginid string) int {
    
    conn := c.NewConnection()

	_manager := models.NewUserManager(conn)
    
    item := _manager.CountByLoginid(loginid)
    
    
    
    c.Set("count", item)
    
    return item
    
}


func (c *UserController) Insert(item *models.UserUpdate) {
    
    
	conn := c.NewConnection()
    
	manager := models.NewUserManager(conn)
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

func (c *UserController) Insertbatch(item *[]models.UserUpdate) {  
    if item == nil || len(*item) == 0 {
        return
    }

    rows := len(*item)
    
    
    
	conn := c.NewConnection()
    
	manager := models.NewUserManager(conn)

    for i := 0; i < rows; i++ {
	    err := manager.Insert(&((*item)[i]))
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *UserController) Update(item *models.UserUpdate) {
    
    
	conn := c.NewConnection()

	manager := models.NewUserManager(conn)
    err := manager.Update(item)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
        return
    }
}

func (c *UserController) Delete(item *models.User) {
    
    
    conn := c.NewConnection()

	manager := models.NewUserManager(conn)

    
	err := manager.Delete(item.Id)
    if err != nil {
        c.Set("code", "error")    
        c.Set("error", err)
    }
}

func (c *UserController) Deletebatch(item *[]models.User) {
    
    
    conn := c.NewConnection()

	manager := models.NewUserManager(conn)

    for _, v := range *item {
        
    
	    err := manager.Delete(v.Id)
        if err != nil {
            c.Set("code", "error")    
            c.Set("error", err)
            return
        }
    }
}

func (c *UserController) Count() {
    
    
	conn := c.NewConnection()

	manager := models.NewUserManager(conn)

    var args []interface{}
    
    _loginid := c.Get("loginid")
    if _loginid != "" {
        args = append(args, models.Where{Column:"loginid", Value:_loginid, Compare:"="})
    }
    _passwd := c.Get("passwd")
    if _passwd != "" {
        args = append(args, models.Where{Column:"passwd", Value:_passwd, Compare:"="})
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"="})
    }
    _level := c.Geti("level")
    if _level != 0 {
        args = append(args, models.Where{Column:"level", Value:_level, Compare:"="})    
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


func (c *UserController) Read(id int64) {
    
    
	conn := c.NewConnection()

	manager := models.NewUserManager(conn)
	item := manager.Get(id)

    
    
    c.Set("item", item)
}

func (c *UserController) Index(page int, pagesize int) {
    
    
	conn := c.NewConnection()

	manager := models.NewUserManager(conn)

    var args []interface{}
    
    _loginid := c.Get("loginid")
    if _loginid != "" {
        args = append(args, models.Where{Column:"loginid", Value:_loginid, Compare:"="})
    }
    _passwd := c.Get("passwd")
    if _passwd != "" {
        args = append(args, models.Where{Column:"passwd", Value:_passwd, Compare:"="})
    }
    _name := c.Get("name")
    if _name != "" {
        args = append(args, models.Where{Column:"name", Value:_name, Compare:"like"})
        
    }
    _email := c.Get("email")
    if _email != "" {
        args = append(args, models.Where{Column:"email", Value:_email, Compare:"="})
    }
    _level := c.Geti("level")
    if _level != 0 {
        args = append(args, models.Where{Column:"level", Value:_level, Compare:"="})    
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
                    str += ", u_" + strings.Trim(v, " ")                
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





