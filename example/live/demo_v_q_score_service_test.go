// Code generated by protoc-gen-volcengine-sdk
// source: VQScoreService
// DO NOT EDIT!

package live

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/live"
	"github.com/volcengine/volc-sdk-golang/service/live/models/request"
)

func Test_CreateVQScoreTask(t *testing.T) {
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.CreateVQScoreTaskRequest{
		MainAddr:      "your MainAddr",
		ContrastAddr:  "your ContrastAddr",
		FrameInterval: 0,
		Duration:      0,
		Algorithm:     "your Algorithm",
	}

	resp, status, err := instance.CreateVQScoreTask(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DescribeVQScoreTask(t *testing.T) {
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.DescribeVQScoreTaskRequest{
		ID: "your ID",
	}

	resp, status, err := instance.DescribeVQScoreTask(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListVQScoreTask(t *testing.T) {
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.ListVQScoreTaskRequest{
		StartTime: "your StartTime",
		EndTime:   "your EndTime",
		PageNum:   0,
		PageSize:  0,
		Status:    0,
	}

	resp, status, err := instance.ListVQScoreTask(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
