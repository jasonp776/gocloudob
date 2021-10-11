package main

import (
	//	"net/http"

	"log"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"psbank.com/gocloudob/controllers"
	"psbank.com/gocloudob/database"
	"psbank.com/gocloudob/middleware"
	"psbank.com/gocloudob/models"
)

func main() {
	router := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))

	log.Println("Connecting to DB")
	initDB()

	router.Use(sessions.Sessions("mysession", store))
	router.POST("/signup", controllers.Signup)
	router.GET("/login", controllers.Login)
	//router.GET("/logout", controllers.Logout)

	auth := router.Group("/auth")
	auth.Use(middleware.Authentication())
	{
		auth.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Everything is ok",
			})
		})
	}
	//router.GET("/login/:username", controllers.Login)

	//profile route
	//router.POST("/profile/:username", controllers.CreateProfile)

	/*router.GET("/customers", controllers.FindCustomers)
	router.POST("/customers", controllers.CreateCustomer)
	router.GET("/customer/:id", controllers.FindCustomerById)
	router.PATCH("/customer/:id", controllers.UpdateCustomerByID) // new
	router.DELETE("/customer/:id", controllers.DeleteCustomerByID)

	router.GET("/products", controllers.FindProducts)
	router.POST("/products", controllers.CreateProduct)
	router.GET("/product/:id", controllers.FindProductById)
	router.PATCH("/product/:id", controllers.UpdateProductByID) // new
	router.DELETE("/product/:id", controllers.DeleteProductByID)
	*/
	router.Run(":9080")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "password",
			DB:         "learning",
			//ServerName: os.Getenv("MySQLAddr"),
			//User:       os.Getenv("DBUSER"),
			//Password:   os.Getenv("DBPASSWORD"),
			//DB:         "learning",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	//database.Connector.AutoMigrate(&models.Credentials)
	//database.MigrateCustomer(&models.Customer{})
	//database.MigrateProduct(&models.Product{})
	database.MigrateCredentials(&models.Credentials{})
	database.MigrateProfile(&models.Profiles{})
}
