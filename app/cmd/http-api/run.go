package httpapi

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var serverID string

func init() {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(800)))
	serverID = uuid.New().String()
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response ID %s %s", serverID, time.Now().Format(time.RFC3339))
}

func handleAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Admin Response ID %s %s", serverID, time.Now().Format(time.RFC3339))
}

func Run(port int, logger *log.Logger) {
	http.HandleFunc("/index", handleIndex)
	http.HandleFunc("/admin", handleAdmin)

	logger.Printf("http-api runnig successfully - port %d\n", port)
	http.ListenAndServe(fmt.Sprint(":", port), nil)
}
