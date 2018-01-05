#!/bin/bash

go run src/*.go -name unit_test.json -init
go run src/*.go -name unit_test.json -sprint s1 -add "Start=42"
go run src/*.go -name unit_test.json -sprint s1 -add "Mid=16"
go run src/*.go -name unit_test.json -sprint s1 -add "Stop=32"

#more unit_test.json
go run src/*.go -name unit_test.json -sprint s1 -remove "Mid"
#more unit_test.json
go run src/*.go -name unit_test.json -sprint s2 -add "Start=24,Stop=23"
go run src/*.go -name unit_test.json -sprint s2 -add "Start=22,Stop=33"

#more unit_test.json
go run src/*.go -name unit_test.json -sprint s3 -add "Start=11,Stop=11"
go run src/*.go -name unit_test.json -sprint s4 -add "Start=20,Stop=19"
go run src/*.go -name unit_test.json -sprint s5 -add "Start=15,Stop=10"
go run src/*.go -name unit_test.json -sprint s6 -add "Start=25,Stop=20"
go run src/*.go -name unit_test.json -sprint s7 -add "Start=30,Stop=23"
go run src/*.go -name unit_test.json -sprint s8 -add "Start=20,Stop=13"
go run src/*.go -name unit_test.json -sprint s9 -add "Start=24,Stop=23"
go run src/*.go -name unit_test.json -sprint s0 -add "Start=25,Stop=23"


go run src/*.go -name unit_test.json -stats "avg,min,max"
echo that was stats
read

go run src/*.go -name unit_test.json -report "csv"
echo that was csv
read

go run src/*.go -name unit_test.json -report "chart=text"