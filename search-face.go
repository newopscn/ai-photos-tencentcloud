package main

import (
        "fmt"
	    "encoding/base64"
        "io/ioutil"
        "os"

        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
        iai "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iai/v20180301"
)

func main() {

        credential := common.NewCredential(
                os.Getenv("TX_SECRET_ID"),
                os.Getenv("TX_SECRET_KEY"),
        )
        cpf := profile.NewClientProfile()
        cpf.HttpProfile.Endpoint = "iai.tencentcloudapi.com"
        client, _ := iai.NewClient(credential, "ap-chengdu", cpf)
        
        request := iai.NewSearchFacesRequest()

        imgfile, _ := ioutil.ReadFile("img/hujuntao/hjt-001.jpeg")
        imgbody := base64.StdEncoding.EncodeToString(imgfile)
        params := `{"GroupIds":["lovehome"],"Image":"` + imgbody + `","MaxPersonNum":1}`
        //fmt.Printf("%s", params)
        err := request.FromJsonString(params)
        if err != nil {
                panic(err)
        }
        response, err := client.SearchFaces(request)
        if _, ok := err.(*errors.TencentCloudSDKError); ok {
                fmt.Printf("An API error has returned: %s", err)
                return
        }
        if err != nil {
                panic(err)
        }
        fmt.Printf("%s", response.ToJsonString())
} 
