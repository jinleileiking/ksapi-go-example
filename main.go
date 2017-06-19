package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bmizerany/aws4"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	r, _ := http.NewRequest("GET", "http://kls.cn-beijing-6.api.ksyun.com?Action=ListStreamDurations&Version=2017-01-01&UniqueName=maxi&App=live&Pubdomain=live.moxiulive.com&StartUnixTime=1492617600&EndUnixTime=1492704000", nil)

	s := &aws4.Service{
		Name:   "kls",
		Region: "cn-beijing-6",
	}

	k := &aws4.Keys{
		AccessKey: "your ak",
		SecretKey: "your sk",
	}

	s.Sign(k, r)

	resp, err := aws4.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)

	/// 接口应答
	spew.Dump(responseString)

}
