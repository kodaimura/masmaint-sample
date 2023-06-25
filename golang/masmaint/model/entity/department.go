package entity


type Department struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	ManagerId int `db:"manager_id" json:"manager_id"`
	Location string `db:"location" json:"location"`
	Budget float32 `db:"budget" json:"budget"`
	CreateAt string `db:"create_at" json:"create_at"`
	UpdateAt string `db:"update_at" json:"update_at"`
}