//Keywords: json golang tree walk

package main

import (
	"encoding/json"
)

var a = []byte(`
{"x":2.5,
"y":	 {
	"z1":"aréác",
	"z2":true	
},
"z":3,
"r":[null,12,"hello"]}`)

type Map map[string]json.RawMessage
type Array []json.RawMessage

func walkJson(raw json.RawMessage) {
	if raw[0] == 123 { //  123 is `{` => object
		var cont Map
		json.Unmarshal(raw, &cont)
		for i, v := range cont {
			println(i, ":")
			walkJson(v)
		}
	} else if raw[0] == 91 { // 91 is `[`  => array
		var cont Array
		json.Unmarshal(raw, &cont)
		for i, v := range cont {
			println(i, ":")
			walkJson(v)
		}

	} else {
		var val interface{}
		json.Unmarshal(raw, &val)
		switch v := val.(type) {
		case float64:
			println("float:", v)
		case string:
			println("string:", v)
		case bool:
			println("bool:", v)
		case nil:
			println("nil")
		default:
			println("unkown type")
		}
	}
}

func main() {
	var tmpJ json.RawMessage

	json.Unmarshal(a, &tmpJ) //compacting it
	walkJson(tmpJ)

}
