package main

import (
	"io"
	"net/http"
	"os"
	"sotsusei/data"
	"sotsusei/db"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func main() {

	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)
	log.SetOutput(gin.DefaultWriter)

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("Session", store))

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "assets")

	// トップページ
	router.GET("/", func(c *gin.Context) {
		// session := sessions.Default(c)

		// log.Print(session.Get("loginUser"))
		user, _ := c.Cookie("loginUser")
		log.Print(user)
		products := db.GetLatestProducts()
		// log.Printf("%+v\n", products)
		// log.Print(products[1])
		c.HTML(http.StatusOK, "index.html", products)
	})

	// 検索結果画面
	router.GET("/search", func(c *gin.Context) {
		name := c.Query("name")
		cat := c.Query("category")

		if name == "" && cat != "" {
			// nameで検索

		} else if name != "" && cat == "" {
			// categoryで検索

		} else {
			// その他はトップに戻る
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	})

	// 商品ページ
	router.GET("/product", func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.Redirect(http.StatusMovedPermanently, "/")
		}
		product := db.GetProduct(id)
		product.FoodName = db.GetFood(product.FoodId)
		log.Printf(string(product.FoodId))
		log.Printf(product.FoodName)
		product.DeadlineStr = product.Deadline.Format("2006年01月02日")
		c.HTML(http.StatusOK, "product.html", product)
	})

	// パスワード認証画面
	router.GET("/passwordcertification", func(c *gin.Context) {
		c.HTML(http.StatusOK, "changecomplete.html", nil)
	})

	// メール送信完了画面
	router.GET("/sendemail", func(c *gin.Context) {
		c.HTML(http.StatusOK, "changepasswordemailcomplete.html", nil)
	})

	// 会員登録画面
	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "createaccount.html", nil)
	})

	router.POST("/signup", func(c *gin.Context) {

		var newUser data.NewUser
		err := c.Bind(&newUser)
		if err != nil {
			c.Status(http.StatusBadRequest)
		} else {
			// emailが重複するかどうか
			if db.IsDuplicate(newUser) {
				c.HTML(http.StatusBadRequest, "createaccount.html", gin.H{"err": "このメールアドレスはすでに登録されています。"})
			} else {
				c.HTML(http.StatusOK, "createaccountconfirm.html", newUser)
			}
		}
	})

	// 会員登録確認画面
	// router.GET("/signupconfirm", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "createaccountconfirm.html", nil)
	// })

	// 会員登録完了画面
	router.GET("/signupcomplete", func(c *gin.Context) {
		c.HTML(http.StatusOK, "accountcomplete.html", nil)
	})

	router.POST("/signupcomplete", func(c *gin.Context) {
		// session := sessions.Default(c)

		var newUser data.NewUser
		err := c.Bind(&newUser)
		if err != nil {
			c.Status(http.StatusBadRequest)
		} else {
			// 会員登録処理
			log.Printf("%+v\n", newUser)
			if err := db.SignUp(newUser); err != nil {
				c.Status(http.StatusBadRequest)
			}

			c.SetCookie("loginUser", newUser.Email, 3600, "/", "localhost", false, false)
			c.SetCookie("loginName", newUser.UserName, 3600, "/", "localhost", false, false)
			c.HTML(http.StatusOK, "accountcomplete.html", newUser)

		}
	})

	// お届け先編集画面
	router.GET("/editdelivery", func(c *gin.Context) {
		c.HTML(http.StatusOK, "editdelivery.html", nil)
	})

	// 会員情報編集画面
	router.GET("/editinfo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "editinfo.html", nil)
	})

	// パスワード編集画面
	router.GET("/editpassword", func(c *gin.Context) {
		c.HTML(http.StatusOK, "editpassword.html", nil)
	})

	// ログイン画面
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "loginaccount.html", nil)
	})

	router.POST("/login", func(c *gin.Context) {
		logintext := ""
		formPassword := c.PostForm("password")
		loginEmail := c.PostForm("email")
		dbPassword := db.GetUser(loginEmail).Password
		// log.Print(dbPassword)
		// session := sessions.Default(c)

		if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(formPassword)); err != nil {
			logintext = "ログインに失敗しました。"
			c.HTML(http.StatusOK, "loginaccount.html", gin.H{
				"text": logintext,
			})
		} else {
			logintext = "ログイン成功"
			// session.Set("loginUser", loginEmail)
			// session.Save()
			// log.Print(session.Get("loginUser"))
			u := db.GetUser(loginEmail)
			loginName := u.UserName
			c.SetCookie("loginUser", string(loginEmail), 3600, "/", "localhost", false, false)
			c.SetCookie("loginName", string(loginName), 3600, "/", "localhost", false, false)

			c.Redirect(http.StatusMovedPermanently, "/")
		}

	})

	// ログアウト(画面なし)
	router.GET("/logout", func(c *gin.Context) {
		// session := sessions.Default(c)

		// session.Delete("loginUser")
		// session.Save()
		c.SetCookie("loginUser", "", -1, "/", "localhost", true, true)
		c.SetCookie("loginName", "", -1, "/", "localhost", true, true)

		log.Print("logout")
		// log.Print(session.Get("loginUser"))

		// ログアウト処理してトップに戻る
		c.HTML(http.StatusOK, "logout.html", nil)

	})

	// email認証画面
	router.GET("/emailcertification", func(c *gin.Context) {
		c.HTML(http.StatusOK, "passwordcertification.html", nil)
	})

	// パスワード設定完了画面
	router.GET("/editpasswordcomplete", func(c *gin.Context) {
		c.HTML(http.StatusOK, "repasswordcomplete.html", nil)
	})

	// 会員管理ページ
	router.GET("/accountinfo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "usermanagement.html", nil)
	})

	router.Run(":3000")
}

func myMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		header := context.GetHeader("X-REQUEST-ID")
		log.Print("RequestId=" + header)
		context.Next()
	}
}
