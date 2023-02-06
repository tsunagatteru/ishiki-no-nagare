package main
 
import(
	"github.com/tsunagatteru/ishiki-no-nagare/server"
	"github.com/tsunagatteru/ishiki-no-nagare/db"
)


func main(){
	dbConn := db.Open()
	defer dbConn.Close()
	db.CreateTable(dbConn)
	router := server.NewRouter()
	server.RunRouter(router, dbConn)
}
