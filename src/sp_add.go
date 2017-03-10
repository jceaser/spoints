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

func add_work(app App_Data, options string) {
    add_handleFlags()
    
    //fmt.Printf("inside of add on '%s' for '%s' with '%s'.\n", app.when, app.sprint, options)
    
    if app.when=="now" {
        app.when = now()
    } else if app.when=="today" {
        app.when = today()
    }
    
    if app.sprint=="" {
        app.sprint = "None"
    }
    
    data := readData(app.file_name)
    
    opt := strings.Split(options, "=")
    name := opt[0]
    value, err := strconv.Atoi(opt[1])
    if err!=nil {
        value = -1
    }
    
    row := CreateRow(app.when, app.sprint, name, value)
    //data.Points = append(data.Points, row)
    data.add(row)
    
    //writeData("out."+app.file_name, data)
    writeData(app.file_name, data)
}

func now() string {
    n := time.Now()
    y, m, d := n.Date()
    h := n.Hour()
    mm := n.Minute()
    s := n.Second()
    out := fmt.Sprintf("%.4d-%.2d-%.2dT%.2d:%.2d:%.2d", y, m, d, h, mm, s)
    return out
}

func today() string {
    n := time.Now()
    y, m, d := n.Date()
    return fmt.Sprintf("%.4d-%.2d-%.2d", y, m, d)
}