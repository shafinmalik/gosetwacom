# gosetwacom
CUI Drawing tablet configuration tool for linux. Uses and requires xsetwacom to be installed. 

# main
Golang CUI for xsetwacom to easily create and store settings for drawing tablet.

Will need to integrate bash into go. 
Hint: once bash is integrated the following command will be a useful starting point:

```
for device in $(xsetwacom list devices)
  do echo "$device"
done
```
