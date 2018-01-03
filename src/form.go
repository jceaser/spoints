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

type App_Data struct {
    workers float64
    days float64
    holidays float64
    vacations float64
    maintenance float64
    points float64
    
    capacity float64
    reserve float64
    velocity float64
}

const (
    ERR_MSG_01 = "01: File '%s' does not exist"
)

var app_data App_Data

func handleFlags() {
}

func add_handleFlags() {

}

func assignNum(msg string, field *float64, def float64) {
    var raw string
    
    fmt.Printf(msg)
    fmt.Scanln(&raw)
    if (raw=="") {
        *field = def
    } else {
        *field, _ = strconv.ParseFloat(raw, 64)
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
    
    assignNum(a + "Num of developers (5.0) ", &app_data.workers, 5.0)
    assignNum(a + "Num of days in sprint(10.0) ", &app_data.days, 10.0)
    assignNum(a + "Holiday days (0.0) ", &app_data.holidays, 0.0)
    assignNum(a + "Vacation days (0.0) ", &app_data.vacations, 0.0)
    assignNum(a + "Sprint points in a day (1.0) ", &app_data.points, 1.0)
    assignNum(a + "Reserve points (10.0) ", &app_data.reserve, 10.0)
    assignNum(a + "Maintenance percentage (40.0) ", &app_data.maintenance, 40.0)
    
    //clean up values
    app_data.maintenance = app_data.maintenance/100.0
	
	//calculate days
    var total_days = app_data.workers*app_data.days
    var outage_days = app_data.vacations + (app_data.holidays*app_data.workers)
    
    //calculate points
    app_data.capacity = app_data.points * (total_days-outage_days)
    app_data.velocity = app_data.capacity - app_data.reserve
    
    var app_target = 1.0 - app_data.maintenance //0.6=1-.4
    var schedule_points = app_data.velocity * app_target
    var maintenance_points = app_data.velocity * app_data.maintenance
    
    fmt.Printf(
    	"Total days for %.2f developers is %.2f, but subtract out %.2f outage days.\n",
    	app_data.workers, total_days, outage_days)
    
    fmt.Printf(
    	"Capacity is %.2f with a velocity of %.2f due to a reserve of %.2f.\n",
    	app_data.capacity, app_data.velocity, app_data.reserve)
	fmt.Printf(
		"Schedule at %.2f given a %0.1f%% target leaving %.2f points for maintenance.\n",
        schedule_points, app_target*100, maintenance_points)
    fmt.Printf("%.2f = (%.2f-%.2f) * (1-%.2f)\n",
    	schedule_points, app_data.capacity, app_data.reserve, app_data.maintenance)
}

func main() {
    handleFlags()
    
    work()
}
