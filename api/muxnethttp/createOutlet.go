package muxnethttp

import(
	"net/http"
	"encoding/json"
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type Outlet struct{
	Name string `json:"outletName"`
	Location string `json:"location"`
}

type CreatedOutlet struct{
	Id int64 `json:"id"`
}


func CreateOutlet(w http.ResponseWriter, r *http.Request) {
	var outlet Outlet
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&outlet)
	// json.NewEncoder(w).Encode(outlet)//using object
	/*  Using JSON String  */
	// w.Write([]byte(fmt.Sprintf(`{"name": "%s", "location":"%s"}`, outlet.Name, outlet.Location)))//using json string
	// fmt.Printf("outlet named: %s, located at %s", outlet.Name, outlet.Location)

	cfg := mysql.Config{
		User: "root",
		Passwd: "mysql",
		Net: "tcp",
		Addr: "127.0.0.1:5432",
		DBName: "outlet_main",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
			log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
			log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	result, err:= db.Exec("insert into outlet (name, location) values (?, ?)", outlet.Name, outlet.Location)

	if err != nil {
		w.WriteHeader(500)
		log.Panic("add outlet: %v", err)
		return
	}

	createdId, err := result.LastInsertId()
	if err != nil {
		w.WriteHeader(500)
		log.Panic("add outlet: %v", err)
		return
	}

	outletId := CreatedOutlet{
		Id: createdId,
	}

	json.NewEncoder(w).Encode(outletId)
	
	
}