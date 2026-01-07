#!/opt/homebrew/bin/php
<?php

function db_connect()
{
    global $GLOBAL_VAR;
    global $_db;

    if (!isset($_db)) {
    $_db = new PDO('mysql:host='.$GLOBAL_VAR['db_host'].';dbname='.$GLOBAL_VAR['db_name'], $GLOBAL_VAR['db_user'], $GLOBAL_VAR['db_pass']);

        $_db->exec('SET NAMES \'UTF8\'');
        //$_db->exec('SET SESSION TIME_ZONE = \'' . date_default_timezone_get() . '\'');
    }

    return $_db;
}

function getGoType($name) {
    if ($name == 'int')
        return 'int';

    if ($name == 'bigint')
        return 'int64';

    if ($name == 'decimal')
        return 'int';

    if ($name == 'float')
        return 'float32';

    if ($name == 'double')
        return 'float64';

    if ($name == 'varchar')
        return 'string';

    if ($name == 'text')
        return 'string';

    if ($name == 'datetime')
        return 'string';

    if ($name == 'date')
        return 'string';

    if ($name == 'time')
        return 'string';
    if ($name == 'json')
        return 'string';
    return $name;
}

function getField($name) {
    return str_replace('!', '', str_replace('#', '', str_replace('<', '', str_replace('>', '', str_replace('@', '', str_replace('^','', $name))))));
}

function getField2($name) {
    if ($name == 'type')
        $name = 'typeid';

    return str_replace('!', '', str_replace('#', '', str_replace('<', '', str_replace('>', '', str_replace('@', '', str_replace('^','', $name))))));
}

function process($table_name) {
    global $GLOBAL_VAR;
    global $_db;

    global $search_array;
    global $alias_array;
    global $increase_array;
    global $delete_array;
    global $update_array;

    global $sum_array;
    global $sum2_array;

    global $cache_array;

    $ret = '';
    $ret_manager = '';

    $viewstatus = false;

    $query = "select column_name as column_name, data_type as data_type from information_schema.columns where table_schema = '".$GLOBAL_VAR['db_name']."' and table_name = '$table_name' order by ORDINAL_POSITION";
    $result = $_db->query($query);

    $table = str_replace("_tb", '', $table_name);

    if (substr($table, -3) == '_vw')
        $table_type = 'view';
    else
        $table_type = 'table';

    $table = str_replace("_vw", '', $table);

    $table[0] = strtoupper($table[0]);

    $struct_column = array();

    foreach($result as $row)
    {
        //$row['column_name'] = $row['COLUMN_NAME'];
        //$row['data_type'] = $row['DATA_TYPE'];

        $column_name = $row['column_name'];

        $temp = explode("_", $column_name);

        if (count($temp) == 2)
        {
            $columns[] = $temp[1];
            $types[] = $row['data_type'];
            $prefix = $temp[0];

            if ($temp[1] == 'viewstatus')
                $viewstatus = true;

            $type_array[$temp[1]] = $row['data_type'];
        }
        else
        {
            return null;
            //$columns[] = $temp[1];
        }

        if ($temp[1] != 'id')
        {
            $insert_columns[] = $temp[1];
            $insert_columns2[] = $row['column_name'];
            $insert_type[] = $row['data_type'];
        }

        $struct_column[] = $row['data_type'];
    }

    if (!isset($columns))
        return null;

    if (!isset($insert_columns))
        return null;

    $ret = "type $table struct {\n";

    for($i = 0; $i < count($columns); $i++) {
        $column = ucfirst($columns[$i]);
        $type = getGoType($struct_column[$i]);

        $ret .= sprintf("    %-25s%s `json:\"%s, %s\" form:\"%s\"`\n", $column, $type, ucfirst($columns[$i]), $type, $columns[$i]);
    }

    $ret .= sprintf("    %-25s%s `form:\"%s\"`\n", 'Extra', 'interface{}', 'extra');
    $ret .= "}\n";

    $insert_column = implode(", ", $insert_columns2);

    for($i = 0; $i < count($insert_columns); $i++)
    {
        $insert[] = ":".$insert_columns[$i];
    }

    $insert_value = implode(', ', $insert);

    $date = 'date';
    $newtable = $table;

    for($i = 0; $i < count($columns); $i++)
    {
        //$bindall[] = "${prefix}_".$columns[$i] . " as " . $columns[$i] . "";
        $bindall[] = "${prefix}_".$columns[$i] . "";
    }

    $bindall_str = implode(', ', $bindall);

    $insert_value = implode(', ', array_fill(0, count($insert_columns), '?'));

    $ret_manager .= "
type ${table}Manager struct {
    Conn    *sql.DB
    Result  *sql.Result
    Prefix  string
    Index   string
}";

    if (in_array(strtolower($table), $cache_array)) {
        $ret_manager .= "

var _cache${table}s map[int64]*${table}

func init() {
    _cache${table}s = make(map[int64]*${table})
}";

    }

    $ret_manager .= "

func New${table}Manager(conn *sql.DB) *${table}Manager {
    var item ${table}Manager

    if conn == nil {
        item.Conn = NewConnection()
    } else {
        item.Conn = conn
    }

    item.Prefix = \"$prefix\"
    item.Index = \"\"

    return &item
}

func (p *${table}Manager) Close() {
    if p.Conn != nil {
        p.Conn.Close()
    }
}

func (p *${table}Manager) GetLast(items *[]${table}) *${table} {
    if items == nil {
        return nil
    } else if len(*items) == 0 {
        return nil
    } else {
        return &(*items)[0]
    }
}

func (p *${table}Manager) SetIndex(index string) {
    p.Index = index
}

func (p *${table}Manager) GetQuery() string {
    ret := \"\"

    tableName := \"$table_name\"
    if config.Database == \"mssql\" || config.Database == \"sqlserver\" {
        tableName = config.Owner + \".$table_name\"
    }

    str := \"select $bindall_str from \" + tableName + \" \"

    if p.Index == \"\" {
        ret = str
    } else {
        ret = str + \" use index(\" + p.Index + \") \"
    }

    return ret;
}

func (p *${table}Manager) GetQuerySelect() string {
    ret := \"\"

    tableName := \"$table_name\"
    if config.Database == \"mssql\" || config.Database == \"sqlserver\" {
        tableName = config.Owner + \".$table_name\"
    }

    str := \"select count(*) from \" + tableName + \" \"

    if p.Index == \"\" {
        ret = str
    } else {
        ret = str + \" use index(\" + p.Index + \") \"
    }

    return ret;
}

func (p *${table}Manager) Insert(item *${table}) error {
    if p.Conn == nil {
        return errors.New(\"Connection Error\")
    }";

    for($l = 0; $l < count($columns); $l++) {
        $column = ucfirst($columns[$l]);

        if ($column == 'Date') {
        continue;
        }

        if ($struct_column[$l] == 'datetime') {
        $ret_manager .= "

    if item.${column} == \"\" {
        item.${column} = \"1000-01-01 00:00:00\"
    }";
        }
    }


    $ret_manager .= "

    if item.Date == \"\" {
        t := time.Now()
        item.Date = fmt.Sprintf(\"%04d-%02d-%02d %02d:%02d:%02d\", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
    }

    tableName := \"$table_name\"
    if config.Database == \"mssql\" || config.Database == \"sqlserver\" {
        tableName = config.Owner + \".$table_name\"
    }

    var err error
    var res sql.Result
    query := \"\"

    if item.Id > 0 {
        query = \"insert into \" + tableName + \" (".$prefix."_id, $insert_column) values (?, $insert_value)\"
        res, err = p.Conn.Exec(query, item.Id, ";

    for($i = 0; $i < count($insert_columns); $i++) {
        if ($i > 0)
            $ret_manager .= ', ';
        $ret_manager .= 'item.' . ucfirst($insert_columns[$i]);
    }

    $ret_manager .= ")
    } else {
        query = \"insert into \" + tableName + \" ($insert_column) values ($insert_value)\"
        res, err = p.Conn.Exec(query, ";

    for($i = 0; $i < count($insert_columns); $i++) {
        if ($i > 0)
            $ret_manager .= ', ';
        $ret_manager .= 'item.' . ucfirst($insert_columns[$i]);
    }

    $ret_manager .= ")
    }

    if err == nil {
        p.Result = &res
";

    if (in_array(strtolower($table), $cache_array)) {

        $ret_manager .= "
        id := p.GetIdentity()

        item.Id = id
        _cache${table}s[id] = item
";
    }

    $ret_manager .= "    } else {
        log.Println(item)
        log.Println(err)
        p.Result = nil
    }

    return err
}";

    // delete start -----------------------------------------------------------------------------------------------------------------

    $ret_manager .= "
func (p *${table}Manager) Delete(id int64) error {
    if p.Conn == nil {
        return errors.New(\"Connection Error\")
    }

    tableName := \"$table_name\"
    if config.Database == \"mssql\" || config.Database == \"sqlserver\" {
        tableName = config.Owner + \".$table_name\"
    }
    query := \"delete from \" + tableName + \" where ${prefix}_id = ?\"
    _, err := p.Conn.Exec(query, id)
";

    if (in_array(strtolower($table), $cache_array)) {
        $ret_manager .= "
    delete(_cache${table}s, id)
";

    }

    $ret_manager .= "
    return err
}";

    if (isset($delete_array[strtolower($table)]))
    {
        for($i = 0; $i < count($delete_array[strtolower($table)]); $i++)
        {
            $field = $delete_array[strtolower($table)][$i];

            if (!is_array($field))
                $field = array($field);

            $fieldname = '';
            for($j = 0; $j < count($field); $j++)
            {
                $fieldname .= ucfirst($field[$j]);
            }

            $ret_manager .= "

func (p *${table}Manager) DeleteBy$fieldname(";

            for($j = 0; $j < count($field); $j++)
            {
                if ($j > 0)
                    $ret_manager .= ", ";

                $ret_manager .= $field[$j] . " " . GetGoType($type_array[$field[$j]]);
            }

            $ret_manager .= ") error {
    if p.Conn == nil {
        return errors.New(\"Connection Error\")
    }

    tableName := \"$table_name\"
    if config.Database == \"mssql\" || config.Database == \"sqlserver\" {
        tableName = config.Owner + \".$table_name\"
    }
    var params []interface{}
    query := \"delete from \" + tableName + \" where ";

            for ($j = 0; $j < count($field); $j++) {
                if ($j > 0)
                    $ret_manager .= " and ";

                $ret_manager .= "${prefix}_".$field[$j] . " = ?";
            }

            $ret_manager .= "\"\n";

            for ($j = 0; $j < count($field); $j++) {
                $ret_manager .= "\tparams = append(params, " . $field[$j] . ")\n";
            }

            $ret_manager .= "
    err := ExecArray(p.Conn, query, params)

    return err
}
";

        }
    }

    // delete end -----------------------------------------------------------------------------------------------------------------


    $ret_manager .= "
func (p *${table}Manager) Update(item *${table}) error {
    if p.Conn == nil {
        return errors.New(\"Connection Error\")
    }
";

    for($l = 0; $l < count($columns); $l++) {
        $column = ucfirst($columns[$l]);

        if ($column == 'Date') {
        continue;
        }

        if ($struct_column[$l] == 'datetime') {
        $ret_manager .= "
    if item.${column} == \"\" {
        item.${column} = \"1000-01-01 00:00:00\"
    }
";
        }
    }

    $update = array();

    for ($i = 0; $i < count($insert_columns); $i++) {
        $update[] = $prefix . '_' . $insert_columns[$i] . ' = ?';
    }

    $ret_manager .= "
    tableName := \"$table_name\"
    if config.Database == \"mssql\" || config.Database == \"sqlserver\" {
        tableName = config.Owner + \".$table_name\"
    }
";
    $ret_manager .= "\n\tquery := \"update \" + tableName + \" set " . implode(',', $update) .  " where ${prefix}_id = ?\"\n";
    $ret_manager .= "\t_, err := p.Conn.Exec(query, ";

    for($i = 0; $i < count($insert_columns); $i++) {
        if ($i > 0)
            $ret_manager .= ', ';
        $ret_manager .= 'item.' . ucfirst($insert_columns[$i]);
    }

    $ret_manager .= ", item.Id)
";

    if (in_array(strtolower($table), $cache_array)) {

        $ret_manager .= "
    _cache${table}s[item.Id] = item
";
    }

    $ret_manager .= "
    return err
}
";
    if (isset($update_array[strtolower($table)]))
    {
        for($i = 0; $i < count($update_array[strtolower($table)]); $i++)
        {
            $field = $update_array[strtolower($table)][$i]['where'];

            if (!is_array($field))
                $field = array($field);

            $fieldname = '';
            for($j = 0; $j < count($field); $j++)
            {
                $fieldname .= ucfirst($field[$j]);
            }


            $field2 = $update_array[strtolower($table)][$i]['field'];

            if (!is_array($field2))
                $field2 = array($field2);

            $fieldname2 = '';
            for($j = 0; $j < count($field2); $j++)
            {
                $fieldname2 .= ucfirst($field2[$j]);
            }


            if ($fieldname2 == $fieldname) {
            $ret_manager .= "

func (p *${table}Manager) Update${fieldname2}By${fieldname}(";


            for($j = 0; $j < count($field2); $j++)
            {
                if (getField2($field2[$j]) == 'default')
                    $ret_manager .= "defaultvalue ".getGoType($type_array[getField($field2[$j])]);
                else
                    $ret_manager .= getField2($field2[$j])." ".getGoType($type_array[getField($field2[$j])]);
                $ret_manager .= ", ";
            }

            for($j = 0; $j < count($field); $j++)
            {
                if ($j > 0)
                    $ret_manager .= ", ";

                if (getField2($field[$j]) == 'default')
                    $ret_manager .= "defaultvalue ".getGoType($type_array[getField($field[$j])]);
                else
                    $ret_manager .= getField2($field[$j])."2 ".getGoType($type_array[getField($field[$j])]);
            }

            } else {
            $ret_manager .= "

func (p *${table}Manager) Update${fieldname2}By$fieldname(";

            for($j = 0; $j < count($field2); $j++)
            {
                if (getField2($field2[$j]) == 'default')
                    $ret_manager .= "defaultvalue ".getGoType($type_array[getField($field2[$j])]);
                else
                    $ret_manager .= getField2($field2[$j])." ".getGoType($type_array[getField($field2[$j])]);
                $ret_manager .= ", ";
            }

            for($j = 0; $j < count($field); $j++)
            {
                if ($j > 0)
                    $ret_manager .= ", ";

                if (getField2($field[$j]) == 'default')
                    $ret_manager .= "defaultvalue ".getGoType($type_array[getField($field[$j])]);
                else
                    $ret_manager .= getField2($field[$j])." ".getGoType($type_array[getField($field[$j])]);
            }
            }

            $ret_manager .= ") error {
    if p.Conn == nil {
        return errors.New(\"Connection Error\")
    }
";

    $update = array();

    $field = $update_array[strtolower($table)][$i]['field'];
    for ($j = 0; $j < count($field); $j++) {
        $update[] = $prefix . '_' . $field[$j] . ' = ?';
    }

    $ret_manager .= "
    tableName := \"$table_name\"
    if config.Database == \"mssql\" || config.Database == \"sqlserver\" {
        tableName = config.Owner + \".$table_name\"
    }
";


    if ($fieldname2 == $fieldname) {
        $ret_manager .= "\n\tquery := \"update \" + tableName + \" set ${prefix}_" . strtolower($fieldname2) . " = ? where ${prefix}_" . strtolower($fieldname) . " = ?\"\n";
    } else {
        $ret_manager .= "\n\tquery := \"update \" + tableName + \" set " . implode(',', $update) .  " where ";
        $field4 = $update_array[strtolower($table)][$i]['where'];

        if (!is_array($field4))
            $field4 = array($field4);

        $fieldname = '';
        for($j = 0; $j < count($field4); $j++) {
            if ($j > 0)
                $ret_manager .= ' and ';

            $ret_manager .= "${prefix}_" . strtolower($field4[$j]) . " = ?";
        }


        $ret_manager .= "\"\n";
    }

    if ($fieldname2 == $fieldname) {
        $ret_manager .= "\t_, err := p.Conn.Exec(query, " . strtolower($fieldname2) . ", " . strtolower($fieldname) . "2)\n";
    } else {
    $ret_manager .= "\t_, err := p.Conn.Exec(query, ";

    for ($j = 0; $j < count($field); $j++) {
        if ($j > 0)
            $ret_manager .= ', ';

        if ($field[$j] == 'default')
            $ret_manager .= 'defaultvalue';
        else
            $ret_manager .= $field[$j];
    }

    $ret_manager .= ", ";

    $field = $update_array[strtolower($table)][$i]['where'];

    if (!is_array($field))
        $field = array($field);

    $fieldname = '';
    for($j = 0; $j < count($field); $j++) {
        if ($j > 0)
            $fieldname .= ', ';

        $fieldname .= $field[$j];
    }

$ret_manager .= $fieldname . ")
";
    }

    if (in_array(strtolower($table), $cache_array)) {

        $ret_manager .= "
    _cache${table}s[item.Id] = item
";
    }

    $ret_manager .= "
    return err
}
";

        }
    }

    if (isset($increase_array[strtolower($table)]))
    {
        for($i = 0; $i < count($increase_array[strtolower($table)]); $i++)
        {
            $field = $increase_array[strtolower($table)][$i]['where'];

            if (!is_array($field))
                $field = array($field);

            $fieldname = '';
            for($j = 0; $j < count($field); $j++)
            {
                $fieldname .= ucfirst($field[$j]);
            }


            $field2 = $increase_array[strtolower($table)][$i]['field'];

            if (!is_array($field2))
                $field2 = array($field2);

            $fieldname2 = '';
            for($j = 0; $j < count($field2); $j++)
            {
                $fieldname2 .= ucfirst($field2[$j]);
            }

            $ret_manager .= "

func (p *${table}Manager) Increase${fieldname2}By$fieldname(";

            for($j = 0; $j < count($field2); $j++)
            {
                 $ret_manager .= getField2($field2[$j])." ".getGoType($type_array[getField($field2[$j])]);
                 $ret_manager .= ", ";
            }

            for($j = 0; $j < count($field); $j++)
            {
                if ($j > 0)
                    $ret_manager .= ", ";

                $ret_manager .= getField2($field[$j])." ".getGoType($type_array[getField($field[$j])]);
            }

            $ret_manager .= ") error {
    if p.Conn == nil {
        return errors.New(\"Connection Error\")
    }
";

    $increase = array();

    $field = $increase_array[strtolower($table)][$i]['field'];
    for ($j = 0; $j < count($field); $j++) {
        $increase[] = $prefix . '_' . $field[$j] . ' = ' . $prefix . '_' . $field[$j] . ' +  ?';
    }

    $ret_manager .= "
    tableName := \"$table_name\"
    if config.Database == \"mssql\" || config.Database == \"sqlserver\" {
        tableName = config.Owner + \".$table_name\"
    }
";
    $ret_manager .= "\n\tquery := \"update \" + tableName + \" set " . implode(',', $increase) .  " where ${prefix}_" . strtolower($fieldname) . " = ?\"\n";
    $ret_manager .= "\t_, err := p.Conn.Exec(query, ";

    for ($j = 0; $j < count($field); $j++) {
        if ($j > 0)
            $ret_manager .= ', ';
        $ret_manager .= $field[$j];
    }

    $ret_manager .= ", ";




            $field = $increase_array[strtolower($table)][$i]['where'];

            if (!is_array($field))
                $field = array($field);

            $fieldname = '';
            for($j = 0; $j < count($field); $j++)
            {
                $fieldname .= $field[$j];
            }

$ret_manager .= $fieldname . ")
";

    if (in_array(strtolower($table), $cache_array)) {

        $ret_manager .= "
    _cache${table}s[item.Id] = item
";
    }

    $ret_manager .= "
    return err
}
";

        }
    }


    $ret_manager .= "
func (p *${table}Manager) GetIdentity() int64 {
    if p.Result == nil {
        return 0
    }

    id, err := (*p.Result).LastInsertId()

    if err != nil {
        return 0
    } else {
        return id
    }
}
    ";


    // list start -----------------------------------------------------------------------------------------------------------------

    $ret_manager .= "

func (p *${table}Manager) ReadRow(rows *sql.Rows) *${table} {
    var item ${table}
    var err error

    if rows.Next() {
        err = rows.Scan(";

    for ($i = 0; $i < count($columns); $i++) {
        if ($i > 0)
            $ret_manager .= ", ";
        $ret_manager .= "&item." . ucfirst($columns[$i]);
    }

    $ret_manager .= ")
";

        for($l = 0; $l < count($columns); $l++) {
        $column = ucfirst($columns[$l]);

        if ($column == 'Date') {
        continue;
        }

        if ($struct_column[$l] == 'datetime') {
        $ret_manager .= "
        if item.${column} == \"1000-01-01 00:00:00\" {
            item.${column} = \"\"
        }
";
        }
    }

    $ret_manager .= "    } else {
        return nil
    }

    if err != nil {
        return nil
    } else {
        return &item
    }
}

func (p *${table}Manager) ReadRows(rows *sql.Rows) *[]${table} {
    var items []${table}
    var err error

    for rows.Next() {
        var item ${table}
        err = rows.Scan(";

    for ($i = 0; $i < count($columns); $i++) {
        if ($i > 0)
            $ret_manager .= ", ";
        $ret_manager .= "&item." . ucfirst($columns[$i]);
    }

    $ret_manager .= ")
";

        for($l = 0; $l < count($columns); $l++) {
        $column = ucfirst($columns[$l]);

        if ($column == 'Date') {
        continue;
        }

        if ($struct_column[$l] == 'datetime') {
        $ret_manager .= "
        if item.${column} == \"1000-01-01 00:00:00\" {
            item.${column} = \"\"
        }
";
        }
    }

    $ret_manager .= "
        items = append(items, item)
    }

    if err != nil {
        return nil
    } else {
        return &items
    }
}

func (p *${table}Manager) Get(id int64) *${table} {
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + \" where ${prefix}_id = ?\"

    rows, err := p.Conn.Query(query, id)

    if err != nil {
        log.Printf(\"query error : %v, %v\\n\", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}
";

$ret_manager .= "
func (p *${table}Manager) GetList(page int, pagesize int, order string) *[]${table} {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    if page > 0 && pagesize > 0 {
        if order == \"\" {
            order = \"".$prefix."_id desc\"
        } else {
            order = \"".$prefix."_\" + order
        }
        query += \" order by \" + order
        if config.Database == \"mysql\" {
            query += \" limit ? offset ?\"
            rows, err = p.Conn.Query(query, pagesize, startpage)
        } else if config.Database == \"mssql\" || config.Database == \"sqlserver\" {
            query += \"OFFSET ? ROWS FETCH NEXT ? ROWS ONLY\"
            rows, err = p.Conn.Query(query, startpage, pagesize)
        }
    } else {
        if order == \"\" {
            order = \"".$prefix."_id\"
        } else {
            order = \"".$prefix."_\" + order
        }
        query += \" order by \" + order
        rows, err = p.Conn.Query(query)
    }

    if err != nil {
        log.Printf(\"query error : %v, %v\\n\", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRows(rows)
}
";

    $ret_manager .= "

func (p *${table}Manager) GetCount() int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    rows, err := p.Conn.Query(query)

    if err != nil {
        log.Printf(\"query error : %v, %v\\n\", err, query)
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

func (p *${table}Manager) GetListInID(ids []int, page int, pagesize int, order string) *[]${table} {
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery()

    var rows *sql.Rows
    var err error

    query = query + \" where ".$prefix."_id in (\" + strings.Trim(strings.Replace(fmt.Sprint(ids), \" \", \", \", -1), \"[]\") + \")\"

    if page > 0 && pagesize > 0 {
        if order == \"\" {
            order = \"".$prefix."_id desc\"
        } else {
            order = \"".$prefix."_\" + order
        }
        query += \" order by \" + order
        if config.Database == \"mysql\" {
            query += \" limit ? offset ?\"
            rows, err = p.Conn.Query(query, pagesize, startpage)
        } else if config.Database == \"mssql\" || config.Database == \"sqlserver\" {
            query += \"OFFSET ? ROWS FETCH NEXT ? ROWS ONLY\"
            rows, err = p.Conn.Query(query, startpage, pagesize)
        }
    } else {
        if order == \"\" {
            order = \"".$prefix."_id\"
        } else {
            order = \"".$prefix."_\" + order
        }
        query += \" order by \" + order
        rows, err = p.Conn.Query(query)
    }

    if err != nil {
        log.Printf(\"query error : %v, %v\\n\", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRows(rows)
}
";

    $ret_manager .= "

func (p *${table}Manager) GetCountInID(ids []int) int {
    if p.Conn == nil {
        return 0
    }

    query := p.GetQuerySelect()

    query = query + \" where ".$prefix."_id in (\" + strings.Trim(strings.Replace(fmt.Sprint(ids), \" \", \", \", -1), \"[]\") + \")\"

    rows, err := p.Conn.Query(query)

    if err != nil {
        log.Printf(\"query error : %v, %v\\n\", err, query)
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
";

    if (isset($search_array[strtolower($table)])) {
        for($i = 0; $i < count($search_array[strtolower($table)]); $i++) {
            $field = $search_array[strtolower($table)][$i];

            if (!is_array($field))
                $field = array($field);

            $fieldname = '';
            for($j = 0; $j < count($field); $j++)
            {
                if (substr($field[$j], 0, 1) == '^')
                    $fieldname .= 'Not';

                $fieldname .= ucfirst(getField($field[$j]));
            }


            if (substr($field[0], 0, 1) == '!') {
                $ret_manager .= "
func (p *${table}Manager) GetBy$fieldname(";

                for($j = 0; $j < count($field); $j++) {
                    if ($j > 0)
                        $ret_manager .= ", ";
                    if (substr($field[$j], 0, 2) == '<>')
                        $ret_manager .= getField($field[$j])."1 string, ". getField($field[$j])."2 string";
                    else if (getField2($field[$j]) == 'default')
                        $ret_manager .= "defaultvalue ".getGoType($type_array[getField($field[$j])]);
                    else
                        $ret_manager .= getField2($field[$j])." ".getGoType($type_array[getField($field[$j])]);

                }

                $ret_manager .= ") *${table} {
        ";

                $ret_manager .= "
    if p.Conn == nil {
        return nil
    }

    query := p.GetQuery() + \" where 1=1 \"
    var params []interface{}

";

                for($j = 0; $j < count($field); $j++) {
                    $t = getGoType($type_array[getField($field[$j])]);
                    if ($t == 'string')
                        $tc = '!= ""';
                    else
                        $tc = '!= 0';

                    if (substr($field[$j], 0, 2) == '<>') {
                        $ret_manager .= "\tif ".getField2($field[$j])."1 != \"\"  && ".getField2($field[$j])."2 != \"\"  {\n\t\tquery += \" and (".$prefix."_".getField($field[$j])." between ? and ?)\"\n\t\tparams = append(params, ".getField2($field[$j])."1)\n\t\tparams = append(params, ".getField2($field[$j])."2)\n\t} ";
                        $ret_manager .= "else if ".getField2($field[$j])."1 != \"\" {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." >= ?\"\n\t\tparams = append(params, ".getField2($field[$j])."1)\n\t} ";
                        $ret_manager .= "else if ".getField2($field[$j])."2 != \"\"  {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." <= ?\"\n\t\tparams = append(params, ".getField2($field[$j])."2)\n\t}\n";
                    } else if (substr($field[$j], 0, 2) == '#>')
                        $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." >= ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                    else if (substr($field[$j], 0, 2) == '#<')
                        $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." <= ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                    else if (substr($field[$j], 0, 1) == '#')
                        $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and date_format(".$prefix."_".getField($field[$j]).", '%Y-%m-%d') = ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                    else if (substr($field[$j], 0, 1) == '@') {
                        $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n";
                        $ret_manager .= "\t\tquery += \" and ".$prefix."_".getField($field[$j])." like ?\"\n";
                        $ret_manager .= "\t\t".getField2($field[$j])."_ := \"%\"+".getField2($field[$j])."+\"%\"\n";

                        $ret_manager .= "\t\tparams = append(params, ".getField2($field[$j])."_)\n";
                        $ret_manager .= "\t}\n";
                    } else if (substr($field[$j], 0, 1) == '>')
                        $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." > ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                    else if (substr($field[$j], 0, 1) == '<')
                        $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." < ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                    else if (substr($field[$j], 0, 1) == '^')
                        $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and not ".$prefix."_".getField($field[$j])." = ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                    else
                        $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." = ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";

                }

                $ret_manager .= "

    rows, err := QueryArray(p.Conn, query, params)

    if err != nil {
        log.Printf(\"query error : %v, %v\\n\", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRow(rows)
}
";
                continue;
            }

            $ret_manager .= "
func (p *${table}Manager) GetListBy$fieldname(";

            for($j = 0; $j < count($field); $j++)
            {
                if (substr($field[$j], 0, 2) == '<>')
                    $ret_manager .= getField($field[$j])."1 string, ". getField($field[$j])."2 string, ";
                else
                    $ret_manager .= getField2($field[$j])." ".getGoType($type_array[getField($field[$j])]).", ";
            }

            $ret_manager .= "page int, pagesize int, orderby string) *[]${table} {
        ";

            $ret_manager .= "
    if p.Conn == nil {
        return nil
    }

    startpage := (page - 1) * pagesize
    query := p.GetQuery() + \" where 1=1 \"
    var params []interface{}

";

            for($j = 0; $j < count($field); $j++)
            {
                $t = getGoType($type_array[getField($field[$j])]);
                if ($t == 'string')
                    $tc = '!= ""';
                else
                    $tc = '!= 0';


                if (substr($field[$j], 0, 2) == '<>') {
                    $ret_manager .= "\tif ".getField2($field[$j])."1 != \"\"  && ".getField2($field[$j])."2 != \"\"  {\n\t\tquery += \" and (".$prefix."_".getField($field[$j])." between ? and ?)\"\n\t\tparams = append(params, ".getField2($field[$j])."1)\n\t\tparams = append(params, ".getField2($field[$j])."2)\n\t} ";
                    $ret_manager .= "else if ".getField2($field[$j])."1 != \"\" {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." >= ?\"\n\t\tparams = append(params, ".getField2($field[$j])."1)\n\t} ";
                    $ret_manager .= "else if ".getField2($field[$j])."2 != \"\"  {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." <= ?\"\n\t\tparams = append(params, ".getField2($field[$j])."2)\n\t}\n";
                } else if (substr($field[$j], 0, 2) == '#>')
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." >= ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                else if (substr($field[$j], 0, 2) == '#<')
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." <= ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                else if (substr($field[$j], 0, 1) == '#')
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and date_format(".$prefix."_".getField($field[$j]).", '%Y-%m-%d') = ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                else if (substr($field[$j], 0, 1) == '@') {
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n";
                    $ret_manager .= "\t\tquery += \" and ".$prefix."_".getField($field[$j])." like ?\"\n";
                    $ret_manager .= "\t\t".getField2($field[$j])."_ := \"%\"+".getField2($field[$j])."+\"%\"\n";

                    $ret_manager .= "\t\tparams = append(params, ".getField2($field[$j])."_)\n";
                    $ret_manager .= "\t}\n";
                } else if (substr($field[$j], 0, 1) == '>')
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." > ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                else if (substr($field[$j], 0, 1) == '<')
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." < ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                else if (substr($field[$j], 0, 1) == '^')
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and not ".$prefix."_".getField($field[$j])." = ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                else
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." = ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";

            }

            $ret_manager .= "

    if page > 0 && pagesize > 0 {
        if orderby == \"\" {
            orderby = \"".$prefix."_id desc\"
        } else {
            orderby = \"".$prefix."_\" + orderby
        }
        query += \" order by \" + orderby
        if config.Database == \"mysql\" {
            query += \" limit ? offset ?\"
            params = append(params, pagesize)
            params = append(params, startpage)
        } else if config.Database == \"mssql\" || config.Database == \"sqlserver\" {
            query += \"OFFSET ? ROWS FETCH NEXT ? ROWS ONLY\"
            params = append(params, startpage)
            params = append(params, pagesize)
        }
    } else {
        if orderby == \"\" {
            orderby = \"".$prefix."_id\"
        } else {
            orderby = \"".$prefix."_\" + orderby
        }
        query += \" order by \" + orderby
    }

    rows, err := QueryArray(p.Conn, query, params)

    if err != nil {
        log.Printf(\"query error : %v, %v\\n\", err, query)
        return nil
    }

    defer rows.Close()

    return p.ReadRows(rows)
}
";

    $ret_manager .= "
func (p *${table}Manager) GetCountBy$fieldname(";

    for($j = 0; $j < count($field); $j++)
    {
        if ($j > 0) $ret_manager .= ", ";

        if (substr($field[$j], 0, 2) == '<>')
            $ret_manager .= getField($field[$j])."1 string, ". getField($field[$j])."2 string";
        else
            $ret_manager .= getField2($field[$j])." ".getGoType($type_array[getField($field[$j])]);
    }

    $ret_manager .= ") int {";

    $ret_manager .= "
    if p.Conn == nil {
        return 0
    }

    var params []interface{}
    query := p.GetQuerySelect() + \" where 1=1 \"
";

    for($j = 0; $j < count($field); $j++)
    {
        $t = getGoType($type_array[getField($field[$j])]);
        if ($t == 'string')
            $tc = '!= ""';
        else
            $tc = '!= 0';


        if (substr($field[$j], 0, 2) == '<>') {
            $ret_manager .= "\tif ".getField2($field[$j])."1 != \"\"  && ".getField2($field[$j])."2 != \"\"  {\n\t\tquery += \" and (".$prefix."_".getField($field[$j])." between ? and ?)\"\n\t\tparams = append(params, ".getField2($field[$j])."1)\n\t\tparams = append(params, ".getField2($field[$j])."2)\n\t} ";
            $ret_manager .= "else if ".getField2($field[$j])."1 != \"\" {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." >= ?\"\n\t\tparams = append(params, ".getField2($field[$j])."1)\n\t} ";
            $ret_manager .= "else if ".getField2($field[$j])."2 != \"\"  {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." <= ?\"\n\t\tparams = append(params, ".getField2($field[$j])."2)\n\t}\n";
        } else if (substr($field[$j], 0, 2) == '#>')
            $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." >= ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
        else if (substr($field[$j], 0, 2) == '#<')
            $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." <= ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
        else if (substr($field[$j], 0, 1) == '#')
            $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and date_format(".$prefix."_".getField($field[$j]).", '%Y-%m-%d') = ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
        else if (substr($field[$j], 0, 1) == '@') {
            $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n";
            $ret_manager .= "\t\tquery += \" and ".$prefix."_".getField($field[$j])." like ?\"\n";
            $ret_manager .= "\t\t".getField2($field[$j])."_ := \"%\"+".getField($field[$j])."+\"%\"\n";

            $ret_manager .= "\t\tparams = append(params, ".getField2($field[$j])."_)\n";
            $ret_manager .= "\t}\n";
        } else if (substr($field[$j], 0, 1) == '>')
            $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." > ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
        else if (substr($field[$j], 0, 1) == '<')
            $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." < ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
         else if (substr($field[$j], 0, 1) == '^')
            $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and not ".$prefix."_".getField($field[$j])." = ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
        else
            $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." = ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";

    }

    $ret_manager .= "
    rows, err := QueryArray(p.Conn, query, params)

    if err != nil {
        log.Printf(\"query error : %v, %v\\n\", err, query)
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
";

        }
    }


    if (isset($sum_array[strtolower($table)])) {
        $zrows = count($sum_array[strtolower($table)]);
        for($i = 0; $i < $zrows; $i++) {
            $field = $sum_array[strtolower($table)][$i]['where'];
            $funcname = ucfirst($sum_array[strtolower($table)][$i]['name']);
            if (isset($sum_array[strtolower($table)][$i]['struct'])) {
                $funcname = ucfirst($sum_array[strtolower($table)][$i]['struct']);
            }

            if (!isset($sum_array[strtolower($table)][$i]['struct'])) {
            $ret .= "
type ${table}Sum${funcname} struct {
";
            }

            if (!is_array($sum_array[strtolower($table)][$i]['sum'])) {
                $sum_array[strtolower($table)][$i]['sum'] = array($sum_array[strtolower($table)][$i]['sum']);
            }

            $sumnames = array();
            for ($j = 0; $j < count($sum_array[strtolower($table)][$i]['sum']); $j++) {
                $sumname = ucfirst($sum_array[strtolower($table)][$i]['sum'][$j]);
                $sumnames[] = $sum_array[strtolower($table)][$i]['sum'][$j];

                if ($type_array[strtolower($sumname)] == 'int' || $type_array[strtolower($sumname)] == 'bigint')
                    $sumtype = 'int64';
                else
                    $sumtype = 'float64';

                if (!isset($sum_array[strtolower($table)][$i]['struct'])) {
                    $ret .=
"    ${sumname} $sumtype
";
                }
            }

            if (!isset($sum_array[strtolower($table)][$i]['struct'])) {
                $ret .= "}
";
            }



            if (!is_array($field))
                $field = array($field);

            $fieldname = '';
            for($j = 0; $j < count($field); $j++)
            {
                $fieldname .= ucfirst(getField($field[$j]));
            }


            $ret_manager .= "
func (p *${table}Manager) GetSum{$funcname}By$fieldname(";

            for($j = 0; $j < count($field); $j++)
            {
                if ($j > 0) $ret_manager .= ", ";

                if (substr($field[$j], 0, 2) == '<>')
                    $ret_manager .= getField($field[$j])."1 string, ". getField($field[$j])."2 string";
                else
                    $ret_manager .= getField2($field[$j])." ".getGoType($type_array[getField($field[$j])]);
            }

            $ret_manager .= ") *{$table}Sum{$funcname} {";

            $ret_manager .= "
    if p.Conn == nil {
        return nil
    }

    var params []interface{}
    query := \"select ";

            for ($j = 0; $j < count($sumnames); $j++) {
                if ($j > 0)
                    $ret_manager .= ", ";

                $ret_manager .= 'sum(' .$prefix . '_' . $sumnames[$j] . ')';
            }


            if ($table_type == 'view')
                $typestring = '_vw';
            else
                $typestring = '_tb';
            $ret_manager .= " from " . strtolower($table) . "$typestring where 1=1 \"
";

            for($j = 0; $j < count($field); $j++)
            {
                $t = getGoType($type_array[getField($field[$j])]);
                if ($t == 'string')
                    $tc = '!= ""';
                else
                    $tc = '!= 0';


                if (substr($field[$j], 0, 2) == '<>') {
                    $ret_manager .= "\tif ".getField2($field[$j])."1 != \"\"  && ".getField2($field[$j])."2 != \"\"  {\n\t\tquery += \" and (".$prefix."_".getField($field[$j])." between ? and ?)\"\n\t\tparams = append(params, ".getField2($field[$j])."1)\n\t\tparams = append(params, ".getField2($field[$j])."2)\n\t} ";
                    $ret_manager .= "else if ".getField2($field[$j])."1 != \"\" {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." >= ?\"\n\t\tparams = append(params, ".getField2($field[$j])."1)\n\t} ";
                    $ret_manager .= "else if ".getField2($field[$j])."2 != \"\"  {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." <= ?\"\n\t\tparams = append(params, ".getField2($field[$j])."2)\n\t}\n";
                } else if (substr($field[$j], 0, 2) == '#>')
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." >= ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                else if (substr($field[$j], 0, 2) == '#<')
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." <= ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                else if (substr($field[$j], 0, 1) == '#')
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and date_format(".$prefix."_".getField($field[$j]).", '%Y-%m-%d') = ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                else if (substr($field[$j], 0, 1) == '@') {
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n";
                    $ret_manager .= "\t\tquery += \" and ".$prefix."_".getField($field[$j])." like ?\"\n";
                    $ret_manager .= "\t\t".getField2($field[$j])."_ := \"%\"+".getField($field[$j])."+\"%\"\n";

                    $ret_manager .= "\t\tparams = append(params, ".getField2($field[$j])."_)\n";
                    $ret_manager .= "\t}\n";
                } else if (substr($field[$j], 0, 1) == '>')
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." > ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                else if (substr($field[$j], 0, 1) == '<')
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." < ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";
                else
                    $ret_manager .= "\tif ".getField2($field[$j])." $tc {\n\t\tquery += \" and ".$prefix."_".getField($field[$j])." = ?\"\n\t\tparams = append(params, ".getField2($field[$j]).")\n\t}\n";

            }

            $ret_manager .= "
    rows, err := QueryArray(p.Conn, query, params)

    if err != nil {
        log.Printf(\"query error : %v, %v\\n\", err, query)
        return nil
    }

    defer rows.Close()

    var item {$table}Sum{$funcname}

    if !rows.Next() {
        return &item
    }

    err = rows.Scan(";

            for ($j = 0; $j < count($sumnames); $j++) {
                if ($j > 0)
                    $ret_manager .= ", ";

                $ret_manager .= '&item.' . ucfirst($sumnames[$j]);
            }

            $ret_manager .= ")

    if err != nil {
        return nil
    } else {
        return &item
    }
}
";


        }
    }







    /*

    if (isset($increase_array[strtolower($table)]))
    {
        for($i = 0; $i < count($increase_array[strtolower($table)]); $i++)
        {
            $field = $increase_array[strtolower($table)][$i];

            if (!is_array($field))
                $field = array($field);

            $fieldname = '';
            for($j = 0; $j < count($field); $j++)
            {
                $fieldname .= ucfirst($field[$j]);
            }

            $ret_manager .= "
    public function increase$fieldname(id";

            for($j = 0; $j < count($field); $j++)
            {
                $ret_manager .= ", ".$field[$j];
            }

            $ret_manager .= ")
    {
        db = this->getDatabase();

";

            for($j = 0; $j < count($field); $j++)
            {
                $ret_manager .= "\t\tupdate[] = \"${prefix}_".$field[$j]." = ${prefix}_".$field[$j]." + :".$field[$j]."\";\n";
            }

            $ret_manager .= "
        stmt = db->prepare(\"update $table_name set \" . implode(', ', update). \" where ${prefix}_id = :id\");

";

            for($j = 0; $j < count($field); $j++)
            {
                $ret_manager .= "\t\tstmt->bindParam(':".$field[$j]."', ".$field[$j].", PDO::PARAM_INT);";
            }

            $ret_manager .= "

        stmt->bindParam(':id', id, PDO::PARAM_INT);

        stmt->execute();
    }";

        }

    }





    if (isset($sum_array[strtolower($table)]))
    {
        for($i = 0; $i < count($sum_array[strtolower($table)]); $i++)
        {
            $field = $sum_array[strtolower($table)][$i];
            $sum_field = $sum2_array[strtolower($table)][$i];

            if (!is_array($field))
                $field = array($field);

            $fieldname = '';
            for($j = 0; $j < count($field); $j++)
            {
                $fieldname .= ucfirst(getField($field[$j]));
            }

            $ret_manager .= "
    public function sum".ucfirst($sum_field)."By$fieldname(";

            for($j = 0; $j < count($field); $j++)
            {
                if ($j > 0) $ret_manager .= ", ";
                $ret_manager .= "".getField($field[$j]);
            }

            $ret_manager .= ")
    {
        db = this->getDatabase();

        query = \"select sum(".$prefix.'_'.$sum_field.") from $table_name where 1=1 \";
";

            for($j = 0; $j < count($field); $j++)
            {
                if ($type_array[getField($field[$j])] == 'int' || $type_array[getField($field[$j])] == 'bigint')
                    $ret_manager .= "\t\tif ".getField($field[$j])." > 0) query += \" and ".$prefix."_".getField($field[$j])." = :".getField($field[$j])."\";\n";
                else
                {
                    if (substr($field[$j], 0, 2) == '<>') {
                        $ret_manager .= "\t\tif ".getField($field[$j])."1 && ".getField($field[$j])."2) query += \" and (".$prefix."_".getField($field[$j])." between :".getField($field[$j])."1 and :".getField($field[$j])."2)\";\n";
                        $ret_manager .= "\t\telse if ".getField($field[$j])."1) query += \" and ".$prefix."_".getField($field[$j])." >= :".getField($field[$j])."1\";\n";
                        $ret_manager .= "\t\telse if ".getField($field[$j])."2) query += \" and ".$prefix."_".getField($field[$j])." <= :".getField($field[$j])."2\";\n";
                    } else if (substr($field[$j], 0, 1) == '@') {
                        $ret_manager .= "\t\tif ".getField($field[$j]).")\n";
                        $ret_manager .= "\t\t{\n";
                        $ret_manager .= "\t\t\tquery += \" and ".$prefix."_".getField($field[$j])." like :".getField($field[$j])."\";\n";
                        $ret_manager .= "\t\t\t".getField($field[$j])."_ = '%'.".getField($field[$j]).".'%';\n";
                        $ret_manager .= "\t\t}\n";
                    }
                    else
                        $ret_manager .= "\t\tif ".getField($field[$j]).") query += \" and ".$prefix."_".getField($field[$j])." = :".getField($field[$j])."\";\n";
                }
            }

            $ret_manager .= "
        stmt = db->prepare(query);
";

            for($j = 0; $j < count($field); $j++)
            {
                if ($type_array[getField($field[$j])] == 'int' || $type_array[getField($field[$j])] == 'bigint')
                    $ret_manager .= "\t\tif ".getField($field[$j])." > 0) stmt->bindParam(':".getField($field[$j])."', ".getField($field[$j]).", PDO::PARAM_INT);\n";
                else
                {
                    if (substr($field[$j], 0, 1) == '@')
                        $ret_manager .= "\t\tif ".getField($field[$j]).") stmt->bindParam(':".getField($field[$j])."', ".getField($field[$j])."_, PDO::PARAM_STR);\n";
                    else
                        $ret_manager .= "\t\tif ".getField($field[$j]).") stmt->bindParam(':".getField($field[$j])."', ".getField($field[$j]).", PDO::PARAM_STR);\n";
                }
            }


            $ret_manager .= "
        stmt->execute();

        result = stmt->fetch(PDO::FETCH_BOTH);

        return result[0];
    }

    ";
        }

    }




    $ret_manager .= "
}

?>";
    */


    // list end -----------------------------------------------------------------------------------------------------------------


    return array($ret, $ret_manager);
}

$argc = count($argv) - 1;

if ($argc == 0) {
    $cwd = getcwd();
} else {
    $cwd = $argv[1];
}

if (is_file($cwd . '/../config/config.json')) {
    $json = json_decode(file_get_contents($cwd . '/../config/config.json'));
    $str = $json->connectionString;
} else {
    $str = "anb:anbdb@tcp(netb.co.kr:3306)/anb2";
}

preg_match('/(^[^:]+):([^@]*)@tcp\\(([^:]+)[^\\/]+\\/([-_0-9a-zA-Z]+)/', $str, $match);

$GLOBAL_VAR['db_host'] = $match[3];
$GLOBAL_VAR['db_name'] = $match[4];
$GLOBAL_VAR['db_user'] = $match[1];
$GLOBAL_VAR['db_pass'] = $match[2];

$config_file = $cwd . '/' . 'model.config.inc';
if (!is_file($config_file)) {
    echo "config file not found\n";
    exit;
}

require_once $config_file;

db_connect();

$query = "select table_name as table_name from information_schema.tables where table_schema = '".$GLOBAL_VAR['db_name']."'";
$result = $_db->query($query);

foreach($result as $row)
{
    //if ($row['table_name'] != 'review_tb')
    //    continue;

    //echo $row['table_name'];
    //echo "\n";

    //$row['table_name'] = $row['TABLE_NAME'];

    $ret = process($row['table_name']);

    $table = str_replace("_tb", '', $row['table_name']);
    $table = str_replace("_vw", '', $table);

    if ($ret == null)
        continue;

    $head = "package models

import (
    \"anb/config\"
    \"database/sql\"
    \"errors\"
    \"fmt\"
    \"log\"
    \"strings\"
    \"time\"

    //_ \"github.com/denisenkom/go-mssqldb\"
    _ \"github.com/go-sql-driver/mysql\"
    _ \"github.com/mattn/go-sqlite3\"
)

";

    $c = $head . $ret[0] . $ret[1];

    $filename = $cwd . "/../models/${table}.go";
    file_put_contents($filename, $c);
}

$str = 'package models

func InitCache() {
    /*
    db := GetConnection()

    if db == nil {
        return
    }

    defer db.Close()
    */
';


for ($i = 0; $i < count($cache_array); $i++) {
    $name = $cache_array[$i];
    $name2 = ucfirst($name);

    $str .= "
    ${name}Manager := New${name2}Manager(db)
    ${name}s := ${name}Manager.GetList(0, 0, \"\")

    for _, item := range (*${name}s) {
        data := item
        _cache{$name2}s[item.Id] = &data
    }
";
}

$str .= "
}
";

for ($i = 0; $i < count($cache_array); $i++) {
    $name = $cache_array[$i];
    $name2 = ucfirst($name);

    $str .= "
func Get${name2}s() map[int64]*${name2} {
    return _cache${name2}s
}

func Get${name2}(id int64) *${name2} {
    if item, ok := _cache${name2}s[id]; ok {
        return item
    }

    return nil
}
";

    if (array_key_exists($name, $cache_search_array)) {
        for ($j = 0; $j < count($cache_search_array[$name]); $j++) {
            $funcname = getField($cache_search_array[$name][$j]);
            $funcname2 = ucfirst($funcname);

            $query = "select column_name as column_name, data_type as data_type from information_schema.columns where table_schema = '".$GLOBAL_VAR['db_name']."' and table_name = '${name}_tb'";
            $result = $_db->query($query);

            $gotype = '';
            foreach ($result as $row) {
                //$row['column_name'] = $row['COLUMN_NAME'];
                //$row['data_type'] = $row['DATA_TYPE'];

                $temp = explode('_', $row['column_name']);

                if ($temp[1] == $funcname) {
                    $gotype = getGoType($row['data_type']);
                    break;
                }
            }

            if (substr($cache_search_array[$name][$j], 0, 1) == '!') {
            $str .= "
func Get${name2}By${funcname2}(val $gotype) *${name2} {
    var ret *${name2}
    for _, item := range _cache${name2}s {
        if item.${funcname2} == val {
            ret = item
            return ret
        }
    }

    return nil
}
";

            } else {
            $str .= "
func Get${name2}sBy${funcname2}(val $gotype) []*${name2} {
    var items []*${name2}
    for _, item := range _cache${name2}s {
        if item.${funcname2} == val {
            items = append(items, item)
        }
    }

    return items
}
";

            }
        }
    }
}


$filename = $cwd . "/../models/cache.go";
file_put_contents($filename, $str);

?>
