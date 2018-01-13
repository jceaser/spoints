package main

import ("os"
    "fmt"
    //"flag"
    //"reflect"
    "math"
    "sort"
    "strings"
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

/**************************************/
/* outputs */

func (r Row) toString() string {
    return fmt.Sprintf("w: %s, s: %s, %s=%d", r.When, r.Sprint, r.Name, r.Value)
}

func quote(raw string) string {
    return strings.Replace(raw, "\"", "\"\"", -1)
}

func (self Row) ToArray() []string {
    list := []string { self.When, self.Sprint, self.Name, fmt.Sprintf("%d",self.Value) }
    return list
}

func (self Row) ToTextList(sep string) string {
    list := self.ToArray()
    return strings.Join(list, sep)
}

func (self Row) ToTab() string {
    return self.ToTextList("\t")
}

func (self Row) ToCsv2() string {
    return self.ToTextList(", ")
}

func (self Row) ToCsv() string {
    return fmt.Sprintf("\"%s\", \"%s\", \"%s\", %d", quote(self.When),
         quote(self.Sprint), quote(self.Name), self.Value)
}

/**************************************/

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
        //add or update
        if obj.Different(r) {//add
            shouldAdd = true
        } else {
            if obj.DifferentValue(r) {
                //update it
                d.Points[idx] = r
            }
            //could be the same
            shouldAdd = false
            break
        }
    }
    if shouldAdd {
        d.Points = append(d.Points, r)
    }
}

func remove(a []Row, i int) []Row {
    b := []Row{}
    
    b = append(b, a[:i]...)
    b = append(b, a[i+1:]...)
    
    return b
}

//not working yet
func (d *Data) remove(i int, r Row) int {
    affected := 0
    
    d.Points = remove(d.Points, i)
    
    affected = 1
    
    return affected
}

func Min(running int, current int) int {
	return int (math.Min(float64(running), float64(current)))
}

func Max(running int, current int) int {
	return int(math.Max(float64(running), float64(current)))
}

func (self Data) Ranges() map[string]stat {
    stats := make(map[string]stat)
    global := stat{}
    global.Name = "global"
    global.Min = 0x7fffffffffffffff	//assume a very large number
    for _, v := range self.Points {
        if val, exists := stats[v.Name]; exists {//update
			val.Value += v.Value
			val.Count += 1
			val.Avg = val.Value/val.Count
			val.Min = Min(val.Min, v.Value)
			val.Max = Max(val.Max, v.Value)
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
        global.Value += v.Value
        global.Count += 1
        global.Min = Min(global.Min, v.Value)
        global.Max = Max(global.Max, v.Value)
    }
    global.Avg = global.Value/global.Count
    stats["global"] = global
    
    return stats
}

func (self Data) UniqueSprints() []string {
    ret := []string{}
    keys := map[string]bool{}
    for _, obj := range self.Points {
        keys[obj.Sprint]=true
    }
    
    for k, _ := range keys {
        ret = append(ret, k)
    }
    
    sort.Strings(ret)
    
    return ret
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

/**************************************/

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
