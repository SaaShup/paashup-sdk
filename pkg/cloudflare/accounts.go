package cloudflare

import (
    "github.com/SaaShup/paashup-sdk/pkg/netbox"
    "encoding/json"
    "fmt"
)

func AccountList() (ZoneAccountListStruct, error){
    var accountList ZoneAccountListStruct
    result, err := netbox.Request("/dns/accounts/", "GET", nil)

    if err != nil {
        return accountList, err
    }
    if err := json.Unmarshal(result, &accountList); err != nil { // Parse []byte to the go struct pointer
        return ZoneAccountListStruct{}, err
    }
    return accountList, nil
}

func AccountSearchByName(name string) (ZoneAccountComplete, error){
	if name == "" {
		return ZoneAccountComplete{}, fmt.Errorf("Account not found")
	}

	url := fmt.Sprintf("/dns/accounts/?name=%s", name)
	var result ZoneAccountListStruct
	resultCall, err := netbox.Request(url, "GET", nil)

	if err != nil {
		return ZoneAccountComplete{}, err
	}

	if err := json.Unmarshal(resultCall, &result); err != nil { // Parse []byte to the go struct pointer
        return ZoneAccountComplete{}, err
	}

	if result.Count == 1 {
		return result.Results[0], nil
	} else {
		return ZoneAccountComplete{}, fmt.Errorf("Account not found")
	}
}
