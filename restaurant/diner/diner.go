package diner

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
)

var duprec []interface{}

func ReadData() map[string]interface{} {

	var filedata map[string]interface{}

	filepath := "d:/programtest/restaurant/diner1.json"

	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error 1: ", err)
	}

	err = json.Unmarshal(data, &filedata)
	if err != nil {
		fmt.Println("Error 2: ", err)
	}
	return filedata
}

func FindDuplicate(data map[string]interface{}) {

	for _, value := range data {

		diner_id := value.(map[string]interface{})["diner_id"]
		foodmenu_id := value.(map[string]interface{})["foodmenu_id"]

		CheckDuplicate(diner_id, foodmenu_id, data)

	}

}

func CheckDuplicate(id interface{}, fid interface{}, data map[string]interface{}) {

	i := 0
	for _, value := range data {

		diner_id := value.(map[string]interface{})["diner_id"]

		if id == diner_id {
			i = i + 1
			foodmenu_id := value.(map[string]interface{})["foodmenu_id"]
			out := reflect.DeepEqual(fid, foodmenu_id)

			if out && i > 1 {

				if duprec != nil {

					for val := range duprec {
						if reflect.DeepEqual(val, id) {
							return
						} else {
							duprec = append(duprec, id)
							fmt.Println("Found duplicate record for the diner_id: ", id)
							return
						}
					}
				} else {
					duprec = append(duprec, id)
				}

			}
		}

	}
}

func ReadTop3Menu(data map[string]interface{}) []Keyvalue {

	count := make(map[string]int)

	for _, value := range data {

		foodmenu_id := value.(map[string]interface{})["foodmenu_id"]

		for _, val := range foodmenu_id.(map[string]interface{}) {

			food_item := val.(map[string]interface{})

			for _, val := range food_item {
				if count == nil {

					count[val.(string)] = 1

				} else {
					found := 0
					for item := range count {
						if item == val {
							count[item] = count[item] + 1
							found = 1
						}
					}
					if found == 0 {
						count[val.(string)] = 1
					}
				}

			}
		}

	}
	var ss []Keyvalue
	var output []Keyvalue
	for k, v := range count {
		ss = append(ss, Keyvalue{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	fmt.Println("Top selling 3 Item")
	for i := 0; i < 3; i++ {
		fmt.Println(ss[i])
		output = append(output, ss[i])
	}
	return output
}

type Keyvalue struct {
	Key   string
	Value int
}
