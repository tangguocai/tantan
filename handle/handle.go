package handle

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	. "tantan/dbhelper"
	"tantan/model"
	"tantan/util"
)

func MyGetHandler(w http.ResponseWriter, r *http.Request) {
	/*// parse query parameter
	vals := r.URL.Query()
	param, _ := vals["servicename"] // get query parameters

	// composite response body
	var res = map[string]string{"result": "succ", "name": param[0]}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)*/
	var res map[string]string = make(map[string]string)
	var status = http.StatusOK

	vals := r.URL.Query()
	param, ok := vals["name"]
	if (!ok) {
		res["result"] = "fail"
		res["error"] = "required parameter name is missing"
		status = http.StatusBadRequest
	} else {
		res["result"] = "succ"
		res["name"] = param[0]
		status = http.StatusOK
	}

	response, _ := json.Marshal(res)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func MyPostHandler(w http.ResponseWriter, r *http.Request) {
	// parse path variable
	vars := mux.Vars(r)
	servicename := vars["servicename"]

	// parse JSON body
	var req map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &req)
	servicetype := req["servicetype"].(string)

	// composite response body
	var res = map[string]string{"result": "succ", "name": servicename, "type": servicetype}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	userDetail := &model.UserDetail{}
	json.Unmarshal(body, &userDetail)
	userDetail.Id, _ = util.NewUUID()

	if err := Pg.Conn.Insert(userDetail); err != nil {
		log.Println("Insert Error:", err)
		return
	}

	// composite response body
	var res = map[string]string{"id": userDetail.Id, "name": userDetail.Name, "type": userDetail.Type}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func ListAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	userInfoArray := []model.UserDetail{}
	Pg.Conn.Model(&userInfoArray).Select()
	userInfoArrayResponse := []model.UserInfo{}
	for _, v := range userInfoArray {
		item := model.UserInfo{}
		item.Id = v.Id
		item.Name = v.Name
		item.Type = v.Type
		userInfoArrayResponse = append(userInfoArrayResponse, item)
	}
	response, _ := json.Marshal(userInfoArrayResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func RelationShipHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	otherUserId := vars["other_user_id"]
	body, _ := ioutil.ReadAll(r.Body)
	relation := model.RelationShip{}
	json.Unmarshal(body, &relation)
	relation.UserId = otherUserId
	relation.Type = "relationship"

	userDetail := &model.UserDetail{Id: userId}
	Pg.Conn.Select(userDetail)
	if relation.State == "liked" {
		for _, v := range userDetail.RelationShip {
			if v.UserId == otherUserId {
				if v.State == relation.State {
					relation.State = "matched"
				}
			}
		}
	}
	relationArray := []model.RelationShip{}
	relationArray = append(relationArray, relation)
	userDetail.RelationShip = relationArray

	Pg.Conn.Update(userDetail)

	var res = map[string]string{"user_id": otherUserId, "state": relation.State, "type": relation.Type}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func ListUserRelationShipHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	userDetail := &model.UserDetail{Id: userId}
	Pg.Conn.Select(userDetail)
	relationArray := []model.RelationShip{}
	for _, v := range userDetail.RelationShip {
		relationArray = append(relationArray, v)
	}

	response, _ := json.Marshal(relationArray)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

