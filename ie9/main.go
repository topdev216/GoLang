package main

import (
	"./src/handlers"
	"./src/middleware"
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"log"
)

func registHandler(g *gin.Engine) {
	home := g.Group("/")
	{
		home.GET("/", handlers.HomeHandler)
		home.GET("/html", handlers.HtmlHandler)
		home.GET("/param/:value", handlers.ParamHandler)
		home.GET("/query", handlers.QueryHandler)
		home.POST("/post", handlers.PostHandler)
		home.POST("/upload", handlers.Uploadhandler)
	}

	admin := g.Group("/admin", handlers.AuthRequired)
	{
		admin.GET("/login", handlers.AdminLogin)
		admin.GET("/logout", handlers.AdminLogout)
	}
}

/* Authentication middleware
func Authentication(con *sql.DB) gin.HandlerFunc {
    query := "SELECT user_id FROM users WHERE token = ?"
    return func(c *gin.Context) {
        tokenHeader := c.Request.Header.Get("Token")
        var user_id int
        err := conn.QueryRow(query, token).Scan(&user_id)
        if err != nil {
            // UNAUTHORISED ERROR ERROR ERROR
        }
        c.Set("UserID", user_id)
        c.Next()
    }
}
*/

/*
func databaseMiddleware(conn string) gin.HandlerFunc {
	db := sqlx.Connect("sqlite3", conn)
	if err := RunMigrations(db); err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
*/

func main() {
	log.Println("开始加载配置……")
	cfg, err := goconfig.LoadConfigFile("web.conf")
	if err != nil {
		panic(err)
	}

	debug, err := cfg.Bool("general", "debug")
	if err != nil {
		panic(err)
	}

	host, err := cfg.GetValue("general", "host")
	if err != nil {
		panic(err)
	}

	conn, err := cfg.GetValue("general", "conn")
	if err != nil {
		panic(err)
	}

	log.Println("开始初始化系统……")
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.LoadHTMLGlob("./views/*")
	g.Static("static", "./static")
	g.StaticFile("/favicon.ico", "./static/favicon.ico")

	log.Println("开始加载数据库……")
	g.Use(Database(conn))

	if debug {
		log.Println("开启 Debug 模式……")
		gin.SetMode(gin.DebugMode)
	}

	log.Println("开始注册路由……")
	registHandler(g)

	log.Println("开始启动服务……")
	g.Run(host)
}
