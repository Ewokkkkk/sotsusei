package main

import (
	"io"
	"net/http"
	"os"
	"sotsusei/data"
	"sotsusei/db"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func main() {

	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)
	log.SetOutput(gin.DefaultWriter)

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "assets")

	// トップページ
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 商品ページ
	router.GET("/product", func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.Redirect(http.StatusMovedPermanently, "/")
		}
		product := db.GetProduct(id)
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
		var newUser data.NewUser
		err := c.Bind(&newUser)
		if err != nil {
			c.Status(http.StatusBadRequest)
		} else {
			// 会員登録処理
			log.Printf("%+v\n", newUser)
			c.HTML(http.StatusOK, "accountcomplete.html", newUser)
			if err := db.SignUp(newUser); err != nil {
				c.Status(http.StatusBadRequest)
			}
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

	router.POST("/login", func(ctx *gin.Context) {
		logintext := ""
		formPassword := ctx.PostForm("password")
		dbPassword := db.GetUser(ctx.PostForm("email")).Password
		log.Print(dbPassword)

		if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(formPassword)); err != nil {
			logintext = "ログイン失敗"
		} else {
			logintext = "ログイン成功"
		}

		ctx.HTML(http.StatusOK, "loginaccount.html", gin.H{
			"text": logintext,
		})
	})

	// ログアウト(画面なし)
	router.GET("/logout", func(c *gin.Context) {

		// ログアウト処理してトップに戻る
		c.Redirect(http.StatusMovedPermanently, "/")
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
