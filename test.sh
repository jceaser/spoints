#!/bin/bash

go run src/*.go -name unit_test.json -init
go run src/*.go -name unit_test.json -sprint s1 -add "Start=42"
go run src/*.go -name unit_test.json -sprint s1 -add "Mid=16"
go run src/*.go -name unit_test.json -sprint s1 -add "Stop=32"
#more unit_test.json
go run src/*.go -name unit_test.json -sprint s1 -remove "Mid"
#more unit_test.json
go run src/*.go -name unit_test.json -sprint s2 -add "Start=24,Stop=23"
more unit_test.json
