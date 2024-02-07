package muxnethttp

import(
	"net/http"
	"encoding/json"
	"fmt"
)

type Outlet struct{
	Name string `json:"outletName"`
	Location string `json:"location"`
}

func CreateOutlet(w http.ResponseWriter, r *http.Request) {
	var outlet Outlet
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&outlet)
	// json.NewEncoder(w).Encode(outlet)//using object
	/*  Using JSON String  */

	w.Write([]byte(fmt.Sprintf(`{"name": "%s", "location":"%s"}`, outlet.Name, outlet.Location)))//using json string
	
	
	
	fmt.Printf("outlet named: %s, located at %s", outlet.Name, outlet.Location)
}