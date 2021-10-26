# gosetwacom
CUI Drawing tablet configuration tool for linux. Uses and requires xsetwacom to be installed. 

The objective is to create a user-friendly interface for xsetwacom that allows for the creation of loadable profiles for saved devices. 

# main
Golang CUI for xsetwacom to easily create and store settings for drawing tablet.

Will need to integrate bash into go. 
Hint: once bash is integrated the following command will be a useful starting point:

```
mapfile -t device_list < <(xsetwacom list devices)

# to loop through the array: 
for element in "${device_list[@]}"
  do
    echo "{element}"
done
```
## gosetwacom.go
- package main
- reserved for CUI handling
- the three views are currently placeholders and will be replaced by "devices", "profiles", and "database"

## tgsh.go
- use to interface with bash

## ttd.go
- contains structs for managing device data and profiles