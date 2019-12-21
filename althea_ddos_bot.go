package main
import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func request(url string) string{
	resp, err := http.Get(url)
	if err != nil {log.Fatalln(err)}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{log.Fatalln(err)}
	return string(body)
}
func main(){for {println(request(os.Args[1]))}}