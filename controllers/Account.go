package controllers                                                                                                                                                                                              

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"time"
	// "khaerus/mini-wallet/conf"

	// "log"
	"github.com/gin-gonic/gin"
    // "github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"

	"khaerus/mini-wallet/models"
	"khaerus/mini-wallet/db/database"
)

func ValidateAccount(c *gin.Context) {

	id := c.PostForm("customer_xid")

	if id == "" {
		c.JSON(http.StatusOK, gin.H{"status": "customer_xid is required!"})
		return
	}

    var account models.Account

	db := database.DBCon
    if err := db.Where("id = ?", id).First(&account).Error; err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Account doesn't exist!"})
        return
    }

	
	token := JWTAuthService().GenerateToken(id, true)
	// token := AuthClaim(id)

	res := new(models.Res1)
	res.Status = "success"
	//this is JSON in database
	var requirement json.RawMessage

	appendRes2 := models.Res2{token}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(appendRes2)

	requirement = []byte(reqBodyBytes.Bytes())

	res.Requirement = &requirement

	c.JSON(200, res)
}

func EnableWallet(c *gin.Context) {

	const BEARER_SCHEMA = "Token "
	
	authHeader := c.GetHeader("Authorization")
	myToken := authHeader[len(BEARER_SCHEMA):]

	token, err := JWTAuthService().ValidateToken(myToken)
	
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {	
		var wallet models.Wallet
		c.BindJSON(&wallet)
	
		id := claims["id"].(string)
		t := time.Now()

		db := database.DBCon
		if err := db.Where("owned_by = ?", id).First(&wallet).Error; err != nil {
		}else{
			c.JSON(http.StatusBadRequest, gin.H{"error": "already enabled!"})
			return
		}

		addWallet := models.Wallet{OwnedBy: id, Status: "enabled", Balance: 0, EnableAt: t}
		if err := db.Create(&addWallet).Error; err != nil {
			fmt.Printf("error add : %3v \n", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		d := models.ShowWallet{addWallet.ID,addWallet.OwnedBy,addWallet.Status,addWallet.Balance,addWallet.EnableAt}

		res := new(models.Res1)
		res.Status = "success"
		//this is JSON in database
		var requirement json.RawMessage

		appendRes3 := models.Res3{d}
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(appendRes3)

		requirement = []byte(reqBodyBytes.Bytes())
	
		res.Requirement = &requirement

		c.JSON(200, res)

	} else {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func GetWallet(c *gin.Context) {

	const BEARER_SCHEMA = "Token "
	
	authHeader := c.GetHeader("Authorization")
	myToken := authHeader[len(BEARER_SCHEMA):]

	token, err := JWTAuthService().ValidateToken(myToken)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {	

		var wallet models.Wallet
		c.BindJSON(&wallet)
	
		id := claims["id"].(string)
		// t := time.Now()

		db := database.DBCon
		if err := db.Where("owned_by = ?", id).First(&wallet).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found!"})
			return
		}

		fmt.Println(wallet)

		d := models.ShowWallet{wallet.ID,wallet.OwnedBy,wallet.Status,wallet.Balance,wallet.EnableAt}

		res := new(models.Res1)
		res.Status = "success"
		//this is JSON in database
		var requirement json.RawMessage

		appendRes3 := models.Res3{d}
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(appendRes3)

		requirement = []byte(reqBodyBytes.Bytes())
	
		res.Requirement = &requirement

		c.JSON(200, res)

	} else {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}