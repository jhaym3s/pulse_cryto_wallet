import (
	"github.com/gin-gonic/gin"
	"log"
	"internal/db"
    "internal/handler"
)

func main() {
    db.InitDB()
    router := gin.Default()
    router.POST("/create-wallet", handler.CreateWallet)
    log.Fatal(router.Run(":8080"))
} 