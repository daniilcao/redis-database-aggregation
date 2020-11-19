//This is collecting information about asic hash
package action

import (
	"encoding/json"
	"fmt"
	"reflect"
	shared_data "telegram-aggregatio-data/shared-data"
	db_conn "telegram-aggregatio-data/db-conn"
)

type GetSumHesh struct {
	Name string
	typeAscList []shared_data.DataAsc
	strResDataS9 string
	strResDataL3Plus string
}

func (GsH *GetSumHesh) getTypeAsc(keys interface{}) {

	keysLen := reflect.ValueOf(keys)
	for i := 0; i < keysLen.Len(); i++ {
		re := keysLen.Index(i).String()
		var d = shared_data.DataAsc{}
		_=json.Unmarshal([]byte(re), &d)
		GsH.typeAscList = append(GsH.typeAscList, d)
	}
}

func (GsH *GetSumHesh) GetHashDirty() {

	switch GsH.Name{
	case "Нижний":
		dbRKulebaka := db_conn.DbConnKulebaka()
		user := dbRKulebaka.DB0.GetData(nil)
		data := dbRKulebaka.DB1.GetData(user)
		GsH.getTypeAsc(data)
		GsH.makeData()
	case "Новинки":
		fmt.Println("")
	case "Пиларама":
		dbRPila:= db_conn.DbConnPila()
		user := dbRPila.DB0.GetData(nil)
		data := dbRPila.DB1.GetData(user)
		GsH.getTypeAsc(data)
		GsH.makeData()
	}

}

func (GsH *GetSumHesh) makeData(){
	for _,v := range GsH.typeAscList{
		switch v.TypeAsc {
		case "AntminerS9":
			data := shared_data.AntminerS9Data{}
			_ = json.Unmarshal([]byte(v.DataAsc),&data)
			GsH.strResDataS9 += fmt.Sprintf("**Type %v******\n*Name* %v\n HeshAVG* %v_T/H\n",v.TypeAsc,v.Name,data.GHSAVG)
		case "AntminerS9j":
			data := shared_data.AntminerS9Data{}
			_ = json.Unmarshal([]byte(v.DataAsc),&data)
			GsH.strResDataS9 += fmt.Sprintf("**Type %v******\n*Name* %v\n HeshAVG* %v_T/H\n",v.TypeAsc,v.Name,data.GHSAVG)
		case "AntminerS9i":
			data := shared_data.AntminerS9Data{}
			_ = json.Unmarshal([]byte(v.DataAsc),&data)
			GsH.strResDataS9 += fmt.Sprintf("**Type %v******\n*Name* %v\n HeshAVG* %v_T/H\n",v.TypeAsc,v.Name,data.GHSAVG)
		case "AntminerL3+":
			data := shared_data.AntminerL3PlusData{}
			_ = json.Unmarshal([]byte(v.DataAsc),&data)
			GsH.strResDataL3Plus += fmt.Sprintf("*Type %v******\n*Name* %v\n HeshAVG* %v_T/H\n",v.TypeAsc,v.Name,data.GHSAVG)
		default:
			fmt.Println("Non GetSumHesh")
		}
	}
}

func  GetHesUser(name string)  {
	gsh := GetSumHesh{Name: name}
	gsh.GetHashDirty()
	var resData string
	if len(gsh.strResDataS9) != 0{resData += gsh.strResDataS9}
	if len(gsh.strResDataL3Plus) != 0{resData += gsh.strResDataL3Plus}
	fmt.Println(resData)
}

func GetHesAllData(name string)  {

}