package main

func main() {
	dependencies := loadDependancies()
  LoadMiddleware(dependencies.gin, dependencies.db)
  LoadRouter(dependencies.gin)
  dependencies.gin.Run()
}
