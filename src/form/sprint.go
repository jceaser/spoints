package main

import ("os"
    "fmt"
    "math"
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
    velocity float64
	load float64
    reserve float64
    reserve_and_maintenance float64
    
}

func (a App_Data) String() string {
    return fmt.Sprintf("\nSummary:\nC=%.2f\nV=%.2f\nL=%.2f\nR=%.2f\n",
    	 a.capacity,
    	 a.velocity,
    	 a.load,
    	 a.reserve_and_maintenance)
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

func help(defaults App_Defaults) {	
	var _, w = termSize()
	flex := int(math.Min(float64(48), float64(w-35))) //20 on left, 15 on right

	format1 := "%4s %-16s  %-" + strconv.Itoa(flex) + "s [%s]\n"
	format2 := "%4s %-16s  %-" + strconv.Itoa(flex) + "s [%2.1f]\n"
	
	fmt.Printf("sprint -H | [ -D -w -d -h -v -p -r -m ] [-i | -I]\n")
	fmt.Printf(format1, "Flag","Long Flag","Description","Default")
	fmt.Printf(format2, "-w", "--workers", "Workers In Sprint", defaults.workers)
	fmt.Printf(format2, "-d", "--days", "Days in Sprint", defaults.days)
	fmt.Printf(format2, "-h", "--holidays", "Holiday Days", defaults.holidays)
	fmt.Printf(format2, "-v", "--vacations", "Vacation Days", defaults.vacation)
	fmt.Printf(format2, "-p", "--points-per-day", "Points to days conversion", defaults.points_per_day)
	fmt.Printf(format2, "-r", "--reserve", "Points to hold back in reserve", defaults.reserve)
	fmt.Printf(format2, "-m", "--maintenance", "% of maintenance", defaults.maintenance)
	fmt.Printf(format2, "-D", "--dump", "don't ask values, just dump", defaults.maintenance)
	fmt.Printf(format2, "-H", "--help", "This help", defaults.maintenance)
	fmt.Printf(format1, "-i", "--inline", "Inline mode, clear screen, one question at time", "")
	fmt.Printf(format1, "-I", "--not-inline", "Not Inline mode, don't manipulate position", "default")
	os.Exit(0)
}

func work() {
	defaults := App_Defaults{5.0, 10.0, 0, 0, 1, 8, 40, true}
    
    var a = ""
    var dump = false
    
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
		
		if arg=="--dump" || arg=="-D" {dump=true;continue}
		
		if arg=="--inline" || arg=="-i" {
			a = "inline"
			continue;
		}
		if arg=="--not-inline" || arg=="-I" {
			a = ""
			continue;
		}
		
		if arg=="--help" || arg=="-H" {help(defaults);}		
		
	}
	
	if dump {
		app_data.workers = defaults.workers
		app_data.days = defaults.days
		app_data.holidays = defaults.holidays
		app_data.vacations = defaults.vacation
		app_data.points = defaults.points_per_day
		app_data.reserve = defaults.reserve
		app_data.maintenance = defaults.maintenance
	} else {
	    if (a!=""){termClear()}
    
		assignNum(a, "Num of developers (%.2f) ", &app_data.workers, defaults.workers)
		assignNum(a, "Num of days in sprint(%.2f) ", &app_data.days, defaults.days)
		assignNum(a, "Holiday days (%.2f) ", &app_data.holidays, defaults.holidays)
		assignNum(a, "Vacation days (%.2f) ", &app_data.vacations, defaults.vacation)
		assignNum(a, "Sprint points in a day (%.2f) ", &app_data.points, defaults.points_per_day)
		assignNum(a, "Reserve points (%.2f) ", &app_data.reserve, defaults.reserve)
		assignNum(a, "Maintenance percentage (%.2f) ", &app_data.maintenance, defaults.maintenance)
    }
    
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
    
    app_data.load = schedule_points
    app_data.reserve_and_maintenance = app_data.reserve + maintenance_points
    
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
    	
    /*fmt.Printf("\n\n C=%.2f\n V=%.2f\n L=%.2f\n R=%.2f\n",
    	 app_data.capacity,
    	 app_data.velocity,
    	 schedule_points,
    	 maintenance_points)
	*/
	fmt.Printf(app_data.String())
}

func main() {
    handleFlags()
    
    work()
}
