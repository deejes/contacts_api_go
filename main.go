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
  ID string `json:"id,omitempty"`
  Firstname string `json:"firstname,omitempty"`
  Lastname string `json:"lastname,omitempty"`
  Age int `json:"age,omitempty"`
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
  //  router.HandleFunc("/people/{id}",CreatePerson).Methods("POST")
  //  router.HandleFunc("/people/{id}",DeletePerson).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000",router))
}



func GetPeople(w http.ResponseWriter, r *http.Request){

  // start session with mongo
  session, err := mgo.Dial("mongodb://localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  // populate db with 2 records
  c := session.DB("go_db").C("contacts")
  err = c.Insert(&Person{"1","Ale", "Da",22},
  &Person{"2","Boba","Da", 53})
  if err != nil {
    log.Fatal(err)
  }

  // c:= session.DB("go-db").C("contacts")
  result := Person{}
  err = c.Find(bson.M{"id": "1"}).One(&result)
  //err = c.Find(nil).One(&result)
  if err != nil {
    log.Fatal(err)
  }

  log.Println(result)
  json.NewEncoder(w).Encode(result)
}
func GetPerson(w http.ResponseWriter, r *http.Request){
  session, err := mgo.Dial("mongodb://localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  params := mux.Vars(r)
  var id_string string = string(params["id"])
  log.Println(id_string)
  c := session.DB("go_db").C("contacts")

  result := Person{}
  err = c.Find(bson.M{"id": id_string}).One(&result)
  //err = c.Find(nil).One(&result)
  if err != nil {
    log.Fatal(err)
  }
  log.Println(result)
  json.NewEncoder(w).Encode(result)

  //      for _, item := range people {
  //      if item.ID == params["id"]{json.NewEncoder(w).Encode(item)}}
}
//func CreatePerson(w http.ResponseWriter, r *http.Request){
//    params := mux.Vars(r)
//    var person Person
//    _ = json.NewDecoder(r.Body).Decode(&person)
//    person.ID = params["id"]
//    people = append(people, person)
//    json.NewEncoder(w).Encode(people)
//     }
//
//func DeletePerson(w http.ResponseWriter, r *http.Request){
//    params := mux.Vars(r)
//    for index,item := range people{
//        if item.ID == params["id"] {
//          people = append(people[:index], people[index+1:]...)
//          break
//        }
//    }
//}
//
//var people string = "a"
