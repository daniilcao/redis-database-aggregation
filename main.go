package main

import (
	db_conn "telegram-aggregatio-data/db-conn"
)


func init()  {
	db_conn.DbConnKulebaka()
	db_conn.DbConnPila()
	db_conn.DbConnUseRegister()
}



func main() {
	dbUser := db_conn.DbConnUseRegister()
	param := dbUser.GetUser("343q353")

	if len(param) == 0{
		var data = map[string]interface{}{}
		data["name"]= "Misha"
		data["id"] = "343q353"
		dbUser.SetData("343q353",data)
	}

	//res := action.GetHashRateAll("Нижний")
	//action.GetHesUser("Пиларама")

	//action.GetHesUser("Нижний")
	//action.GetHesUser("Пиларама")

}
