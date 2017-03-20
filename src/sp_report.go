package main

import (//"os"
    "fmt"
    //"time"
    //"strings"
    //"strconv"
    //"flag"
    //"encoding/json"
    //"io/ioutil"
    )

func report_handleFlags() {

}

/**
actions:
simple
csv
text
*/
func report_work(app App_Data, options string) {
    report_handleFlags()
    data := readData(app.file_name)
    
    /*
    opt := strings.Split(options, "=")
    name := opt[0]
    value, err := strconv.Atoi(opt[1])
    if err!=nil {
        value = -1
    }*/
    
    if options=="simple" {
    
        for _, obj := range data.Points {
            fmt.Printf("%s\n", obj.toString())
        }
    } else if options=="csv" {
        for _, obj := range data.Points {
            fmt.Printf("%s\n", obj.ToCsv())
        }
    } else if options=="tab" {
        for _, obj := range data.Points {
            fmt.Printf("%s\n", obj.ToTab())
        }
    } else if options=="chart=text" {
        fmt.Printf(ChartText(app, data))
    }
    
}

func ChartText(app App_Data, data Data) string {
    out := ""
    grid := [24][40]string{}
    
    //grid[0][0] = "24"
    //grid[23][0] = "0"
    
    found := 1
    for _, obj := range data.Points {
        if (obj.Name=="Start") {
            grid[obj.Value][found] = " X"
            found = found + 1
        }
    }
    
    for y:= range grid {
        out = fmt.Sprintf("%s%.2d|", out, 24-y)
        for x:= range grid[y]{
            out = fmt.Sprintf("%s%s", out, grid[y][x])
        }
        out = out + "\n"
    }
    
    data.UniqueSprints()
    
    return out
}

/*
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
*/