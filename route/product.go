package route

import (
	"khub/database"
	"khub/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Insert_reviews_produk(c* fiber.Ctx) error {
	 var member models.Member
	 var product models.Product
	 id_member := c.Params("id_member")
	 id_product := c.Params("id_product")

	 //periksa apakah ada member sesuai dengan database 
	 database.DB.Where("id_member = ?", id_member).Select("id_member").Find(&member)
	
	if x := member.Id_member ; x == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "false",
            "message": "Member yang ditemukan tidak valid",
        })
    }
	//periksa apakah ada product ada di database atau tidak
	database.DB.Where("id_product = ?", id_product).Select("id_product").Find(&product)
	
	if x := product.Id_product ; x == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "false",
            "message": "data product yang ditemukan tidak valid",
        })
    }

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"Massage": "Harap Masukan Data ",
		})
	}
	type Review_product struct {  
		Id_review int `json:"id_review"`
		Id_product  int    
		Id_member   int    
		Decs_review string 
	}
	//var review_product Review_product
	reviews := models.Review_product{
		Id_product: product.Id_product,
		Id_member :member.Id_member,
		Decs_review: data["decs_review"],
	}
	database.DB.Create(&reviews)

	return c.JSON(fiber.Map{
        "status":  "true",
        "message": "Reviews Product Telah di tambah",
		"data": reviews,
    })



}


func Review_produk(c *fiber.Ctx) error {

	type Review_produk struct {
		Products struct {
			Id_product int
			Name_product string
			Price        float32
		}
		Reviews []struct {
			Id_member   int
			Decs_review string
			Username   string
			Gender     string
			Skintype   string
			Skincolor  string
		}
	}
	var productss Review_produk
	type Product struct {
		Id_product int
	}
	var  product Product

	//variabel id_produk atau bisa di bikin query
	id_produk := c.Params("id_product")
	id_produkss,err := strconv.Atoi(id_produk)
	if err!= nil {
        return nil
   }

	//cek apakah data product ada di database ? 
	database.DB.Where("id_product = ?", id_produkss).Select("id_product").Find(&product)
	
	if x := product.Id_product ; x == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "false",
            "message": "data product yang ditemukan tidak valid",
        })
    }

	//queri 1 untuk menyimpan id_product
	err2 := database.DB.Table("review_products").
		Select("review_products.id_product, review_products.id_member, review_products.decs_review, members.username, members.gender, members.skintype, members.skincolor, products.name_product, products.price").
		Joins("INNER JOIN members ON review_products.id_member = members.id_member").
		Joins("INNER JOIN products ON review_products.id_product = products.id_product").
		Where("review_products.id_product = ?", id_produk).
		Scan(&productss.Products)
	if err2.Error != nil {
		return c.JSON(fiber.Map{
			"error": err2.Error.Error(),
		})
	}

	//queri 2 unutk menymoan review products
	err3 := database.DB.Table("review_products").
		Select("review_products.id_product, review_products.id_member, review_products.decs_review, members.username, members.gender, members.skintype, members.skincolor, products.name_product, products.price").
		Joins("INNER JOIN members ON review_products.id_member = members.id_member").
		Joins("INNER JOIN products ON review_products.id_product = products.id_product").
		Where("review_products.id_product = ?", id_produk).
		Scan(&productss.Reviews)
	if err3.Error != nil {
		return c.JSON(fiber.Map{
			"error": err3.Error.Error(),
		})
	}
	//ambil data like dari product
	type Like_review struct{
		Is_like string
	}
	var like []Like_review
	
	//cek data like dari product
	database.DB.Model(&Like_review{}).Where("id_product = ? and is_like = ?", id_produk, 1).Find(&like)
	likes := len(like)

	return c.JSON(fiber.Map{
		"data":   productss,
		"like": likes,
		"status": "ok",

	})
}



func Cek_like_product(c *fiber.Ctx) error {

	id_member := c.Params("id_member")
	id_produk := c.Params("id_product")

	//cek data prod
	var produk models.Product
	err := database.DB.Where("id_product = ?", id_produk).Select("id_product", "name_product", "price").Find(&produk)
	if err.Error != nil {
		return c.JSON(fiber.Map{
			"error": err.Error.Error(),
		})
	}

	if x := produk.Id_product; x == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "false",
            "message": "data product yang ditemukan tidak valid",
        })
	}
	//cek apakah data member 
	var member models.Member
    err = database.DB.Where("id_member =?", id_member).Select("id_member").Find(&member)
    if err.Error!= nil {
        return c.JSON(fiber.Map{
            "error": err.Error.Error(),
        })
    }

    if x := member.Id_member; x == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "false",
            "message": "data member yang ditemukan tidak valid",
        })
    }

	//cek produk sudah di like atau belum
	type Like_review struct {
		Is_like int
	}
	var like_review Like_review

	err2 := database.DB.Where("id_product = ? AND id_member = ?", id_produk, id_member).Select("is_like").Find(&like_review)
	if err2.Error != nil {
		return c.JSON(fiber.Map{
			"error": err2.Error.Error(),
		})
	}

	if hasil := like_review.Is_like; hasil == 0 {
		c.JSON(fiber.Map{
			"is_like": 0,
		})
	 
	} else {
		c.JSON(fiber.Map{
			"is_like": 1,
		})
	
	}

	return c.JSON(fiber.Map{
		"status": "ok",
		"datalike": Like_review{
			Is_like: like_review.Is_like,
		},
		"product": produk,
	})

}

func Like_unlike_product(c *fiber.Ctx) error {
	id_member := c.Params("id_member")
	id_produk := c.Params("id_product")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	type Like_review struct {
		Id_review  int
		Id_like    int
		Is_like    string
		Id_product string
		Id_member  string
	}
	var like_review Like_review

	//action := 1

	//cek data produk
	var produk models.Product
	err := database.DB.Where("id_product = ?", id_produk).Select("id_product", "name_product", "price").Find(&produk)
	if err.Error != nil {
		return c.JSON(fiber.Map{
			"error": err.Error.Error(),
		})
	}

	if x := produk.Id_product; x == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "TIDAK ADA PRODUCT SILAHKAN PILIH YANG BENAR",
		})
	}
	//Cek Data member 
	var member models.Member
    err = database.DB.Where("id_member =?", id_member).Select("id_member").Find(&member)
    if err.Error!= nil {
        return c.JSON(fiber.Map{
            "error": err.Error.Error(),
        })
    }

    if x := member.Id_member; x == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "data member yang ditemukan tidak valid",
        })
	}

	//cek apakah data like sudah ada ? kalo belum ada di tambah kalau sudah ada di update

	err3 := database.DB.Where("id_product = ? AND id_member = ?", id_produk, id_member).Select("is_like", "id_like", "id_product", "id_member").Find(&like_review)
	if err3.Error != nil {
		return c.JSON(fiber.Map{
			"error": err3.Error.Error(),
		})
	}

	type Action struct {
		Action string
	}
	action := Action{
		Action: data["action"],
	}

	//Function unutk menambhakan dan mengupdate
	//jika action ngelike = 1
	if y := action.Action; y == "1" {
		//jika id like belum ada maka di tambahkan data like
		if x := like_review.Id_like; x == 0 {
			like := Like_review{
				Is_like:    action.Action,
				Id_product: id_produk,
				Id_member:  id_member,
			}
			database.DB.Select("id_product", "id_member", "is_like").Create(&like)
			return c.JSON(like_review)

		} else {
			//kalau data like dengan id_like nya ada sesuai dengan id_product dan id_member, bisa ubah data ke like = 1

			id_like := like_review.Id_like
			database.DB.Model(&like_review).Where("id_like =?", id_like).Update("is_like", 1)
		}

	} else {
		if x := like_review.Id_like; x == 0 {
			return c.JSON(fiber.Map{
				"error": "anda blom pernah like !!!",
			})
		}else {
		
			id_like := like_review.Id_like
			database.DB.Model(&like_review).Where("id_like =?", id_like).Update("is_like", 0)
		}
	}

	return c.JSON(fiber.Map{
		"data": like_review,
	})
}
