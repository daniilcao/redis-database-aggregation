package helper
//
//import (
//	db_redis "db-redis"
//	"encoding/json"
//	"fmt"
//	"os"
//	"path/filepath"
//	"sync"
//)
//
//type UserData struct {
//	Type string
//	Name string
//	Ip string
//}
//
//type SetDb0 struct {
//	db_redis.RedisDb0
//}
//
//func (sd *SetDb0) connDb(){
//
//}
//
//func (sd *SetDb0) readJson()  {
//	//filename is the path to the json config file
//	var abs string
//	abs, _ = filepath.Abs("./json-files/json-for-do0.json")
//	fmt.Println(abs)
//	file, err := os.Open(abs)
//	if err != nil {
//		panic("Error read  config file")
//	}
//	defer func() { _ = file.Close() }()
//	var MN []UserData
//	decoder := json.NewDecoder(file)
//	_ = decoder.Decode(&MN)
//
//	wg := sync.WaitGroup{}
//	for _,v := range MN{
//		wg.Add(1)
//		go func(v UserData) {
//			defer wg.Done()
//			sd.SetDataAsc(v.Name,v.Type,v.Name,v.Ip)
//		}(v)
//	}
//	wg.Wait()
//
//}
//
//func (sd *SetDb0) SetData() {
//	sd.connDb()
//	sd.readJson()
//	//sd.ConnCloe()
//}