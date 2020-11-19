package db_conn

import (
	"sync"
)

var (
	DBCk *dbConnectorKulebaka
	DBCp *dbConnectorPila
	DBCu *dbConnUseRegister
	onceKulebaka sync.Once
	oncePila sync.Once
	onceConnUseRegister sync.Once
)

type dbConnectorKulebaka struct {
	DB0
	DB1
	DB2
	DB5
}

type dbConnectorPila struct {
	DB0
	DB1
	DB2
	DB5
}

type dbConnUseRegister struct {
	DB0USReg

}

func DbConnUseRegister() *dbConnUseRegister{
	onceConnUseRegister .Do(func() {
		confDb := map[string]interface{}{
			"DbHost":      "-.-.-.-",
			"DbNum":      0,
			"DbPassword": "-",
			"DbPort":     "6002",
			"PoolSiz":    10,
		}

		DBCu = &dbConnUseRegister{}
		conn := NewConnectDbFactory(confDb)

		DBCu.DB0USReg = conn.CreateUseRegistryDB0()
		//DBCk.DB1 = conn.CreateDB1()
		//DBCk.DB2 = conn.CreateDB2()
		//DBCk.DB5 = conn.CreateDB5()
	})
	return DBCu
}

func DbConnKulebaka()  *dbConnectorKulebaka{
	onceKulebaka.Do(func() {
		confDb := map[string]interface{}{
			"DbHost":     "-.-.-.-",
			"DbNum":      0,
			"DbPassword": "-",
			"DbPort":     "6000",
			"PoolSiz":    10,
		}
		DBCk = &dbConnectorKulebaka{}
		conn := NewConnectDbFactory(confDb)

		DBCk.DB0 = conn.CreateDB0()
		DBCk.DB1 = conn.CreateDB1()
		DBCk.DB2 = conn.CreateDB2()
		DBCk.DB5 = conn.CreateDB5()
	})
	return DBCk
}


func DbConnPila()  *dbConnectorPila {
	oncePila.Do(func() {
		confDb := map[string]interface{}{
			"DbHost":      "-.-.-.-",
			"DbNum":      0,
			"DbPassword": "-",
			"DbPort":     "6001",
			"PoolSiz":    10,
		}
		DBCp = &dbConnectorPila{}
		conn := NewConnectDbFactory(confDb)

		DBCp.DB0 = conn.CreateDB0()
		DBCp.DB1 = conn.CreateDB1()
		DBCp.DB2 = conn.CreateDB2()
		DBCp.DB5 = conn.CreateDB5()
	})
	return DBCp
}
