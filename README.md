# SprintPOINT converter #

A program to generate charts of sprint points over time.

## Usage ##

    go run src/*.go -name test.json -sprint sprint01 -add "start=44" -when today

meaning:

* (-name) use test.json for data
* (-sprint) in sprint01
* (-when) use todays date
* (-add) add value 44 for starting sprint points

### Commands ###

* add (in development)
* report (future)
    * count
    * average
    * min
    * max
* remove (future)
* list (future)

## Design ##
* To be able to query by date, sort by sprint and chart the sprint points
* To be able to query by date, sort by sprint and chart the avg sprint points

date name=value

2017-02-09 "sprint 74" sprint.all=34
2017-02-09 "sprint 74" sprint.finish=34
2017-02-09 "sprint 74" sprint.push=34

select 2017-02* "sprint 74" sprint.all value
select 2017-02-13:2017-02-23 "sprint 74" sprint.all value


----


raw notes:

count, average, min, max

    [
        {when:"", sprint:"", name:"", value:""}
        ,{when:"", sprint:"", name:"", value:""}
    ]
