package handlers

import (
	"encoding/json"
	"test-d-2/connection"
	"test-d-2/structs"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var(
	stock 	 float32
	bond	 float32
	mm		 float32
)

func CountRisk(age int){
	var sum = 55 - age

	if sum >= 30 {
		stock = 72.5
		bond= 21.5
		mm = 100 - (stock + bond)
		return
	}else if sum >=20 {
		stock = 54.5
		bond = 25.5
		mm = 100 - (stock + bond)
		return
	}else if sum < 20{
		stock = 34.5
		bond = 45.5
		mm = 100 - (stock + bond)
		return
	} 
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


func CreateUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	
	var user structs.User

	json.Unmarshal(payloads, &user)
	
	hash, _ := HashPassword(user.Password)
	user.Password = hash
	
	CountRisk(user.AGE)

	connection.DB.Create(&user)

	risk_profile := structs.RiskProfile{UserID :user.UserID, MM:mm, Bond:bond, Stock:stock}
	connection.DB.Create(risk_profile)

	connection.DB.Preload("RiskProfiles").
	First(&user, user.UserID)
	res := structs.Result{Code: 200, Data: user, Message: "Create User Success"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func LoginUser(w http.ResponseWriter, r *http.Request){
	payloads, _ := ioutil.ReadAll(r.Body)
	var user structs.User
	var userPayload structs.User

	json.Unmarshal(payloads, &userPayload)
	connection.DB.First(&user, userPayload.UserID)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userPayload.Password))	
	
	if err != nil {
		res := structs.Result{Code: 400, Data: "", Message: "Wrong userid/password"}
		result, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(result)
		return
	}

	res := structs.Result{Code: 200, Data: user.NAME, Message: "Login Success"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetAllUser(w http.ResponseWriter, r *http.Request){

	user := []structs.User{}

	take := r.FormValue("take")
	page := r.FormValue("page")
	if take == "" || page == "" {
		take = "10"
		page = "0"
	} 


	connection.DB.
		Limit(take).
		Offset(page).
		Find(&user)


	res := structs.Result{Code: 200, Data: user, Message: "Get All Users, page = "+ page +" take = "+ take +" Success"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    var user structs.User
    connection.DB.Preload("RiskProfiles").
	First(&user, id)

	res := structs.Result{Code: 200, Data: user, Message: "Get User Detail"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var userUpdate structs.User
	json.Unmarshal(payloads, &userUpdate)

	var user structs.User
	connection.DB.First(&user, userID)
	connection.DB.Model(&user).Updates(&userUpdate)
	
	CountRisk(user.AGE)

	risk_profile := structs.RiskProfile{MM:mm, Bond:bond, Stock:stock}
	var riskProfile structs.RiskProfile
	connection.DB.Model(&riskProfile).Where("user_id = ?",user.UserID).Updates(&risk_profile)
	
	connection.DB.Preload("RiskProfiles").
	First(&user, userID)
 
	res := structs.Result{Code: 200, Data: user, Message: "Success update article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user structs.User
	connection.DB.First(&user, userID)
	connection.DB.Delete(&user)

	var riskProfile structs.RiskProfile
	connection.DB.Where("user_id = ?",user.UserID).Delete(&riskProfile)

	res := structs.Result{Code: 200, Message: "Success delete article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
