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

func WithdrawWallet(c *gin.Context) {

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

		addWithdraw := models.Withdraw{WithdrawnBy: id, Status: "success", Amount: amount, WithdrawnAt: t, ReferenceId: referenceId}
		if err := db.Create(&addWithdraw).Error; err != nil {
			fmt.Printf("error add : %3v \n", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}else{
			newBalance := wallet.Balance - amount
			db.Model(models.Wallet{}).Where("owned_by = ?", id).Updates(map[string]interface{}{"balance": newBalance, "updated_at": t})
		}

		d := models.ShowWithdraw{addWithdraw.ID,addWithdraw.WithdrawnBy,addWithdraw.Status,addWithdraw.WithdrawnAt,addWithdraw.Amount,addWithdraw.ReferenceId}

		res := new(models.Res1)
		res.Status = "success"
		//this is JSON in database
		var requirement json.RawMessage

		appendRes5 := models.Res5{d}
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(appendRes5)

		requirement = []byte(reqBodyBytes.Bytes())
	
		res.Requirement = &requirement

		c.JSON(200, res)

	} else {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}