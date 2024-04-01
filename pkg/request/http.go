package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Send(url string, method string, token string) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
