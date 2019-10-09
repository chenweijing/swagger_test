package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/chenweijing/swagger_test/docs"
	"github.com/gin-gonic/gin"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

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

	fmt.Println("swagger server start.")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello/{name}", index).Methods("GET")
	router.HandleFunc("/swagger.yaml", swagger).Methods("GET")
	router.HandleFunc("/api/v1/mytest", handleMytest).Methods("GET")

	// router.HandleFunc("/api/docs", handleDocs).Methods("GET")

	router.PathPrefix("/docs").HandlerFunc(handleDocs).Methods(http.MethodGet)

	// handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":9091", router))
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

// GroupListWithMember 展示附加成员信息的群组信息列表
// @Tags group-info 群聊信息
// @Summary 展示群信息列表
// @Description 展示附加成员信息的群组信息列表
// @Produce  json
// @Param operator_id query int false "运营商"
// @Param group_id query string false "组ID"
// @Param group_no query string false "组编号"
// @Param group_name query string false "群名称"
// @Param member_id query string false "成员用户ID"
// @Param member_mlid query string false "成员用户密聊ID"
// @Param member_role query int false "成员用户角色"
// @Param state query int false "群组状态;-1 已冻结,0-已关闭, 1-开启中"
// @Param start_time query string false "创建时间-开始范围2018-01-01 00:00:00"
// @Param end_time query string false "创建时间-结束范围2020-01-01 00:00:00"
// @Param offset query int false "偏移量" default(0)
// @Param limit query int false "限制数" default(20)
// @Success 200 {object} string "附加成员信息的群组信息"
// @Failure 500 {object} string "内部错误"
// @Security ApiKeyAuth
// @Security OAuth2Implicit[admin, write]
// @Router /cms/api/group [get]
func handleMytest(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to /api/v1/mytest request")
	fmt.Fprintln(w, "Hello:测试完毕")
}
