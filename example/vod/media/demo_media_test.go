package media

import (
	"encoding/json"
	"fmt"
	"testing"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func TestVod_GetMediaInfos(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.volc/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	// instance.SetAccessKey("")
	// instance.SetSecretKey("")
	vids := "your vids"
	// Media Info
	query := &request.VodGetMediaInfosRequest{
		Vids: vids,
	}
	resp, code, err := instance.GetMediaInfos(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVod_GetRecommendedPoster(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.volc/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	// vod.NewInstance().SetAccessKey("your ak")
	// vod.NewInstance().SetSecretKey("your sk")
	vid := "your vid"
	query := &request.VodGetRecommendedPosterRequest{
		Vids: vid,
	}
	resp, code, err := instance.GetRecommendedPoster(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVod_UpdateMediaInfo(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.volc/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	// vod.NewInstance().SetAccessKey("your ak")
	// vod.NewInstance().SetSecretKey("your sk")
	vid := "your vid"
	title := "your title"
	query := &request.VodUpdateMediaInfoRequest{
		Vid:         vid,
		Title:       wrapperspb.String(title),
		PosterUri:   wrapperspb.String("PosterUri"),
		Description: wrapperspb.String("description"),
		Tags:        wrapperspb.String("tag1,tag2"),
	}
	resp, code, err := instance.UpdateMediaInfo(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVod_UpdateMediaPublishStatus(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.volc/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	// vod.NewInstance().SetAccessKey("your ak")
	// vod.NewInstance().SetSecretKey("your sk")
	vid := "your vid"
	query := &request.VodUpdateMediaPublishStatusRequest{
		Vid:    vid,
		Status: "Unpublished",
	}
	resp, code, err := instance.UpdateMediaPublishStatus(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVod_DeleteMedia(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.volc/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	// vod.NewInstance().SetAccessKey("your ak")
	// vod.NewInstance().SetSecretKey("your sk")
	vids := "your vids"
	callbackArgs := "callbackArgs"
	query := &request.VodDeleteMediaRequest{
		Vids:         vids,
		CallbackArgs: callbackArgs,
	}
	resp, code, err := instance.DeleteMedia(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVod_DeleteTranscodes(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.volc/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	// vod.NewInstance().SetAccessKey("your ak")
	// vod.NewInstance().SetSecretKey("your sk")
	vid := "your vid"
	fileIds := "your file ids"
	callbackArgs := "callbackArgs"
	query := &request.VodDeleteTranscodesRequest{
		Vid:          vid,
		FileIds:      fileIds,
		CallbackArgs: callbackArgs,
	}
	resp, code, err := instance.DeleteTranscodes(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

func TestVod_GetMediaList(t *testing.T) {
	instance := vod.NewInstance()
	// call below method if you dont set ak and sk in ～/.volc/config
	//vod.NewInstance().SetCredential(base.Credentials{
	//	AccessKeyID:     "your ak",
	//	SecretAccessKey: "your sk",
	//})

	// or set ak and ak as follow
	// instance.SetAccessKey("")
	// instance.SetSecretKey("")
	spaceName := "you space"
	vid := "your vid"
	status := "your status" // Published or Unpublished
	order := "your order"   // Desc or Asc
	tags := "your tags"
	startTime := "2021-01-01T00:00:00Z"
	endTime := "2021-04-01T00:00:00Z"
	offset := "0"
	pageSize := "10" // pageSize <= 100

	// Media Info
	query := &request.VodGetMediaListRequest{
		SpaceName: spaceName,
		Vid:       vid,
		Status:    status,
		Order:     order,
		Tags:      tags,
		StartTime: startTime,
		EndTime:   endTime,
		Offset:    offset,
		PageSize:  pageSize,
	}
	resp, code, err := instance.GetMediaList(query)
	fmt.Println(code)
	fmt.Println(err)
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}
