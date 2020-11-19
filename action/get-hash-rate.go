package action

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	db_conn "telegram-aggregatio-data/db-conn"
	"telegram-aggregatio-data/shared-data"
)

type GetHashRate struct {
	locationName string
	listPage string
	hesh int
	heshAVG int
	typeAscList []shared_data.AntminerS9Data
	sync.Mutex
}




func (GHR *GetHashRate) listPageF(){
	GHRr := strconv.Itoa(GHR.hesh)
	GHRavg := strconv.Itoa(GHR.heshAVG)
	if len(GHRr) == 3{
		GHRr = GHRr[0:2]+"."+GHRr[2:]
		GHRavg = GHRavg[0:2]+"."+GHRavg[2:]
		GHR.listPage = fmt.Sprintf("*Hesh* %v_T/H\n HeshAVG* %v_T/H\n",GHRr,GHRavg)
	}

	if len(GHRr) == 4{
		GHRr = GHRr[0:3]+"."+GHRr[3:]
		GHRavg = GHRavg[0:3]+"."+GHRavg[3:]
		GHR.listPage = fmt.Sprintf("*Hesh* %v_T/H\n HeshAVG* %v_T/H\n",GHRr,GHRavg)
	}


}


func (GHR *GetHashRate) getHesh(keys interface{}) {
	keysLen := reflect.ValueOf(keys)
	for i := 0; i < keysLen.Len(); i++ {
		re := keysLen.Index(i).String()
		var d = shared_data.DataAsc{}
		var dd = shared_data.AntminerS9Data{}
		_=json.Unmarshal([]byte(re), &d)
		_=json.Unmarshal([]byte(d.DataAsc), &dd)
		GHR.typeAscList = append(GHR.typeAscList, dd)
	}
}


func (GHR *GetHashRate) getAllNameMisha()  {
	connClubacka := db_conn.DbConnKulebaka()
	allNameKu := connClubacka.DB2.GetData("misha")
	dataKu := connClubacka.DB1.GetData(allNameKu)
	GHR.getHesh(dataKu)

	connPila := db_conn.DbConnPila()
	pial := connPila.DB2.GetData("misha")
	dataPila := connPila.DB1.GetData(pial)
	GHR.getHesh(dataPila)

}

func (GHR *GetHashRate)ConcatHesh(GH,t string)  {
	nonFractionalPart := strings.Split(GH, ".")
	strH := nonFractionalPart[0]

	if len(strH) > 3 {
		//fmt.Println(len(strH),strH)
		var n int
		n, _ = strconv.Atoi(strH)
		if len(strH) == 5 {
			n, _ = strconv.Atoi(strH[0:3])
			//n = 0
		}
		if len(strH) == 4 {
			n, _ = strconv.Atoi(strH[0:2])
		}

		GHR.Lock()
		if t == "GHS"{
			GHR.hesh += n
		}else {
			GHR.heshAVG += n
		}

		GHR.Unlock()
	}

}

//func (GHR *GetHashRate) getAllVal(name string) {
	//dbConn := db_redis.DBConnector()
	//DB1 := dbConn.DB1
	//if res := DB1.GetKeyVal(name); len(res) !=0 {
	//	var dataAsc = &shared_data.DataAsc{}
	//
	//	st := res[0]
	//	dByte := []byte(st)
	//	err := json.Unmarshal(dByte,&dataAsc)
	//	//fmt.Println(dataAsc)
	//	if err != nil{
	//		fmt.Println(err,"GetHashRate m:getAllVal")
	//	}
	//
	//	PDS9 := &shared_data.AntminerS9Data{}
	//	b := []byte(dataAsc.DataAsc)
	//	err = json.Unmarshal(b,&PDS9)
	//	if err != nil{
	//		fmt.Println(err,"GetHashRate m:getAllVal")
	//	}

		//GHR.ConcatHesh(PDS9.GHS,"GHS")
		//GHR.ConcatHesh(PDS9.GHSAVG,"AVG")

	//}

//}

func (GHR *GetHashRate) mixData() {
	GHR.getAllNameMisha()

	for _,v := range GHR.typeAscList{
		GHR.ConcatHesh(v.GHS,"GHS")
		GHR.ConcatHesh(v.GHSAVG,"GHSAVG")
	}

}
func (GHR *GetHashRate) get()  {
	GHR.mixData()
	GHR.listPageF()
}

func GetHashRateAll(name string) string  {
	H := &GetHashRate{locationName:name}
	H.get()
	return H.listPage
}