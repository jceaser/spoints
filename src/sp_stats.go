package main

import (//"os"
    "fmt"
    //"time"
    //"strings"
    //"strconv"
    //"flag"
    "math"
    "reflect"
    //"encoding/json"
    //"io/ioutil"
    )

func stats_handleFlags() {

}

type stat struct {
    Name string 
    Value int
    Count int
    Avg int
    Min int
    Max int
}

func (s stat) String() string {
    return fmt.Sprintf("%d from %d points ; %d<%d<%d.",
        s.Value, s.Count, s.Min, s.Avg, s.Max)
}


/**
Do the Add process
@param app current app status
@param options what to add, should be name1=value1,name2=value2
*/
func stats_work(app App_Data, options string) {
    stats_handleFlags()
    
    data := readData(app.file_name)
    
    uniq_sprints := make(map[string]bool)
    uniq_names := make(map[string]bool)
    count_by_name := make(map[string]int)
    stats := make(map[string]stat)
    
    for _, v := range data.Points {
        if uniq_sprints[v.Sprint]==false {
            uniq_sprints[v.Sprint] = true
        }
        if uniq_names[v.Name]==false {
            uniq_names[v.Name] = true
        }
        count_by_name[v.Name] += v.Value
        
        if val, exists := stats[v.Name]; exists {//update
            val.Value += v.Value
            val.Count += 1
            val.Avg = val.Value/val.Count
            val.Min = int(math.Min(float64(val.Min), float64(v.Value)))
            val.Max = int(math.Max(float64(val.Max), float64(v.Value)))
            stats[v.Name] = val
        } else {//add
            s := stat{}
            s.Name = v.Name
            s.Value = v.Value
            s.Count = 1
            s.Avg = s.Value
            s.Min = s.Value
            s.Max = s.Value
            stats[v.Name] = s
        }
    }
    keys := reflect.ValueOf(uniq_sprints).MapKeys()
    names := reflect.ValueOf(uniq_names).MapKeys()
    
    /*
      | Start | Stop |
    ------------------
    s1|  42   |  32  |
    s2|  24   |  23  |
    ------------------
    avg  33   |  26.5
    */
    
    fmt.Printf("%s -> %s\n", keys, names)
    for _, v := range stats {
        fmt.Printf("stats for %s: %s\n", v.Name, v)
    }
}