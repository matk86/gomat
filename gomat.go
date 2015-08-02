package main

import (
       "encoding/json"
       "fmt"
       "io/ioutil"
       "net/http"
       "os"
       )

func main() {
	preamble := "https://www.materialsproject.org/rest/v1"
	request_type := "materials" //"materials", "battery", "reaction", "mpquery" and "api_check"
	identifier := os.Args[1] 
	data_type := "vasp"
	mapi_key := os.Getenv("MAPI_KEY")
	url := fmt.Sprintf(
		"%s/%s/%s/%s?API_KEY=%s",
		preamble,
		request_type,
		identifier,
		data_type,
		mapi_key )
	client := &http.Client{}
	req, httperr := http.NewRequest("GET", url, nil)
	if httperr != nil {
		fmt.Errorf("gomat: %s", httperr)
	}
	resp, reqerr := client.Do(req)
	if reqerr != nil {
		fmt.Errorf("gomat: %s", reqerr)
	}
	defer resp.Body.Close()
	body, dataerr := ioutil.ReadAll(resp.Body)
	if dataerr != nil {
		fmt.Errorf("gomat: %s", dataerr)
	}
	var output interface{}
	jsonerr := json.Unmarshal(body, &output)
	if jsonerr != nil {
		fmt.Errorf("gomat: %s", jsonerr)
	}
	m := output.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, " :", vv)
		case int:
			fmt.Println(k, " :", vv)
		case []interface{}:
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "unknown type")
		}
	}
}
