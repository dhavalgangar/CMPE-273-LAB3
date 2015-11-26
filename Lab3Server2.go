package main

import (
	"encoding/json"
	"fmt"
	"httprouter"
	"net/http"
	"strconv"
)

//MyController structure
type MyController struct{}

//NewMyController function
func NewMyController() *MyController {
	return &MyController{}
}

//DataCache struct
type DataCache struct {
	DataCachekey   int    `json:"key"`
	DataCachevalue string `json:"value"`
}

var myMap = make(map[int]string)

func main() {
	// Instantiate a new router
	router := httprouter.New()

	// Get a controller instance
	mycontroller := NewMyController()

	// Add handlers
	router.GET("/keys/:id", mycontroller.GetKey)
	router.GET("/keys", mycontroller.GetAllKeys)
	router.PUT("/keys/:id/:value", mycontroller.PutKey)

	// Expose the server at port 3000
	http.ListenAndServe(":3001", router)
}

//GetKey to display single key value
func (uc MyController) GetKey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	eid, _ := strconv.Atoi(id)

	var key int
	var val string

	fmt.Println("in GET", myMap)
	for key, val = range myMap {
		if key == eid {
			val = myMap[key]
			break
		}
	}
	uj := DataCache{
		DataCachekey:   key,
		DataCachevalue: val,
	}

	result, _ := json.Marshal(uj)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", result)
}

//GetAllKeys to display all the key values
func (uc MyController) GetAllKeys(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	type data []DataCache
	var dataArray = make(data, 0)

	for key, val := range myMap {
		uj := DataCache{key, val}
		dataArray = append(dataArray, uj)
	}

	result, _ := json.Marshal(dataArray)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", result)

}

//PutKey to insert value in a map
func (uc MyController) PutKey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	val := p.ByName("value")
	eid, _ := strconv.Atoi(id)

	myMap[eid] = val
	fmt.Println(myMap)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
