package fetcher

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Status code is not okay")
	}
	// when function ends, close the Closer.
	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
