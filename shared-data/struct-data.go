package shared_data

type DataAsc struct {
	PingStatus string
	Ip         string
	DataAsc   string
	Name string
	TypeAsc string
	TimeQuery string
}

type AntminerS9Data struct {
	GHS string `json:"ghs"`
	GHSAVG string `json:"ghsavg"`
	ChainAcs5 string `json:"chain_acs_5"`
	ChainAcs6 string `json:"chain_acs_6"`
	ChainAcs7 string `json:"chain_acs_7"`
	ChainAcn6 string `json:"chain_acn_6"`
	ChainAcn7 string `json:"chain_acn_7"`
	ChainAcn8 string `json:"chain_acn_8"`
	ChainRateideal6 string `json:"chain_rateideal_6"`
	ChainRateideal7 string `json:"chain_rateideal_7"`
	ChainRateideal8 string `json:"chain_rateideal_8"`
	ChainRate6 string `json:"chain_rate_6"`
	ChainRate7 string `json:"chain_rate_7"`
	ChainRate8 string `json:"chain_rate_8"`
	Temp26 string `json:"temp_2_6"`
	Temp27 string `json:"temp_2_7"`
	Temp28 string `json:"temp_2_8"`
	Temp6 string `json:"temp_6"`
	Temp7 string `json:"temp_7"`
	Temp8 string `json:"temp_8"`
	Fan5 string `json:"fan_5"`
	Fan6 string `json:"fan_6"`
	Freqavg6 string `json:"freqavg_6"`
	Freqavg7 string `json:"freqavg_7"`
	Freqavg8 string `json:"freqavg_8"`
	Worker1 string `json:"worker_1"`
	Worker2 string `json:"worker_2"`
	Worker3 string `json:"worker_3"`
	User1 string `json:"user_1"`
	User2 string `json:"user_2"`
	User3 string `json:"user_3"`
}

type AntminerL3PlusData struct {
	GHS string `json:"ghs"`
	GHSAVG string `json:"ghsavg"`
	ChainAcs1 string `json:"chain_acs1"`
	ChainAcs2 string `json:"chain_acs2"`
	ChainAcs3 string `json:"chain_acs3"`
	ChainAcs4 string `json:"chain_acs4"`
	ChainAcn1 string `json:"chain_acn1"`
	ChainAcn2 string `json:"chain_acn2"`
	ChainAcn3 string `json:"chain_acn3"`
	ChainAcn4 string `json:"chain_acn4"`
	ChainRate1 string `json:"chain_rate1"`
	ChainRate2 string `json:"chain_rate2"`
	ChainRate3 string `json:"chain_rate3"`
	ChainRate4 string `json:"chain_rate4"`
	Temp1 string `json:"temp1"`
	Temp2 string `json:"temp2"`
	Temp3 string `json:"temp3"`
	Temp4 string `json:"temp4"`
	Fan1 string `json:"fan1"`
	Fan2 string `json:"fan2"`
	Frequency string `json:"frequency"`
	Worker1 string `json:"worker_1"`
	Worker2 string `json:"worker_2"`
	Worker3 string `json:"worker_3"`
	User1 string `json:"user_1"`
	User2 string `json:"user_2"`
	User3 string `json:"user_3"`
}
