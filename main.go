package main
import(
	"fmt"
	"log"
	"github.com/Santy13226/connectiongorm.git"
)

func main() {
	conn, err := GetConnection("DESKTOP-QV1GQ7E", "sa", "200018S@nty", "chatbot", "1433")
	errorFatal(err)
	defer conn.Close()

	fmt.Println("Conexión exitosa")
}
func errorFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}