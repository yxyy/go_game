package corn

import (
	"io/ioutil"
	"net/http"
)

type Corn interface {
	Run()
}

func Run(corn Corn)  {
	corn.Run()
}

func GetHtml(url string) ([]byte,error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		return nil,err
	}

	return bytes,nil
}