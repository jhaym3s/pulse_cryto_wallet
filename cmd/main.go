

func main() {
    db.InitDB()
    router := gin.Default()
    router.POST("/create-wallet", handler.CreateWallet)
    log.Fatal(router.Run(":8080"))
} 