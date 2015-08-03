package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func RecursiveDataProcess(d map[string]interface{}) {
	for k, v := range d {
		switch dd := v.(type) {
		case []interface{}:
			for _, u := range dd {
				uu, _ := u.(map[string]interface{})
				RecursiveDataProcess(uu)
			}
		case map[string]interface{}:
			fmt.Println(k, " :")
			for l, m := range dd {
				fmt.Println("    ", l, " :", m)
			}
		default:
			fmt.Println(k, " :", dd)
		}
	}

}

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
		mapi_key)
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
	//fmt.Printf(string(body))
	if dataerr != nil {
		fmt.Errorf("gomat: %s", dataerr)
	}
	var data interface{}
	jsonerr := json.Unmarshal(body, &data)
	if jsonerr != nil {
		fmt.Errorf("gomat: %s", jsonerr)
	}
	d := data.(map[string]interface{})
	RecursiveDataProcess(d)
	fmt.Println("Done")
}
