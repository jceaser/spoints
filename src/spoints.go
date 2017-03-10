package main

/*
find files in the pattern of name.ext and roll them back to name.num.ext. Any name.num.ext should also be rolled back by one

This class dispatches to the handler that will handle the actual request, current requests are 

* add
* remove
* report
* help

*/

import ("os"
    "fmt"
    "flag"
    //"path/filepath"
    )

type App_Data struct {
    file_name string
    mode_init bool
    mode_add string
    mode_remove bool
    mode_report bool
    
    when string
    sprint string
    value string
}

const (
    ERR_MSG_01 = "01: File '%s' does not exist"
)

var app_data App_Data

func handleFlags() {
    orig := os.Args
    
    raw_help := flag.Bool("help", false, "help")
    raw_name := flag.String("name", "data.json", "location of data file")
    
    raw_init := flag.Bool("init", false, "add a new data point")
    raw_add := flag.String("add", "", "add a new data point")
    raw_remove := flag.Bool("remove", false, "remove a data point")
    raw_report := flag.Bool("report", false, "generate a report")
    
    raw_date := flag.String("when", "today", "when did the data point happen")
    raw_sprint := flag.String("sprint", "", "which sprint is this for")
    raw_value := flag.String("value", "", "name=value")
    
    flag.Parse()
    
    if *raw_help {
        fmt.Printf("cmd --help --name\n")
        fmt.Printf("\n")
        fmt.Printf("calculate average sprint points over time")
        os.Exit(-1)
    }
    app_data.file_name = *raw_name
    app_data.mode_init = *raw_init
    app_data.mode_add = *raw_add
    app_data.mode_remove = *raw_remove
    app_data.mode_report = *raw_report
    
    app_data.when = *raw_date
    app_data.sprint = *raw_sprint
    app_data.value = *raw_value
    
    os.Args = orig
}

func exists(path string) bool {
    var e = false
    if _, err := os.Stat(path); err == nil {e = true}
    return e;
}

/** primary entry point to task at hand */
func work() {
    if exists(app_data.file_name) {
        //something to do
        //var dir,full_name = filepath.Split(app_data.file_name)
        //var ext = filepath.Ext(full_name)
        //var name = full_name[0:len(full_name)-len(ext)]
        
        /*fmt.Printf("r:%s -> d:'%s' has '%s' of type %s.\n",
            app_data.file_name, dir, name, ext)*/
        
        if app_data.mode_add != "" {
            //run add 
            add_work(app_data, app_data.mode_add)
        }
        if app_data.mode_remove {
            //run remove command
        }
        if app_data.mode_report {
            //run report command
        }
    } else {//nothing to do
        fmt.Printf(ERR_MSG_01, app_data.file_name)
    }
}

/*
*/
func main() {
    handleFlags()
    
    work()

}
