package cloudflare

import (
    "github.com/SaaShup/paashup-sdk/netbox"
    "encoding/json"
    "fmt"
)

func DnsRecordList() (DnsRecordListStruct error){
    var DnsRecordList DnsRecordListStruct
    result, err := netbox.Request("/dns/records/", "GET", nil)

    if err != nil {
        return DnsRecordList, err
    }
    if err := json.Unmarshal(result, &DnsRecordList); err != nil { // Parse []byte to the go struct pointer
        return DnsRecordList, err
    }
    return DnsRecordList, nil
}

func DnsRecordListByAccount(hostId int) (DnsRecordListStruct error){
    var DnsRecordList DnsRecordListStruct
    result, err := netbox.Request(fmt.Sprintf("/dns/records/?account_id=%d", hostId), "GET", nil)

    if err != nil {
        return DnsRecordList, err
    }
    if err := json.Unmarshal(result, &DnsRecordList); err != nil {
        return DnsRecordList, err
    }
    return DnsRecordList, nil
}

func DnsRecordInspect(DnsRecordId int) (DnsRecordComplete, error){
    var DnsRecord DnsRecordComplete
    url := fmt.Sprintf("/dns/records/%d/", DnsRecordId)
    result, err := netbox.Request(url, "GET", nil)

    if err != nil {
        return DnsRecord, err
    }
    if err := json.Unmarshal(result, &DnsRecord); err != nil {
        return DnsRecord, err
    }
    return DnsRecord, nil
}

func DnsRecordSearchByName(name string, accountId int) (DnsRecordComplete, error){
	if name == "" {
		return DnsRecordComplete{}, fmt.Errorf("Record not found")
	}

	url := fmt.Sprintf("/dns/record/?name=%s&account_id=%d", name, accountId)
	var result DnsRecordListStruct
	resultCall, err := netbox.Request(url, "GET", nil)

	if err != nil {
		return DnsRecordComplete{}, err
	}

	if err := json.Unmarshal(resultCall, &result); err != nil {
		return DnsRecordComplete{}, err
	}

	if result.Count == 1 {
		return result.Results[0], nil
	} else {
		return DnsRecordComplete{}, fmt.Errorf("Record not found")
	}
}

func DnsRecordCreate(DnsRecord DnsRecordCreateStruct) (DnsRecordComplete, error){
    var DnsRecordResponse DnsRecordComplete
    jsonStr, _ := json.Marshal(DnsRecord)
    result, err := netbox.Request("/dns/records/", "POST", jsonStr)

    if err != nil {
        return DnsRecordResponse, err
    }
    if err := json.Unmarshal(result, &DnsRecordResponse); err != nil {
        return DnsRecordResponse, err
    }
    return DnsRecordResponse, nil
}

func DnsRecordDelete(dnsrecordId int) error{
    url := fmt.Sprintf("/dns/records/%d/", dnsrecordId)
    _, err := netbox.Request(url, "DELETE", nil)

    if err != nil {
        return err
    }
    return nil
}
