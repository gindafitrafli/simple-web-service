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
	json.NewDecoder(r.Body).Decode(&outlet)
	fmt.Fprintf(w, "outlet named: %s, located at %s", outlet.Name, outlet.Location)
}