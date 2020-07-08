package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong  status code: %d\n", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

// func Fetch(url string) ([]byte, error) {
// 	client := &http.Client{}
// 	newUrl := strings.Replace(url, "http://", "https://", 1)
// 	req, err := http.NewRequest("GET", newUrl, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")
// 	cookie1 := "sid=c5243abb-4aff-44e3-ba91-7d688332d517; FSSBBIl1UgzbN7NO=5uvWSdF68TJWuRViTFSaJGNWKt3bcTRIg3aGx5H4XQ_HCYvm6pBBXvP9rws_jLy.P3D29rTVxv2kEd9M6x58NSG; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1594018083,1594044360,1594110087,1594111025; ec=qxpgJ2LR-1594147442756-17157044fe7a31427679333; _efmdata=zyEXXHfgFMTGV1KUcSuytot9sckzAbDl6AZYujKO%2FwUeixHOb0ULZ%2FoeIOdbVQBm0%2F5SyWQL85bCs7H%2F9Mxod7qZmWvNoAa2azJk5V9AC0Y%3D; _exid=ws63t9DSqFHPeuZ1fM2LFiJMAezwaS0KU8p3BhCmO%2B3OKbriZrc%2FfiTnrWw68cga%2FPYzDiAD8NxVcwI3GHqzUQ%3D%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1594149050; FSSBBIl1UgzbN7NP=5UasEX2Zlnh9qqqmut.xfzG2Gk5WvbmPaPMkIOLIVndHutpw9DUuhhME5gJ5BOsQqeOOU0DSuCIzm0gKF.djRALpegny75sZk9oMfoWSkTlwf7LXM9uQnYzxJxRSa4dWNmWv7cEipw_1fEmsAJALVdRZ6JXagNxcNNDJhRfYV0I4mBqjH1P6KWskYuF4hnLl0EAzEShJlxApw2PmvCVlPqxiebROxDZMdBGZZlvNygX72HdNbwXG.KHMT_1WC4YTKbc0WeaJmQKVcfuLakc4D21yyyohFVG91nYfZj6IRfj7.iYZ_7WCC3GyLEOI8XgLoQ"
// 	req.Header.Add("cookie", cookie1)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
// 	}
// 	return ioutil.ReadAll(resp.Body)
// }
