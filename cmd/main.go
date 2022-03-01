package main

func main() {
	dependencies := loadDependancies()
	dependencies.gin.Run()
}
