package main

import (
	"mhilmi999/project-2-mhilmi999/db"
	"mhilmi999/project-2-mhilmi999/routes"
)



func main(){
	db.Init()
	
	e := routes.Init() 

	e.Logger.Fatal(e.Start(":1234")) // Running proses web server di port http://localhost:1234
}