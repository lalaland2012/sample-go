package router

import (
	"net/http"
	"sample/app/hello"
	"sample/app/infrastructure"
	"sample/app/login"
	"sample/app/logout"
	"sample/app/profile"
	"sample/app/register"
	"sample/app/shared/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router is application struct hold Mux and db connection
type Router struct {
	Mux                *chi.Mux
	SQLHandler         *infrastructure.SQL
	CacheHandler       *infrastructure.Cache
	LoggerHandler      *infrastructure.Logger
	TranslationHandler *infrastructure.Translation
}

// InitializeRouter initializes Mux and middleware
func (r *Router) InitializeRouter() {
	r.Mux.Use(middleware.RequestID)
	r.Mux.Use(middleware.RealIP)
	// Custom middleware(Translation)
	// r.Mux.Use(r.TranslationHandler.Middleware.Middleware)
	// // Custom middleware(Logger)
	// r.Mux.Use(mMiddleware.Logger(r.LoggerHandler))

}

// SetupHandler set database and redis and usecase.
func (r *Router) SetupHandler() {
	// error handler set.
	eh := handler.NewHTTPErrorHandler(r.LoggerHandler.Log)
	r.Mux.NotFound(eh.StatusNotFound)
	r.Mux.MethodNotAllowed(eh.StatusMethodNotAllowed)

	r.Mux.Method(http.MethodGet, "/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Mux.HandleFunc("/terms-of-use", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/terms-of-use.html")
	})

	ah := handler.NewApplicationHTTPHandler(r.LoggerHandler.Log)

	// base set.
	// bh := handler.NewBaseHTTPHandler(r.LoggerHandler.Log)
	// // base set.
	// br := repository.NewBaseRepository(r.LoggerHandler.Log)
	// // base set.
	// bu := usecase.NewBaseUsecase(r.LoggerHandler.Log)

	// uh := user.NewHTTPHandler(br, bu, bh, r.SQLHandler, r.CacheHandler)

	hw := hello.NewHTTPHandler(ah)
	lg := login.NewLoginHTTPHandler(ah)
	lo := logout.NewLogoutHTTPHandler(ah)
	rg := register.NewRegisterHTTPHandler(ah)
	pf := profile.NewProfileHTTPHandler(ah)

	r.Mux.Route("/", func(cr chi.Router) {
		cr.Get("/hello", hw.HelloWorld)
		cr.Get("/login", lg.Login)
		cr.Post("/login", lg.HandleLogin)
		cr.Post("/logout", lo.HandleLogout)
		cr.Get("/register", rg.Register)
		cr.Post("/register", rg.HandleRegister)
		cr.Get("/profile", pf.ShowProfile)
		cr.Post("/profile",pf.HandleProfile)
	})
}
