package main

import (
        "encoding/base64"
        "fmt"
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

        request := iai.NewDetectFaceRequest()
        imgfile, _ := ioutil.ReadFile("img/wanglin/wl-001.jpeg")
        imgbody := base64.StdEncoding.EncodeToString(imgfile)
        params := `{"Image":"` + imgbody + `","NeedFaceAttributes":1,"NeedQualityDetection":1}`
        err := request.FromJsonString(params)
        if err != nil {
                panic(err)
        }
        response, err := client.DetectFace(request)
        if _, ok := err.(*errors.TencentCloudSDKError); ok {
                fmt.Printf("An API error has returned: %s", err)
                return
        }
        if err != nil {
                panic(err)
        }
        fmt.Printf("%s", response.ToJsonString())
}