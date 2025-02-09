// Code generated by protoc-gen-volcengine-sdk
// source: PullToPushTaskService
// DO NOT EDIT!

package live

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/live"
	"github.com/volcengine/volc-sdk-golang/service/live/models/request"
)

func Test_CreatePullToPushTask(t *testing.T) {
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.CreatePullToPushTaskRequest{
		Title:       "your Title",
		StartTime:   0,
		EndTime:     0,
		CallbackURL: "your CallbackURL",
		Type:        0,
		CycleMode:   0,
		DstAddr:     "your DstAddr",
		SrcAddr:     "your SrcAddr",
		SrcAddrS:    []string{"your SrcAddrS"},
	}

	resp, status, err := instance.CreatePullToPushTask(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListPullToPushTask(t *testing.T) {
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.ListPullToPushTaskRequest{
		Page:  0,
		Size:  0,
		Title: "your Title",
	}

	resp, status, err := instance.ListPullToPushTask(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdatePullToPushTask(t *testing.T) {
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.UpdatePullToPushTaskRequest{
		Title:       "your Title",
		TaskId:      "your TaskId",
		StartTime:   0,
		EndTime:     0,
		CallbackURL: "your CallbackURL",
		Type:        0,
		CycleMode:   0,
		DstAddr:     "your DstAddr",
		SrcAddr:     "your SrcAddr",
		SrcAddrS:    []string{"your SrcAddrS"},
	}

	resp, status, err := instance.UpdatePullToPushTask(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_StopPullToPushTask(t *testing.T) {
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.StopPullToPushTaskRequest{
		TaskId: "your TaskId",
	}

	resp, status, err := instance.StopPullToPushTask(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_RestartPullToPushTask(t *testing.T) {
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.RestartPullToPushTaskRequest{
		TaskId: "your TaskId",
	}

	resp, status, err := instance.RestartPullToPushTask(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DeletePullToPushTask(t *testing.T) {
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.DeletePullToPushTaskRequest{
		TaskId: "your TaskId",
	}

	resp, status, err := instance.DeletePullToPushTask(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
