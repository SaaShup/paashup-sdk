package docker

type ZoneAccount struct {
	Id                 int    `json:"id" yaml:"id"`
	Zone_Name          string `json:"zone_name" yaml:"zone_name"`
	Zone_Id            string `json:"zone_id" yaml:"zone_id"`
	Token              string `json:"token" yaml:"token"`
}

type DnsRecord struct {
	Id            int    `json:"id" yaml:"id"`
	Record_Id     string `json:"record_id" yaml:"record_id"`
	Name          string `json:"name" yaml:"name"`
	Zone          int    `json:"zone" yaml:"zone"`
	Type          string `json:"type" yaml:"type"`
	Content       string `json:"content" yaml:"content"`
	Ttl           string `json:"ttl" yaml:"ttl"`
	Proxied	      bool   `json:"proxied" yaml:"proxied"`
}

type DnsRecordComplete struct {
    DnsRecord
    ZoneAccount  ZoneAccount  `json:"accounts" yaml:"host"`
    Images []Image `json:"images" yaml:"images"`
}

type DnsRecordCreateStruct struct {
    Name          string `json:"name" yaml:"name"`
    ServerAddress string `json:"serveraddress" yaml:"serveraddress"`
    Username      string `json:"username" yaml:"username"`
    Password      string `json:"password" yaml:"password"`
    Email         string `json:"email" yaml:"email"`
    ZoneAccount   int    `json:"host" yaml:"host"`
}

type DnsRecordListStruct struct {
    Count    int              `json:"count" yaml:"count"`
    Next     string           `json:"next" yaml:"next"`
    Previous string           `json:"previous" yaml:"previous"`
    Results  []DnsRecordComplete `json:"results" yaml:"results"`
}

type ZoneAccountComplete struct {
	Token struct {
		Id            int    `json:"id" yaml:"id"`
		Url           string `json:"url" yaml:"url"`
		Display       string `json:"display" yaml:"display"`
		Key           string `json:"key" yaml:"key"`
		Write_enabled bool   `json:"write_enabled" yaml:"write_enabled"`
	}
	Netbox_base_url string      `json:"netbox_base_url" yaml:"netbox_base_url"`
	Custom_fields   CustomField `json:"custom_fields" yaml:"custom_fields"`
	Last_updated    string      `json:"last_updated" yaml:"last_updated"`
	Tags            []string    `json:"tags" yaml:"tags"`
	DnsRecords      []DnsRecord  `json:"registries" yaml:"registries"`
	ZoneAccount
}

type ZoneAccountListStruct struct {
	Count    int            `json:"count" yaml:"count"`
	Next     string         `json:"next" yaml:"next"`
	Previous string         `json:"previous" yaml:"previous"`
	Results  []ZoneAccountComplete `json:"results" yaml:"results"`
}

