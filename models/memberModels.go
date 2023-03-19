package models

type Member struct {
	Id_member int    `json:"id_member,string" gorm:"primaryKey;autoIncrement:true"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
	Skintype  string `json:"skintype"`
	Skincolor string `json:"skincolor"`
}

type Product struct {
	Id_product   int     `json:"id_product,string" gorm:"primaryKey;autoIncrement:true"`
	Name_product string  `json:"name_product"`
	Price        float64 `json:"price"`
}

type Review_product struct {
	Id_review   int    `json:"id_review,string" gorm:"primaryKey;autoIncrement:true"`
	Id_product  int    `json:"id_product"`
	Id_member   int    `json:"id_member"`
	Decs_review string `json:"decs_review"`
}

type Like_review struct {
	Id_like   int    `json:"id_like,string" gorm:"primaryKey;autoIncrement:true"`
	Id_review int		`json:"id_review"`
	Id_product  int    `json:"id_product"`
	Id_member   int    `json:"id_member"`
	
}
