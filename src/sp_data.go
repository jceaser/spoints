package main

import ("os"
    "fmt"
    //"flag"
    "encoding/json"
    "io/ioutil"
    )


/*
{
    "_format":1.0
    ,"points":
    [
        {"when":"2017-01-01", "sprint":"s01", "name":"start", "value":32}
        ,{"when":"2017-01-15", "sprint":"s02", "name":"start", "value":43}
    ]
}
*/

type Row struct {
    When string `json:when`
    Sprint string `json:sprint`
    Name string `json:name`
    Value int `json:value`
}
func (r Row) toString() string {
    return fmt.Sprintf("w: %s, s: %s, %s=%d", r.When, r.Sprint, r.Name, r.Value)
}

func (self Row) DifferentValue(other Row) bool {
    ret := true  //assume different
    if self.Sprint==other.Sprint && self.Name==other.Name && self.Value==other.Value {
        ret = false //have same date
    }
    return ret
}

func (self Row) Different(other Row) bool {
    ret := true  //assume different
    if self.Sprint==other.Sprint && self.Name==other.Name {
        ret = false //have same date
    }
    return ret
}

func CreateRow(w, s, n string, v int) Row {
    m := Row{}
    m.When = w
    m.Sprint = s
    m.Name = n
    m.Value = v
    return m
}

/******************************************************************************/
/* Data */

type Data struct {
    Format float32  `json:"_format"`
    Points []Row   `json:"points"`
}

func CreateData(format float32) Data {
    obj := Data{}
    obj.Format = format
    return obj
}

/** mutable method to add a row to the list of data points */
func (d *Data) add(r Row) {
    shouldAdd := true
    for idx, obj := range d.Points {
        if obj.Different(r) {
            fmt.Printf("not the same, add it\n")
        } else {
            if obj.DifferentValue(r) {
                fmt.Printf("found one, lets update instead\n")
                d.Points[idx] = r
            }
            shouldAdd = false
            break
        }
    }
    if shouldAdd {
        d.Points = append(d.Points, r)
    }
}

func (d Data) toString() string {
    out := fmt.Sprintf("%1.1f [", d.Format)
    for i, v := range d.Points {
        if i!=0 {
            out = fmt.Sprintf("%s,%s", out, v.toString())
        } else {
            out = fmt.Sprintf("%s%s", out, v.toString())
        }
    }
    out = fmt.Sprintf("%s]", out)
    return out
}

func readData(file string) Data {
    raw, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var c Data
    json.Unmarshal(raw, &c)
    return c
}


func writeData(file string, data Data) {
    //json_data, _ := json.Marshal(data)
    json_data, _ := json.MarshalIndent(data, "", "    ")
    err := ioutil.WriteFile(file, json_data, 0644)
    if err!=nil {fmt.Printf("Error: %s\n", err)}
}
