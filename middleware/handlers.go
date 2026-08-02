package middleware 

import(
	"database/sql"
	"github.com/joho/godotenv"
	"fmt"
	"log"
	"models"
	"encoding/json"
	"net/http"
	"os"
	"go-postgres-yt/models"
	"github.com/gorilla/mux"
)

type response struct(
	ID int64       'josn:"id,omitempty"'
	Message string 'json:"message,omitempty"'
)

func CreateConnection() *sql.DB{
	err := godotenv.Load(".env")

	if ee !=nil{
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil{
		panic(err)
	}

	err = db.Ping()

	if err!= nil{
		panic(err)
	}

	fmt.Println("Successfully connected to postgres")
	return db
}


//functions creation

func CreateStock(w http.ResponseWriter, r *http.Request){
	var stock models.stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil{
		log.Fatal("Unable to decode the request body. %v", err)
	}

	insertID := insertStock(stock)

	res := response{
		ID : insertID,
		Message: "stock created successfully"
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	id, err := trconv.Atoi(params["id"])

	if err != nil{
		log.Fatal("Unable to convert the string into int")
	}
}

func GETAllStock(){

}

func UpdateStock(){

}

func Deletestock(){

}