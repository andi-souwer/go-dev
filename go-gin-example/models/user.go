package models

type User_statu struct {
	Model
	User_name  string `josn: "user_name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func Getusers(pageNum int, pageSize int, maps interface{}) (users []User_statu){
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users)

	return
}