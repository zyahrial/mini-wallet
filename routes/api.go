  package routes

  import (
    "time"
    "fmt"
    // "os"
  
    "github.com/gin-gonic/gin"
    migrations "khaerus/mini-wallet/db/migrations"
    controllers "khaerus/mini-wallet/controllers"
    // "khaerus-mini-wallet/models"
  	// "khaerus-mini-wallet/conf"
    "khaerus/mini-wallet/db/database"
  )

  func Route() {
    // time := conf.Viper("APP_TZ")
    fmt.Printf("Started at : %3v \n", time.Now())

    //InitPostgres()
    database.InitDB()
    // database.InitGormPostgres()
    defer database.DBCon.Close()

    migrations.Migrate()

    // Set the router as the default one shipped with Gin
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()

    // Setup route group for the API
    api := router.Group("/api/v1")
    api.POST("/wallet", controllers.ValidateAccount)
    api.POST("/wallet/enable", controllers.EnableWallet)
    api.GET("/wallet", controllers.GetWallet)
    api.PATCH("/wallet", controllers.DisableWallet)
    
    api.POST("/wallet/deposits", controllers.DepositWallet)
    api.POST("/wallet/withdrawals", controllers.WithdrawWallet)

    router.Run(":8000")
  }