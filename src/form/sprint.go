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

type App_Defaults struct {
	workers float64
	days float64
    holidays float64
    vacation float64
    points_per_day float64
    reserve float64
    maintenance float64
    
    inline_mode bool
}

const (
    ERR_MSG_01 = "01: File '%s' does not exist"
)

var app_data App_Data

func handleFlags() {
}

func add_handleFlags() {

}

func assignNum(prefix string, msg string, field *float64, def float64) {
    var raw string
    
    if 0<len(prefix) {
    	termMoveToFirst()
    }
    
    fmt.Printf(msg, def)
    fmt.Scanln(&raw)
    if (raw=="") {
        *field = def
    } else {
        *field, _ = strconv.ParseFloat(raw, 64)
    }
    
    if 0<len(prefix) {
    	termMoveToFirst()
	    var _, w = termSize()
    	for i:=0; i<w; i++ {
        	fmt.Printf(" ");
    	}
    }
    fmt.Printf("\n")
}

func termClear() {
    fmt.Print("\033[2J")
}

func termMoveToFirst() {
	// send a "\033[1;1H"
	termMoveTo(1,1)
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

func set_def(field *float64, long_flag string, short_flag string, arg string, next string) (bool){
	if arg==long_flag || arg==short_flag {
		*field, _ = strconv.ParseFloat(next, 64)
		return true
	}
	return false
}

func work() {
	defaults := App_Defaults{5.0, 10.0, 0, 0, 1, 8, 40}
    
    var a = ""
    
    for i:=1; i<len(os.Args); i++ {
		arg := os.Args[i]
		next := ""
		if (i+1)<len(os.Args) {next = os.Args[i+1]}
		
		if set_def(&defaults.workers, "--workers", "-w", arg, next){continue}
		if set_def(&defaults.days, "--days", "-d", arg, next){continue}
		if set_def(&defaults.holidays, "--holidays", "-h", arg, next){continue}
		if set_def(&defaults.vacation, "--vacations", "-v", arg, next){continue}
		if set_def(&defaults.points_per_day, "--points-per-day", "-p", arg, next){continue}
		if set_def(&defaults.reserve, "--reserve", "-r", arg, next){continue}
		if set_def(&defaults.maintenance, "--maintenance", "-m", arg, next){continue}
		
		if arg=="--inline" || arg=="-i" {
			a = "inline"
			continue;
		}
		if arg=="--not-inline" || arg=="-I" {
			a = ""
			continue;
		}
		
	}
	    
    termClear()
    
    assignNum(a, "Num of developers (%.2f) ", &app_data.workers, defaults.workers)
    assignNum(a, "Num of days in sprint(%.2f) ", &app_data.days, defaults.days)
    assignNum(a, "Holiday days (%.2f) ", &app_data.holidays, defaults.holidays)
    assignNum(a, "Vacation days (%.2f) ", &app_data.vacations, defaults.vacation)
    assignNum(a, "Sprint points in a day (%.2f) ", &app_data.points, defaults.points_per_day)
    assignNum(a, "Reserve points (%.2f) ", &app_data.reserve, defaults.reserve)
    assignNum(a, "Maintenance percentage (%.2f) ", &app_data.maintenance, defaults.maintenance)
    
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
