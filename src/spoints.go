package main

/*
find files in the pattern of name.ext and roll them back to name.num.ext. Any name.num.ext should also be rolled back by one

This class dispatches to the handler that will handle the actual request, current requests are 

CRUD

* add (Create), (Update)
* report (Read)
* remove (delete)
* help

p|             x
o|         x
i| x               x
n|     x
t|                     x
s|
-------------------------------
   s1  s2  s3  s4  s5  s6

*/

import ("os"
    "fmt"
    "flag"
    //"path/filepath"
    )

type App_Data struct {
    file_name string
    dry_run bool
    mode_init bool
    mode_add string     //start=44
    mode_remove string  //sprint:name
    mode_report bool
    mode_stats bool
    
    
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
    raw_dry := flag.Bool("dry-run", false, "don't write any new data")
    raw_add := flag.String("add", "", "add a new data point")
    raw_remove := flag.String("remove", "", "remove a data point")
    raw_report := flag.Bool("report", false, "generate a report")
    raw_stats := flag.Bool("stats", false, "print out some basic statistics")
    
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
    app_data.dry_run = *raw_dry
    app_data.mode_add = *raw_add
    app_data.mode_remove = *raw_remove
    app_data.mode_report = *raw_report
    app_data.mode_stats = *raw_stats
    
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
        
        if app_data.mode_add != "" {
            //run add 
            add_work(app_data, app_data.mode_add)
        }
        if app_data.mode_remove != "" {
            rm_work(app_data, app_data.mode_remove)
        }
        if app_data.mode_report {
            //run report command
        }
        if app_data.mode_stats {
            stats_work(app_data, "")
        }
    } else {//nothing to do
        if app_data.mode_init {
            //create data file
            writeData(app_data.file_name, CreateData(1.0))
        } else {
            fmt.Printf(ERR_MSG_01, app_data.file_name)
        }
    }
}

/*
*/
func main() {
    handleFlags()
    
    work()
}
