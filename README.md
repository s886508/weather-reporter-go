
# Introduction
This is a simple crawler built by Go. Input city to lookup weather data from next few days.

# Start
- Run by interpreter
`go run main`
- Build binary
`go install main`

# Example
The tool run until user enter `Q`. Enter`city name` and `days` to print retrieved data.
```
Please enter city name(Q to exit): 臺北市
Please enter days to show(1~7): 5
============ 臺北市 ============

 03/13 星期三     白天:15 ~ 22°C:多雲
 03/13 星期三     晚上:17 ~ 20°C:陰時多雲短暫陣雨

 03/14 星期四     白天:17 ~ 21°C:陰短暫陣雨
 03/14 星期四     晚上:16 ~ 19°C:多雲時陰短暫陣雨

 03/15 星期五     白天:16 ~ 19°C:多雲時陰短暫雨
 03/15 星期五     晚上:16 ~ 17°C:多雲短暫雨

 03/16 星期六     白天:16 ~ 19°C:陰短暫雨
 03/16 星期六     晚上:16 ~ 18°C:陰時多雲短暫雨

 03/17 星期日     白天:16 ~ 19°C:多雲時陰短暫陣雨
 03/17 星期日     晚上:15 ~ 17°C:陰短暫陣雨

============ 臺北市 ============
```
