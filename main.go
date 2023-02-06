package main
 
import(
	"github.com/tsunagatteru/ishiki-no-nagare/server"
	"github.com/tsunagatteru/ishiki-no-nagare/db"
	_ "github.com/tsunagatteru/ishiki-no-nagare/model"
	"github.com/tsunagatteru/ishiki-no-nagare/config"
)


func main(){
	config := config.Read()
	dbConn := db.Open()
	defer dbConn.Close()
	db.CreateTable(dbConn)
	router := server.NewRouter()
	server.RunRouter(router, dbConn, config)
}
