package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/chenweijing/swagger_test/docs"
	"github.com/gin-gonic/gin"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// https://razeencheng.com/post/go-swagger.html
// @Param 1.参数名 2.参数类型 3.参数数据类型 4.是否必须 5.参数描述 6.其他属性
// @Param id path integer true "文件ID"

// 参数类型
// path 	该类型参数直接拼接在URL中，如Demo中HandleGetFile：
// query 	该类型参数一般是组合在URL中的，如Demo中HandleHello
// formData 该类型参数一般是POST,PUT方法所用，如Demo中HandleLogin
// body 	当Accept是JSON格式时，我们使用该字段指定接收的JSON类型

// @Summary 测试SayHello
// @Description 向你说Hello
// @Tags 测试
// @Accept mpfd
// @Produce json
// @Param who query string true "人名"
// @Success 200 {string} json "{"msg": "hello Razeen"}"
// @Failure 400 {string} json "{"msg": "who are you"}"
// @Router /hello [get]

func HandleHello(c *gin.Context) {
	who := c.Query("who")

	if who == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "who are u?"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "hello " + who})
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://razeen.me

// @contact.name Razeen
// @contact.url https://razeen.me
// @contact.email me@razeen.me

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:9090
// @BasePath /api/v1

func main() {
	//r := gin.Default()
	//store := sessions.NewCookieStore([]byte("secret"))
	// r.Use(sessions.Sessions("mysession", store))

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//v1 := r.Group("/api/v1")
	//{
	//	v1.GET("/hello", HandleHello)
	//v1.POST("/login", HandleLogin)
	//v1Auth := r.Use(HandleAuth)
	//{
	//	v1Auth.POST("/upload", HandleUpload)
	//	v1Auth.GET("/list", HandleList)
	//}
	//}

	//r.Run(":9090")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello/{name}", index).Methods("GET")
	router.HandleFunc("/swagger.json", swagger).Methods("GET")
	router.HandleFunc("/docs", handleDocs).Methods("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":9090", handler))
}

func swagger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "docs/swagger.json")
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to /hello request")
	log.Println(r.UserAgent())

	vars := mux.Vars(r)
	name := vars["name"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello:", name)
}

func handleDocs(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to /docs request")

	if r.RequestURI == "/docs" {
		http.Redirect(w, r, "docs/index.html", http.StatusTemporaryRedirect)
	} else {
		httpSwagger.WrapHandler.ServeHTTP(w, r)
	}
}
