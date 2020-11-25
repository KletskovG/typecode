package main
import (
	"log"
	"github.com/kletskovg/typecode/server/src/internal/app/apiserver"
	"os"
)

func main() {
	log.Printf(os.Getenv("DB_CONNECTION"))
	config := apiserver.NewConfig()
	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
