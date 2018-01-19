package main

import (
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "encoding/json"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Person struct{
  ID string `json:"id,omitempty" bson:"id,omitempty"`
  Firstname string `json:"firstname" bson:"firstname"`
  Lastname string `json:"lastname" bson:"lastname"`
  Age int `json:"age" bson:"age"`
}


func main(){
  // start session with mongo
  //session, err := mgo.Dial("mongodb://localhost:27017")
  //if err != nil {
  //  panic(err)
  //}
  //defer session.Close()

  // populate db with 2 records
  // c := session.DB("go_db").C("contacts")
  // err = c.Insert(&Person{"1","Ale", "Da",22},
  //               &Person{"2","Boba","Da", 53})
  // if err != nil {
  //   log.Fatal(err)
  // }


  router := mux.NewRouter()
  router.HandleFunc("/people",GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}",GetPerson).Methods("GET")
  router.HandleFunc("/people/{id}",CreatePerson).Methods("POST")
  router.HandleFunc("/people/{id}",DeletePerson).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000",router))
}



func GetPeople(w http.ResponseWriter, r *http.Request){

  // start session with mongo
  session, err := mgo.Dial("mongodb://localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  c := session.DB("go_db").C("contacts")

  // grab all results from collection, and store in result
  result := []Person{}
  err = c.Find(nil).All(&result)
  if err != nil {
    log.Fatal(err)
  }

  json.NewEncoder(w).Encode(result)
}


func GetPerson(w http.ResponseWriter, r *http.Request){
  // start mongo session
  session, err := mgo.Dial("mongodb://localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  c := session.DB("go_db").C("contacts")

  // grab id from request, and store as string
  params := mux.Vars(r)
  var id_string string = string(params["id"])

  // search mongo and store result in result variable
  result := Person{}
  err = c.Find(bson.M{"id": id_string}).One(&result)
  if err != nil {
    log.Fatal(err)
  }

  // encode and send back a json response
  json.NewEncoder(w).Encode(result)

}


func CreatePerson(w http.ResponseWriter, r *http.Request){
  // start mongo session
  session, err := mgo.Dial("mongodb://localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  c := session.DB("go_db").C("contacts")
  // assign decoded request to person variable and insert to collection
  params := mux.Vars(r)
  var person Person
  _ = json.NewDecoder(r.Body).Decode(&person)
  person.ID = params["id"]
  c.Insert(person)
  json.NewEncoder(w).Encode(person)

}



func DeletePerson(w http.ResponseWriter, r *http.Request){
  // start mongo session
  session, err := mgo.Dial("mongodb://localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  c := session.DB("go_db").C("contacts")

  // grab id from request, and store as string
  params := mux.Vars(r)
  var id_string string = string(params["id"])

  // Delete record
  err = c.Remove(bson.M{"id": id_string})
  if err != nil {
    log.Printf("remove fail %v\n", err)
  }
  json.NewEncoder(w).Encode("record deleted")

}

