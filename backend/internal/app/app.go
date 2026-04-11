package app

func NewApp() *App {
	logger, _ := zap.NewProduction()
}

func (a *App) Run() error {

}

r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", handler.GetTasksHandler)
	})
	http.ListenAndServe(":3000", r)