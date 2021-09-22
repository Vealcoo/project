package view

import (
	"fmt"
	"net/http"
	"tyr-project/controller"

	"github.com/gin-gonic/gin"
)

func WebPage(l *gin.Context) {
	l.HTML(http.StatusOK, "index.html", nil)
}

func LoginAuth(l *gin.Context) {
	info := controller.LoginInfo{}
	l.BindJSON(&info)

	id := info.UserId
	pw := info.UserPw
	auth := controller.Login(id, pw)
	if auth {
		token, _ := controller.SetToken(id)
		l.JSON(http.StatusOK, gin.H{
			"userid":  id,
			"message": "login success",
			"token":   token,
		})
	}
	if !auth {
		l.JSON(http.StatusOK, gin.H{
			"message": "login fail",
		})
	}
}

func LogoutAuth(l *gin.Context) {

}

func RegisterAuth(l *gin.Context) {
	info := controller.RegisterInfo{}
	l.BindJSON(&info)

	id := info.UserId
	pw := info.UserPw
	name := info.UserName
	err, auth := controller.Register(id, pw, name)
	if err != nil {
		fmt.Println("register fail")
	}
	if auth {
		l.JSON(http.StatusOK, gin.H{
			"userid":   id,
			"username": name,
			"message":  "register success!",
		})
	}
	if !auth {
		l.JSON(http.StatusOK, gin.H{
			"userid":   id,
			"username": name,
			"message":  "id exist!",
		})
	}
}

func ListInsert(l *gin.Context) {
	token := l.Request.Header.Get("Authorization")
	id := l.Query("userid")
	title := l.Query("listtitle")
	context := l.Query("listcontext")
	start := l.Query("starttime")
	end := l.Query("endtime")
	fmt.Println(id, title, context, start, end)
	timeup := false
	if controller.AuthJWT(token, id) {
		err := controller.Insert(id, title, context, start, end, timeup)
		if err != nil {
			fmt.Println(err)
		}
		l.JSON(http.StatusOK, gin.H{
			"userid":  id,
			"title":   title,
			"context": context,
			"message": "list build success",
		})
	}
	if !controller.AuthJWT(token, id) {
		l.JSON(http.StatusOK, gin.H{
			"userid":    id,
			"usertoken": token,
			"message":   "token verify fail",
		})
	}
}

func ListDelete(l *gin.Context) {
	token := l.Request.Header.Get("Authorization")
	id := l.Query("userid")
	listid := l.Query("listid")
	if controller.AuthJWT(token, id) {
		err := controller.Delete(listid)
		if err != nil {
			fmt.Println(err)
		}
		l.JSON(http.StatusOK, gin.H{
			"message": "list remove success",
		})
	}
	if !controller.AuthJWT(token, id) {
		l.JSON(http.StatusOK, gin.H{
			"userid":    id,
			"usertoken": token,
			"message":   "token verify fail",
		})
	}
}

func ListUpdate(l *gin.Context) {
	token := l.Request.Header.Get("Authorization")
	info := controller.UpdateListInfo{}
	l.BindJSON(&info)
	listid := info.ListId
	id := info.UserId
	title := info.ListTitle
	context := info.ListContext
	start := info.StartTime
	end := info.EndTime
	timeup := false
	if controller.AuthJWT(token, id) {
		err := controller.Update(listid, id, title, context, start, end, timeup)
		if err != nil {
			fmt.Println("")
		}
		l.JSON(http.StatusOK, gin.H{
			"userid":    id,
			"title":     title,
			"context":   context,
			"strattime": start,
			"endtime":   end,
			"message":   "list build success",
		})
	}
	if !controller.AuthJWT(token, id) {
		l.JSON(http.StatusOK, gin.H{
			"userid":    id,
			"usertoken": token,
			"message":   "token verify fail",
		})
	}
}

func ListDisplay(l *gin.Context) {
	token := l.Request.Header.Get("Authorization")
	info := controller.DisplayInfo{}
	l.BindJSON(&info)
	id := info.UserId
	if controller.AuthJWT(token, id) {
		listinfo := controller.Display(id)
		l.JSON(http.StatusOK, listinfo)
	}
	if !controller.AuthJWT(token, id) {
		l.JSON(http.StatusOK, gin.H{
			"userid":    id,
			"usertoken": token,
			"message":   "token verify fail",
		})
	}
}

func StartServer() {

	server := gin.Default()
	server.LoadHTMLGlob("view/template/html/*")
	//設定靜態資源的讀取
	// server.Static("/assets", "./template/assets")
	server.GET("/", WebPage)
	server.POST("/api/member", LoginAuth)
	server.GET("/api/member", RegisterAuth)

	server.GET("/api/member/todo", ListInsert)
	server.DELETE("/api/member/todo", ListDelete)
	server.PUT("/api/member/todo", ListUpdate)
	server.POST("/api/member/todo", ListDisplay)

	server.Run(":8887")
}
