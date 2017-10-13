package main

import ("os"
    "fmt"
    //"time"
    //"strings"
    "strconv"
    "os/exec"
    //"flag"
    //"encoding/json"
    //"io/ioutil"
    )


/**
v-38
r-10
l-28
*/

type App_Data struct {
    workers int
    days int
    holidays int
    vacations int
    maintenance int
    
    velocity int
    reserve int
    load int
}

const (
    ERR_MSG_01 = "01: File '%s' does not exist"
)

var app_data App_Data

func handleFlags() {
}

func add_handleFlags() {

}

func assignNum(msg string, field *int, def int) {
    var raw string
    
    fmt.Printf(msg)
    fmt.Scanln(&raw)
    if (raw=="") {
        *field = def
    } else {
        *field, _ = strconv.Atoi(raw)
    }
    fmt.Printf("\033[1;1H")
    var _, w = termSize()
    for i:=0; i<w; i++ {
        fmt.Printf(" ");
    }
    fmt.Printf("\n")
    
    //return field
}

func termClear() {
    fmt.Print("\033[2J")
}

func termMoveTo(x int, y int) {
    fmt.Printf("\033[%d;%dH", x, y)
}

func termSize() (int, int) {
    var h, w int
    cmd := exec.Command("stty", "size")
    cmd.Stdin = os.Stdin
    d, _ := cmd.Output()
    fmt.Sscan(string(d), &h, &w)
    
    return h, w;
}

func work() {
    //if exists(app_data.file_name) {
        //something to do
    //}
    //var raw string
    
    termClear()
    
    var a = "\033[1;1H"     //move to 1,1
    
    assignNum(a + "Num of developers (5) ", &app_data.workers, 5)
    assignNum(a + "Num of days (10) ", &app_data.days, 10)
    assignNum(a + "Holidays (0) ", &app_data.holidays, 0)
    assignNum(a + "Vacations (0) ", &app_data.vacations, 0)
    assignNum(a + "Reserve (10) ", &app_data.reserve, 10)
    assignNum(a + "Maintenance (40) ", &app_data.maintenance, 40)

    var tdays = app_data.workers*app_data.days
    var outage = app_data.vacations+app_data.holidays*app_data.workers
    
    app_data.velocity = tdays-outage
    app_data.load = tdays-outage-app_data.reserve
    
    var app_target = 1.0 - (float64(app_data.maintenance*1.0)/100.0)
    
    fmt.Printf("Total days for %d developers is %d, but subtract out %d outage days. velocity is %d but with a reserve of %d schedule at %f.\n",
        app_data.workers,
        tdays,
        outage,
        app_data.velocity,
        app_data.reserve,
        float64(app_data.load) * app_target)
        //(1.0*app_data.load) * app_target )
}

func main() {
    handleFlags()
    
    work()
}
