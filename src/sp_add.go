package main

import (//"os"
    "fmt"
    "time"
    "strings"
    "strconv"
    //"flag"
    //"encoding/json"
    //"io/ioutil"
    )

func add_handleFlags() {

}

/**
Do the Add process
@param app current app status
@param options what to add, should be name1=value1,name2=value2
*/
func add_work(app App_Data, options string) {
    add_handleFlags()
    
    if app.when=="now" {
        app.when = now()
    } else if app.when=="today" {
        app.when = today()
    }
    
    if app.sprint=="" {
        app.sprint = "None"
    }
    
    data := readData(app.file_name)
    
    for _, pair := range strings.Split(options, ",") {
        opt := strings.Split(pair, "=")
        name := opt[0]
        value, err := strconv.Atoi(opt[1])
        if err!=nil {value = -1}
    
        row := CreateRow(app.when, app.sprint, name, value)
        data.add(row)
    }
    
    if !app.dry_run {
        writeData(app.file_name, data)
    }
}

/**
@return current ISO date time
*/
func now() string {
    n := time.Now()
    y, m, d := n.Date()
    h := n.Hour()
    mm := n.Minute()
    s := n.Second()
    out := fmt.Sprintf("%.4d-%.2d-%.2dT%.2d:%.2d:%.2d", y, m, d, h, mm, s)
    return out
}

/**
@return current ISO date
*/
func today() string {
    n := time.Now()
    y, m, d := n.Date()
    return fmt.Sprintf("%.4d-%.2d-%.2d", y, m, d)
}