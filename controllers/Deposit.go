package controllers                                                                                                                                                                                              

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"time"
	// "khaerus/mini-wallet/conf"
	"strconv"
	// "log"
	"github.com/gin-gonic/gin"
    // "github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"

	"khaerus/mini-wallet/models"
	"khaerus/mini-wallet/db/database"
)

func DepositWallet(c *gin.Context) {

	// var deposit models.Deposit
	var wallet models.Wallet

	// t := time.Now()
	amount, _ := strconv.ParseInt(c.PostForm("amount"), 10, 64)
	referenceId := c.PostForm("reference_id")

	const BEARER_SCHEMA = "Token "
	
	authHeader := c.GetHeader("Authorization")
	myToken := authHeader[len(BEARER_SCHEMA):]

	token, err := JWTAuthService().ValidateToken(myToken)
	
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {	

		id := claims["id"].(string)

		db := database.DBCon
		if err := db.Where("owned_by = ?", id).First(&wallet).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found!"})
			return
		}

		t := time.Now()

		addDeposit := models.Deposit{DepositedBy: id, Status: "success", Amount: amount, DepositeAt: t, ReferenceId: referenceId}
		if err := db.Create(&addDeposit).Error; err != nil {
			fmt.Printf("error add : %3v \n", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}else{

			// geWallet := db.Where("owned_by = ?", id).Select("balance").First(&wallet)
			// fmt.Println(geWallet)
			newBalance := wallet.Balance + amount

			db.Model(models.Wallet{}).Where("owned_by = ?", id).Updates(map[string]interface{}{"balance": newBalance, "updated_at": t})
		}

		d := models.ShowDeposit{addDeposit.ID,addDeposit.DepositedBy,addDeposit.Status,addDeposit.DepositeAt,addDeposit.Amount,addDeposit.ReferenceId}

		res := new(models.Res1)
		res.Status = "success"
		//this is JSON in database
		var requirement json.RawMessage

		appendRes4 := models.Res4{d}
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(appendRes4)

		requirement = []byte(reqBodyBytes.Bytes())
	
		res.Requirement = &requirement

		c.JSON(200, res)

	} else {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}