package main

import (//"os"
    //"fmt"
    //"time"
    "strings"
    //"strconv"
    //"flag"
    //"encoding/json"
    //"io/ioutil"
    )

func rm_handleFlags() {

}

/*
option is sprint:value
*/
func rm_work(app App_Data, options string) {
    add_handleFlags()
    
    if app.sprint=="" {
        app.sprint = "None"
    }
    
    if app.value=="" {
        app.value = "None"
    }
    
    data := readData(app.file_name)
    
    opt := strings.Split(options, ":")
    name := opt[0]
    
    affected := 0
    for i := len(data.Points)-1; 0<=i; i-- {
        v := data.Points[i]
        
        if v.Sprint==app.sprint && v.Name==name {
            affected += data.remove(i, v)
        }
    }
    
    if !app.dry_run {
        writeData(app.file_name, data)
    }
}
