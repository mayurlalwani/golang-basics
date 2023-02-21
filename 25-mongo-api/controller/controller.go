package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mayurlalwani/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func init(){
	 dotenv := goDotEnvVariable("MONGO_URI")
	clientOption := options.Client().ApplyURI(dotenv)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("MongoDb connection success!!")

	collection = (*mongo.Collection)(client.Database(dbName).Collection(collectionName))

	fmt.Println("Collection instance is ready!!")
}

func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with ID", inserted.InsertedID)
}

func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set":bson.M{"watched":true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Modified count:", result.ModifiedCount)
}

func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Deleted count:", result)
}

func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}},nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted: ",deleteResult)
	return deleteResult.DeletedCount
}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M
		err := cur.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}