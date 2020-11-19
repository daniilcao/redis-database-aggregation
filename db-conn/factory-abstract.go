package db_conn

type AbstractFactory interface {
	CreateDB0() DB0// Names are stored here asc
	CreateDB1() DB1 // This is stored data asc
	CreateDB2() DB2 // This is where users ' ASICs are stored
	CreateDB5() DB5 // This is stored status
	//------------------------------
	CreateUseRegistryDB0() DB0USReg
}


type ConnectDbFactory struct {
	configDb map[string]interface{}
}


func (f *ConnectDbFactory) CreateDB0() DB0 {
	conn := &Db0{ConfigDb: f.configDb}
	conn.Connect()
	return conn
}

func (f *ConnectDbFactory) CreateDB1() DB1 {
	conn := &Db1{ConfigDb: f.configDb}
	conn.Connect()
	return conn
}

func (f *ConnectDbFactory) CreateDB2() DB2 {
	conn := &Db2{ConfigDb: f.configDb}
	conn.Connect()
	return conn
}

func (f *ConnectDbFactory) CreateDB5() DB5 {
	conn := &Db5{ConfigDb: f.configDb}
	conn.Connect()
	return conn
}

//----------------------User-----------------
func (f *ConnectDbFactory) CreateUseRegistryDB0() DB0USReg {
	conn := &Db0usReg{ConfigDb: f.configDb}
	conn.Connect()
	return conn
}
func NewConnectDbFactory(conf map[string]interface{}) *ConnectDbFactory {
	return &ConnectDbFactory{configDb:conf}
}
