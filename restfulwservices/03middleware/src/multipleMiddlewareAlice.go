import (
	"encoding/json"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	chain := alice.New(filterContentType, setServerTimeCookie).Then(mainLogicHandler)
	http.Handle("/city", chain)
	http.ListenAndServe(":8000", nil)
}
