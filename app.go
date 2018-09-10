package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	. "github.com/gonzaloescobar/prescriptions/config"
	. "github.com/gonzaloescobar/prescriptions/dao"
	. "github.com/gonzaloescobar/prescriptions/models"
	"github.com/gorilla/mux"
)

var config = Config{}
var dao = PrescriptionsDAO{}

	// Suma devuelve el resultado de la adición de dos números
func Suma(numero1, numero2 int) (resultado int) {
	resultado = numero1 + numero2
	return
}

// GET list of prescriptions
func AllPrescriptionsEndPoint(w http.ResponseWriter, r *http.Request) {
	prescriptions, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, prescriptions)
}

// GET a prescription by its ID
func FindPrescriptionEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	prescription, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Prescription ID")
		return
	}
	respondWithJson(w, http.StatusOK, prescription)
}

// POST a new prescription
func CreatePrescriptionEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var prescription Prescription
	if err := json.NewDecoder(r.Body).Decode(&prescription); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	prescription.ID = bson.NewObjectId()
	if err := dao.Insert(prescription); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, prescription)
}

// PUT update an existing prescription
func UpdatePrescriptionEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var prescription Prescription
	if err := json.NewDecoder(r.Body).Decode(&prescription); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(prescription); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing prescription
func DeletePrescriptionEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var prescription Prescription
	if err := json.NewDecoder(r.Body).Decode(&prescription); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(prescription); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/prescriptions", AllPrescriptionsEndPoint).Methods("GET")
	r.HandleFunc("/prescriptions", CreatePrescriptionEndPoint).Methods("POST")
	r.HandleFunc("/prescriptions", UpdatePrescriptionEndPoint).Methods("PUT")
	r.HandleFunc("/prescriptions", DeletePrescriptionEndPoint).Methods("DELETE")
	r.HandleFunc("/prescriptions/{id}", FindPrescriptionEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatal(err)
	}


}
