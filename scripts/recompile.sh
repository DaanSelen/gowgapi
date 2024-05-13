#!/bin/bash

echo "Compiling"

cd ../src
go build -o ../

echo "Executing"
cd ../
./gowgapi