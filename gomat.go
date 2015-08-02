package main

import (
       "encoding/json"
       "fmt"
       "io/ioutil"
       "net/http"
       "os"
       )

func main() {
	url := fmt.Sprintf("https://www.materialsproject.org/rest/v1/materials/%s/vasp?API_KEY=%s",
		os.Args[1],
		os.Getenv("MAPI_KEY") )
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("gomat: %s", err)
	}
	resp, requestErr := client.Do(req)
	if requestErr != nil {
		fmt.Errorf("gomat: %s", requestErr)
	}
	defer resp.Body.Close()
	body, dataReadErr := ioutil.ReadAll(resp.Body)
	if dataReadErr != nil {
		fmt.Errorf("gomat: %s", dataReadErr)
	}
	var output interface{}
	err = json.Unmarshal(body, &output)
	m := output.(map[string]interface{})
	if err != nil {
		fmt.Errorf("gomat: %s", err)
	}
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
