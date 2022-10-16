package fetcher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	url = strings.Replace(url, "http://", "https://", 1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	// req.Header.Set("cookie", "sid=c6563e3f-de2a-46ff-8a6b-e21637f0bfe2; ec=ByYNpWEp-1664770533078-546fa9766ad9d-1406334761; FSSBBIl1UgzbN7NO=5UrGi7lb0b_qRwrMrnJLxctzOXTWMlz9jEHxDi4nqDbSCDENnylQX1eG.V4YdWz76KgrVOx9n1IXsfutRLNXZdG; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1664770582,1664856374; hcbm_tk=MjZvkSuzE8En8AWeqsdF3nQ3yXUG1nIwNk+4wuPcMKTiFN0oGKABhT3w8caAxU9D0c9/KmuoBPvaXpIxKkye78OeplPgi5+4HBl6HSkBBO1vO4cyNolYelpYji9NAXz2otN2fjKK77TscE4s5mt7L27r1QxOIepMEhXvYDiCCGEElymWW7R2/vnIk3KLbmGkCY0hQumPk5FTlI4yeawODwVTXh0GUJo/OydLuScdvCESbMS7kg6+wVeyj6CoTT+7Lggz75qVfLEDSnfszV3oQuqb9HSM+QmF5065nUJULdd+HD9zWqekwPDpW72aonLORdLz7g5whq38UWVgF5HR+WqftHFNXGQzWE1ZRwe43Xvc47axwaV80YeBzLitn7z/EzFwwIiP75a9ODlv5//nNkxRjI67qFRbKXG3KgvR9Q==_RERXV4AmIa4VjY7e; _exid=JipazBJP/PYcq8u1QVZD1eKDwsq109VJriY63W4rDDM3pkE9I6Ha1JGdV1nLCaDyqkhKcRR29YOYHjU68vRzjA==; _efmdata=NhwqfSGacsvHV6lmW4O+iw+ZVYnzRbjk+lHJThjDon4fvR+CKtSXaue/dkPPq2madN11PSOUKpUcKZGCnj2CowZDa73VpA8Ts7/d8cTmK58=; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1665275318; FSSBBIl1UgzbN7NP=530EpODJcl3AqqqDkXPMicaLS5_Rha7_KJ1ZftMImK2xivxdAI1fkYdRaznr.nwY1QNuNV7OGpdx1REic8mCGLORDbpFZS.rconudgBYEulxhMlCMR1klHAHrDC96QMMFTLjCFPeVk8mPkcNFitG2FXArjJxT.Dmwprnw5k4sUyN6f39tkYqD_WZ3sDdiB9ZA7RZSeyTocPpYOdAlOaQ91WbKmCJT0FhPNo7AZk3D0oSFsqTBAPtUlKOQb0WLmv3WAYjA1OGFcJFxDXo8JLRXIbFBqZ_aPUvi0eJGij3p2IoVw7INbSELUjHR5v2gLU9Cq")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.StatusCode)
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
