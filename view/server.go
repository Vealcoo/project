package view

import (
	"fmt"
	"net/http"
	"os"
	"tyr-project/controller"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY)")))

func WebPage(l *gin.Context) {
	l.HTML(http.StatusOK, "index.html", nil)
}

func sessionget(id string) {
	var w http.ResponseWriter
	var r *http.Request
	session, _ := store.Get(r, id)
	err := session.Save(r, w)
	if err != nil {
		fmt.Println(err)
	}
}

func LoginAuth(l *gin.Context) {
	id, _ := l.GetPostForm("userid")
	pw, _ := l.GetPostForm("userpw")
	auth := controller.Login(id, pw)
	if auth {
		sessionget(id)
	}
}

func RegisterAuth(l *gin.Context) {
	id, _ := l.GetPostForm("userid")
	pw, _ := l.GetPostForm("userpw")
	name, _ := l.GetPostForm("username")
	err := controller.Register(id, pw, name)
	if err != nil {
		fmt.Println("register fail")
	}
}

func ListInsert(l *gin.Context) {
	id, _ := l.GetPostForm("userid")
	title, _ := l.GetPostForm("listtitle")
	context, _ := l.GetPostForm("listcontext")
	start, _ := l.GetPostForm("starttime")
	end, _ := l.GetPostForm("endtime")
	timeup := false
	err := controller.Insert(id, title, context, start, end, timeup)
	if err != nil {
		fmt.Println(err)
	}
}

func ListDelete(l *gin.Context) {
	listid, _ := l.GetPostForm("listid")
	err := controller.Delete(listid)
	if err != nil {
		fmt.Println(err)
	}
}

func ListUpdate(l *gin.Context) {
	listid, _ := l.GetPostForm("listid")
	id, _ := l.GetPostForm("userid")
	title, _ := l.GetPostForm("listtitle")
	context, _ := l.GetPostForm("listcontext")
	start, _ := l.GetPostForm("starttime")
	end, _ := l.GetPostForm("endtime")
	timeup := false
	err := controller.Update(listid, id, title, context, start, end, timeup)
	if err != nil {
		fmt.Println("")
	}
}

func StartServer() {
	server := gin.Default()
	server.LoadHTMLGlob("template/html/*")
	//設定靜態資源的讀取
	// server.Static("/assets", "./template/assets")
	server.GET("/", WebPage)
	server.POST("/login", LoginAuth)
	server.POST("/register", RegisterAuth)
	server.POST("/insert", ListInsert)
	server.POST("/delete", ListDelete)
	server.POST("/update", ListUpdate)

	server.Run(":8887")
}
