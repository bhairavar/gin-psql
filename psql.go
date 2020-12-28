package gin_psql



import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //import postgres
	"log"
	"os"
)

// DB ...

type DB struct {
	*sql.DB
}

var db *gorp.DbMap


// Initialize ...

func Init(user string, password string, host string, name string, SslMode string, Trace bool ) (status int){

	//DbUser ...
	var DbUser = user
	//DbPassword ...
	var DbPassword = password
	//DbHost
	var DbHost = host
	//DbName ...
	var DbName =  name
	//Db Ssl Mode
	var DbSSLMode = SslMode


	DbInfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s",
		DbUser, DbPassword, DbHost, DbName, DbSSLMode)

	var err error
	db, err, status = ConnectDB(DbInfo, Trace)
	if err != nil {
		fmt.Println(err)
		return status
	}
	return status

}


//Connect DB ...
func ConnectDB(dataSourceName string, Trace bool) ( gp *gorp.DbMap, err error, status int) {
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		return nil, err, 1
	}
	if err = db.Ping(); err != nil {
		return nil, err, 1
	}
	trace := Trace
	DbMap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	if trace == true{
		DbMap.TraceOn("[gorp]", log.New(os.Stdout, "golang[gin]:", log.Lmicroseconds)) //Trace database requests
	}

	//
	return DbMap, nil, 0
}

//Get DB ...
func GetDB() *gorp.DbMap {
	return db
}

