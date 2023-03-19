package route

import (

	"khub/database"

	"khub/models"

	"github.com/gofiber/fiber/v2"
)

//insert data ke dalam table member
func Insert_member(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"Massage": "Harap Masukan Data ",
		})
	}
	user := models.Member{
		Username:  data["username"],
		Gender:    data["gender"],
		Skintype:  data["skintype"],
		Skincolor: data["skincolor"],
	}
	if user.Username ==""{
	    return c.JSON(fiber.Map{
            "Massage": "Harap Masukan Username",
        })
	}
	//apakah username telah di pakai ?
	var username models.Member
	database.DB.Where("username = ?", user.Username).Select("username").Find(&username)
	if user.Username == username.Username {
	    return c.JSON(fiber.Map{
			"status": "invalid",
            "Massage": "Username Telah Digunakan Harap Masukan Username yang lain",
        })
	} else {
		database.DB.Create(&user)
		return c.JSON(fiber.Map{
		"status": "Success",
		"massage": "Selamat datang",
	})}
}

//Update data member berdasarkan id member yang di 
func Update_member(c *fiber.Ctx) error {
	var data map[string]string
	
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var member models.Member
	id_member := c.Params("id_member")
	//cek apakah id_member sudah terdata di database 
	database.DB.Where("id_member", id_member).Find(&member)
	if member.Id_member == 0 {
		return c.JSON(fiber.Map{
            "Massage": "Maaf tidak ada data member, mohon daftarkan member terlebih dahulu ",
        })
    }
	//masukan data yang akan di ubah
	update := models.Member{
		Username:  data["username"],
        Gender:    data["gender"],
        Skintype:  data["skintype"],
        Skincolor: data["skincolor"],
	}

	//query update member
	database.DB.Model(&member).Where("Id_member = ?", id_member).Updates(update)
	return c.JSON(fiber.Map{
		"massage": "Data berhasil di ubah, mohon di cek kembali ya",
	})
}
func Getdatamember(c* fiber.Ctx)error{
	var user models.Member
	
	id_member := c.Params("id_member")
	//cek apakah id_member sudah terdata di database
	database.DB.Where("id_member", id_member).Find(&user)
    if user.Id_member == 0 {
        return c.JSON(fiber.Map{
            "status": "invalid",
            "Massage": "Member tidak di temukan",
        })
    } else {
		return c.JSON(fiber.Map{
        "status": "Success",
        "data": user,})
	}
}

func Getall_member(c *fiber.Ctx) error {

	var users []models.Member

	result1 := database.DB.Find(&users)
	if result1.Error != nil {
		return c.JSON(fiber.Map{
			"error": result1.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":  users,
		"count": len(users),
	})

}

func Delete_member(c *fiber.Ctx) error {
	//Delate data Member di lakukan secara permanenen yang dmna data di haous dari dataabse member 
	id_member := c.Params("id_member")
	// cek apakah member terdaftar di dalam database
	var member models.Member
	var review_products models.Review_product
	var like_review models.Like_review
	database.DB.Where("id_member", id_member).Find(&member)
	if member.Id_member == 0 {
		return c.JSON(fiber.Map{
            "Massage": "Maaf Data yang anda masukkan tidak ada, mohon masukan data terlebih dahulu ",
        })
    }
	//hapus database member beserta product yang di like dan review yang di berikan 
	database.DB.Where("id_member =?", id_member).Delete(&review_products)
	//hapus Database Like_review 
	database.DB.Where("id_member =?", id_member).Delete(&like_review)
	//hapus data member
    database.DB.Where("id_member =?", id_member).Delete(&member)

    return c.JSON(fiber.Map{
        "massage": "Member berhasil dihapus ",
    })


}