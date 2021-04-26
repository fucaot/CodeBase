#! /bin/bash

# output: 2021年 4月26日 星期一 14时27分50秒 CST
date=$(date)
echo $date

# 参数 'y' 具体年份 / parameter '%Y' specifical year
# output: 2021
yeardate=$(date +%Y)
echo $yeardate

# 参数 '%y' 年份最后两位数 / parameter '%y' the last two digits of this year
# output: 21
yeardate2=$(date +%y)
echo $yeardate2

# 参数 '%m' 月份 / parameter '%m' month
# output: 21
month=$(date +%m)
echo $month

# 参数 '%d' 一个月的第几天 / parameter '%d' what a day of month
# output: 21
day=$(date +%d)
echo $day

# 参数 '%d' 一个月的第几天 / parameter '%d' what a day of month
# output: 21
day=$(date +%d)
echo $day
