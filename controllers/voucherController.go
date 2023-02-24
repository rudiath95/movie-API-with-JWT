package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rudiath95/movie-API-with-JWT/ini"
	"github.com/rudiath95/movie-API-with-JWT/models"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func AddVoucher(c *gin.Context) {

	//Call DB
	var body struct {
		VoucherAmount uint
	}

	c.Bind(&body)

	//Create Query
	post := models.VoucherList{VoucherCode: randSeq(10), VoucherAmount: body.VoucherAmount}

	result := ini.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to add voucher",
		})
		return
	}

	//Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func RedeemVoucher(c *gin.Context) {
	//get id
	user, _ := c.Get("user")
	id := user.(models.User).ID

	//Call DB
	var body struct {
		VoucherCode string
	}
	var voc models.VoucherList

	c.Bind(&body)
	ini.DB.Raw("SELECT * FROM voucher_lists WHERE voucher_code = ?", body.VoucherCode).Scan(&voc)

	if voc.VoucherStatus == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Voucher already Redeemed",
		})
		return
	}

	var post models.UserInfo
	ini.DB.Raw("SELECT current_balance FROM user_infos WHERE user_id = ?", id).Scan(&post)

	total := post.CurrentBalance + voc.VoucherAmount
	ini.DB.Raw("UPDATE user_infos SET current_balance = ? WHERE user_id = ? ",
		total, id).Scan(&post)
	ini.DB.Raw("UPDATE voucher_lists SET voucher_status = ? WHERE voucher_code = ? ",
		false, body.VoucherCode).Scan(&voc)

	//Return
	c.JSON(200, gin.H{
		"Success Redeemed":    voc.VoucherAmount,
		"Your Balance Become": total,
	})
}
