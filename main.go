package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/bmizerany/aws4"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	r, _ := http.NewRequest("GET", "http://kls.cn-beijing-6.api.ksyun.com?Action=ListStreamDurations&Version=2017-01-01&UniqueName=maxi&App=live&Pubdomain=live.moxiulive.com&StartUnixTime=1492617600&EndUnixTime=1492704000", nil)

	now := time.Now().UTC().Format("2006-01-02 15:04:05")
	amzDate := now[0:4] + now[5:7] + now[8:10] + "T" + now[11:13] + now[14:16] + now[17:19] + "Z"
	r.Header.Set("X-Amz-Date", amzDate)
	
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
