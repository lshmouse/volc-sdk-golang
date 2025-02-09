// Code generated by protoc-gen-volcengine-sdk
// source: DenyConfigService
// DO NOT EDIT!

package live

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/live"
	"github.com/volcengine/volc-sdk-golang/service/live/models/request"
)

func Test_UpdateDenyConfig(t *testing.T) {
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.UpdateDenyConfigRequest{
		Vhost:  "your Vhost",
		Domain: "your Domain",
		App:    "your App",
	}

	resp, status, err := instance.UpdateDenyConfig(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DescribeDenyConfig(t *testing.T) {
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.DescribeDenyConfigRequest{
		Vhost:  "your Vhost",
		Domain: "your Domain",
		App:    "your App",
	}

	resp, status, err := instance.DescribeDenyConfig(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
