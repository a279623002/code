package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	PId  int
	Date string
}

func NewDatas() []*Data {
	datas := []*Data{}
	dataA := &Data{
		PId:  1,
		Date: "2023-02-01,2023-05-26,2024-02-01,2023-05-26",
	}
	dataB := &Data{
		PId:  2,
		Date: "2023-05-26,2023-12-01,2025-01-01,2025-12-01",
	}
	dataC := &Data{
		PId:  2,
		Date: "2023-05-26,2023-12-01,2024-01-01,2025-12-01",
	}
	datas = append(datas, dataA)
	datas = append(datas, dataB)
	datas = append(datas, dataC)
	return datas
}

func Judge(data *Data) (res bool) {
	timeArr := strings.Split(data.Date, ",")
	if len(timeArr) < 1 {
		return false
	}
	// 按日期排序
	sort.Slice(timeArr, func(i, j int) bool {
		return timeArr[i] < timeArr[j]
	})
	defer func() {
		fmt.Println(timeArr, res)
	}()

	start := timeArr[0]
	end := timeArr[len(timeArr)-1]
	flag := false
	for k, v := range timeArr {
		// 第一个跟第二个日期不判断是否连续
		if k <= 1 {
			continue
		}
		// 最后一个跟最后第二个日期不判断是否连续
		if k == len(timeArr)-1 {
			continue
		}
		// 从第3个日期判断是否跟上一个日期是否连续
		yearA, _ := strconv.Atoi(v[0:4])
		monthA, _ := strconv.Atoi(v[5:7])
		// 上一个日期
		yearB, _ := strconv.Atoi(timeArr[k-1][0:4])
		monthB, _ := strconv.Atoi(timeArr[k-1][5:7])

		if yearA == yearB && monthB-monthA <= 1 {
			// 年份相同 月份不超过一个月 则连续
			flag = true
		} else if yearA != yearB && yearA-yearB == 1 && monthA == 1 && monthB == 12 {
			// 年份相差1年 12月要衔接到1月 则连续
			flag = true
		} else {
			flag = false
			break
		}
	}
	if !flag {
		return false
	}
	// 中间日期连续 判断是否满12个月
	sYear, _ := strconv.Atoi(start[0:4])
	sMonth, _ := strconv.Atoi(start[5:7])
	eYear, _ := strconv.Atoi(end[0:4])
	eMonth, _ := strconv.Atoi(end[5:7])
	if sYear == eYear && sMonth == 1 && eMonth == 12 {
		return true
	}
	if sYear != eYear {
		if eYear-sYear > 1 {
			// 连续且大于2年 一定满12个月
			return true
		}
		// 小于2年
		if eMonth < sMonth {
			// eMonth = 3 sMonth = 4
			if sMonth-eMonth <= 1 {
				return true
			}
		} else {
			//eMonth >= sMonth
			return true
		}
	}
	return false
}

func test1(m map[string]string) map[string]string {
	fmt.Printf("%p\n", m)
	if _, ok := m["a"]; ok {
		m["a"] = "b"
	}
	m["c"] = "d"
	m["e"] = "f"
	m["g"] = "h"
	m["i"] = "j"
	return m
}

func main() {
	// datas := NewDatas()
	// for _, v := range datas {
	// 	Judge(v)
	// }
	m := map[string]string{"a": "1"}
	fmt.Printf("%p\n", m)
	m1 := test1(m)
	fmt.Println(m1)
	fmt.Println(m)
}
