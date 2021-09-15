package view

import (
	"fmt"
	"net/http"
	"tyr-project/controller"

	"github.com/gin-gonic/gin"
)

// var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY)")))

func WebPage(l *gin.Context) {
	l.HTML(http.StatusOK, "index.html", nil)
}

func LoginAuth(l *gin.Context) {
	id := l.Query("userid")
	pw := l.Query("userpw")
	auth := controller.Login(id, pw)
	if auth {
		// store.MaxAge(1000)
		// session := store.Get(l.Request, id)
		// session.Values["auth"] = true
		// err := session.Save(l.Request, l.Writer)
		// if err != nil {
		// 	http.Error(l.Writer, err.Error(), http.StatusInternalServerError)
		// }
		controller.SetSession(id, l.Writer)
		l.JSON(http.StatusOK, gin.H{
			"userid":  id,
			"message": "login success",
		})
	}
	if !auth {
		l.JSON(http.StatusOK, gin.H{
			"message": "login fail",
		})
	}
}

func LogoutAuth(l *gin.Context) {
	fmt.Println("session clear")
	controller.ClearSession(l.Writer)
}

func RegisterAuth(l *gin.Context) {
	id := l.Query("userid")
	pw := l.Query("userpw")
	name := l.Query("username")
	err := controller.Register(id, pw, name)
	if err != nil {
		fmt.Println("register fail")

	}
	l.JSON(http.StatusOK, gin.H{
		"userid":   id,
		"username": name,
		"message":  "register success",
	})
}

func ListInsert(l *gin.Context) {
	id := l.Query("userid")
	title := l.Query("listtitle")
	context := l.Query("listcontext")
	start := l.Query("starttime")
	end := l.Query("endtime")
	timeup := false
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

func ListDelete(l *gin.Context) {
	listid := l.Query("listid")
	err := controller.Delete(listid)
	if err != nil {
		fmt.Println(err)
	}
	l.JSON(http.StatusOK, gin.H{
		"message": "list remove success",
	})
}

func ListUpdate(l *gin.Context) {
	listid := l.Query("listid")
	id := l.Query("userid")
	title := l.Query("listtitle")
	context := l.Query("listcontext")
	start := l.Query("starttime")
	end := l.Query("endtime")
	timeup := false
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

func ListDisplay(l *gin.Context) {
	// id, _ := l.GetPostForm("userid")
	// listinfo := controller.Display(id)
}

// func test(w http.ResponseWriter, r *http.Request) {

// }

// func LoginAuth(w http.ResponseWriter, r *http.Request) {
// 	id := r.FormValue("userid")
// 	pw := r.FormValue("userpw")
// 	auth := controller.Login(id, pw)
// 	if auth {
// 		sessionget(id)
// 	}
// }

// func RegisterAuth(w http.ResponseWriter, r *http.Request) {
// 	id := r.FormValue("userid")
// 	pw := r.FormValue("userpw")
// 	name := r.FormValue(("username"))
// 	err := controller.Register(id, pw, name)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

func StartServer() {
	// http.HandleFunc("/", test)
	// http.HandleFunc("/login", LoginAuth)
	// http.HandleFunc("/register", RegisterAuth)
	// err := http.ListenAndServe(":9090", nil) //設定監聽的埠
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
	server := gin.Default()
	server.LoadHTMLGlob("view/template/html/*")
	//設定靜態資源的讀取
	// server.Static("/assets", "./template/assets")
	server.GET("/", WebPage)
	server.POST("/api/login", LoginAuth)
	server.POST("/api/logout", LogoutAuth)
	server.POST("/api/register", RegisterAuth)
	server.POST("/api/insert", ListInsert)
	server.POST("/api/delete", ListDelete)
	server.POST("/api/update", ListUpdate)
	server.POST("/api/dispaly", ListDisplay)

	server.Run(":8887")
}
