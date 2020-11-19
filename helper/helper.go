package helper

import (
	"math"
	"reflect"
	"strconv"
	"time"
)

//Делает лист из map
func MakeListInMap(m map[string]string) ([]string, []string) {
	var param0 []string
	var param1 []string
	for val0, val1 := range m {
		param0 = append(param0, val0)
		param1 = append(param1, val1)
	}
	return param0, param1
}

func GetTime() string {
	dt := time.Now().UTC()
	//dt = time.Date(
	//	dt.Year(),
	//	dt.Month(),
	//	dt.Day(),
	//	dt.Hour(),
	//	dt.Minute(),
	//	dt.Second(),
	//	dt.Nanosecond(),
	//	time.UTC)
	return dt.String()
}

func DiffTest(t string) bool {
	dt := time.Now()
	layout := "2006-01-02 15:04:05 -0700 MST"
	v, _ := time.Parse(layout, t)

	n := strconv.Itoa(dt.Minute())
	asc := strconv.Itoa(v.Minute())


	if  v.Minute() <= dt.Minute() {
		return true
	}else if len(asc) != len(n) {
		return true
	}
	return false
}

func NextTime(m int) string {
	dt := time.Now()
	dt = time.Date(
		dt.Year(),
		dt.Month(),
		dt.Day(),
		dt.Hour(),
		dt.Minute()+m,
		dt.Second(),
		dt.Nanosecond(),
		time.UTC)
	return dt.String()
}

func MakeMapInMap(m map[string]map[string]string, n int) []map[string]map[string]string {

	var elementMap []map[string]map[string]string
	breakingNum := math.Ceil(float64(len(m)) / float64(n))
	iter := 0
	i := 0
	for name, data := range m {

		if i%int(breakingNum) == 0 {
			m := map[string]map[string]string{}
			elementMap = append(elementMap, m)
			iter++
		}
		i++

		if len(elementMap) != 0 {
			elementMap[iter-1][name] = data
		}
	}

	return elementMap
}


func MakeInListInList(m []string, n int) [][]string  {
	var elementMap [][]string
	breakingNum := math.Ceil(float64(len(m)) / float64(n))

	iter := 0
	i := 0

	for _,name := range m {

		if i%int(breakingNum) == 0 {
			var m  []string
			elementMap = append(elementMap, m)
			iter++
			i = 0
		}
		i++
		if len(elementMap) != 0 {
			elementMap[iter-1] = append(elementMap[iter-1],name)
		}

	}

	return elementMap
}

func MakeInMap(m map[string]string, n int) []map[string]string {
	var elementMap []map[string]string
	breakingNum := math.Ceil(float64(len(m)) / float64(n))
	iter := 0
	i := 0
	for name, ip := range m {
		if i%int(breakingNum) == 0 {
			m := make(map[string]string)
			elementMap = append(elementMap, m)
			iter++
		}
		i++
		if len(elementMap) != 0 {
			elementMap[iter-1][name] = ip
		}
	}

	return elementMap
}

type PairKeyValue struct {
	Id    string
	ID string
	data map[string]interface{}

}

func InterFaceMapConvertMap(data reflect.Value) map[string]string  {
	mapRes := make(map[string]string)
	for _,key := range data.MapKeys(){
		p := data.MapIndex(key)
		r:= p.Interface()
		mapRes[key.String()] = r.(string)
	}
	return mapRes
}