package routes

import (
	"html/template"
	"mhilmi999/project-2-mhilmi999/controllers"

	. "mhilmi999/project-2-mhilmi999/helpers"
	"mhilmi999/project-2-mhilmi999/middlewares"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

func Init() *echo.Echo {
	e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Ini percobaan Echo nya dah jalan belum")
	// }) // Mendefinisikan routing awal akses nya

	// Coba template lain (tampilan html)
	//---------------------------------------------//
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.Static("assets"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("UIDToken"))))
	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseFiles("view/header.html", "view/home.html", "view/footer.html"))
	templates["login.html"] = template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseFiles("view/login.html"))
	templates["register.html"] = template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseFiles("view/register.html"))
	templates["dashboard.html"] = template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseFiles("view/header.html","view/dashboard.html", "view/footer.html"))
	e.Renderer = &DaftarTemplate{
		Templates: templates,
	}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &middlewares.CustomContext{c}
			return next(cc)
		}
	})

	e.GET("/", controllers.HomeView)
	e.GET("/login", controllers.LoginView)
	// Melakukan authorization user
	e.POST("/login", controllers.CheckingLogin)
	
	e.GET("/register", controllers.RegisterView)
	e.POST("/register", controllers.NewRegister)

	e.GET("/dashboard", controllers.DashboardView)
	
	//---------------------------------------------//
	// Coba template lain (tampilan html)

	//--------------------------------------------------------------------//


	// API DOCS with Swagger
	e.GET("/documentation/*", echoSwagger.WrapHandler)

	// RESTFUL API nya
	// Pemanggilan GET Methods pada routes author dari database
	e.GET("/author", controllers.FetchAllAuthor)
	
	// Pemanggilan Post Methods pada routes author dari database
	e.POST("/author", controllers.StoreAuthor, middlewares.IsAuth)
	
	// Pemanggilan Put Methods pada routes author untuk ke database
	e.PUT("/author", controllers.UpdateAuthor, middlewares.IsAuth)
	
	// Pemanggilan Delete Methods pada routes author untuk ke database
	e.DELETE("/author", controllers.HapusAuthor, middlewares.IsAuth)
	
	// Melakukan validasi user secara sederhana (login) dengan metode bcrypt
	e.GET("/make-hash/:password", controllers.GenerateHashPassword)
	//--------------------------------------------------------------------//


	return e
}
