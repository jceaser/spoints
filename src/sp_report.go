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
    grid := [24][80]string{}
    RangeY := 24
    //RangeX := 80
    
    max := data.Ranges()["global"].Max
    //min = data.Ranges()["min-max"].Min
    
    //fill grid
    found := 1
    for uidx, u := range data.UniqueSprints() {
        fmt.Printf("sprint %s\n", u)
        for _, obj := range data.Points {
            if (obj.Sprint==u && obj.Name=="Start") {
                ratio := float64(max)/float64(RangeY)
                y := int(float64(obj.Value) / ratio) -1
                
                grid[y][(uidx+1)] = "X"
                found = found + 1
            } else if obj.Sprint==u && obj.Name=="Stop" {
                ratio := float64(max)/float64(RangeY)
                y := int(float64(obj.Value) / ratio) -1
                
                grid[y][(uidx+1)] = "+"
                found = found + 1
            }
        }
    }
    
    //draw the grid out
    for y:= range grid {
        out = fmt.Sprintf("%s%.2d|", out, RangeY-y)
        x_line := ""
        for x := range grid[y]{
            cell := "."
            if 0<len(grid[y][x]) {
                cell = grid[y][x]
            }
            
            x_line = fmt.Sprintf("%s%s", x_line, cell)
        }
        out = fmt.Sprintf("%s%s\n", out, x_line)
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