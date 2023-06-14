package diner

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestReadtop3Menu(t *testing.T) {

	data := `{
    
		"1":{
			"diner_id": 1,
			"foodmenu_id": {
				"f1": {
					"1": "Item A",
					"2": "Item B",
					"3": "Item C",
					"4": "Item D"
				}
			}
		},
		"2":{
		
			"diner_id": 2,
			"foodmenu_id": {
				"f2": {
					"2": "Item A",
					"10": "Item K",
					"3": "Item C"
				}
			}
		},
		"3":{
		
			"diner_id": 3,
			"foodmenu_id": {
				"f3": {
					"6": "Item F",
					"2": "Item B",
					"5": "Item E"
				}
			}
		},
		"4":{
		
			"diner_id": 3,
			"foodmenu_id": {
				"f3": {
					"6": "Item F",
					"2": "Item B",
					"5": "Item E"
				}
			}
		},
		"5":{
		
			"diner_id": 4,
			"foodmenu_id": {
				"f2": {
					"2": "Item B",
					"1": "Item A",
					"3": "Item C"
				}
			}
		},
		"6":{
		
			"diner_id": 5,
			"foodmenu_id": {
				"f2": {
					"2": "Item B",
					"4": "Item D",
					"3": "Item C"
				}
			}
		},
		"7":{
		
			"diner_id": 6,
			"foodmenu_id": {
				"f2": {
					"1": "Item A",
					"4": "Item D",
					"3": "Item C"
				}
			}
		}
	
	}`
	var datainter map[string]interface{}

	json.Unmarshal([]byte(data), &datainter)
	want := []Keyvalue{{"Item C", 5}, {"Item B", 5}, {"Item A", 4}}
	got := ReadTop3Menu(datainter)
	if reflect.DeepEqual(want, got) {
		t.Error("Not the expected output:", want, got)
	}
}
