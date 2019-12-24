package model

type UserInfo struct {
	Id           string         `json:"id" sql:"type:varchar(255), primary key, not null, unique"`
	Name         string         `json:"name" sql:"type:varchar(255), not null, unique"`
	Type         string         `json:"type" sql:"type:varchar(255), not null"`
	RelationShip []RelationShip `json:"relation_ship"`
}

type RelationShip struct {
	UserId string `json:"user_id" sql:"type:varchar(255), not null, unique"`
	State  string `json:"state" sql:"type:varchar(255), not null"`
	Type   string `json:"type" sql:"type:varchar(255), not null"`
}
