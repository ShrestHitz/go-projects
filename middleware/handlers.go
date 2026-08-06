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

	json.NewEncoder(w).Encode(stock)
}

func GETAllStock(w http.ResponseWriter, r *http.Request){
	stocks, err := GETAllStock()

	if err!= nil {
		log.Fatal("Unable to get all the stocks %v", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	id, er := strconv.Atoi(params["id"])

	if err!= nil{
		log.Fatalf("Unable to convert the string int. %v", err)
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil{
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	updatedRows := UpdateStock(int64(id), )

	msg := fmt.Sprintf("Stock updated successfully. Total rows/records affected %v", updatedRows)
	res := response{
		ID: int64(id),
		Message: msg,
	}

	json.NewEnecoder(w).Encode(res)
}

func Deletestock(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string to int %v", err)
	}

	deletedRows := deletestock(int64(id))

	msg := fmt.Sprintf("Stock deleted successfully. Total rows/recods %v", deletedRows)

	res := response{
		ID : int64(id),
		Message: msg,
	}


	json.NewEncoder(w).Encode(res)
}

func insertStock() int64{

}

func getStock(id int64)(){

}

func GETAllStock()(){

}

func UpdateStock(id int64, stock models.Stock) int64{

}

func deletestock