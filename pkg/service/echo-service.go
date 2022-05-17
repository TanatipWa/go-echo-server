package service

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/TanatipWa/go-echo-server/pkg/proto/kbtg/kbtg1000"
)

type echoServiceServerImpl struct {
	// put initial database connection here
}

func NewEchoServiceServer() pb.EchoServiceServer {
	return nil
}

func (i *echoServiceServerImpl) Query(ctx context.Context, echo *pb.Echo) (*pb.EchoList, error) {
	fmt.Printf("Request :%v\n", echo)
	if echo == nil {
		return nil, nil
	}
	// test return error if echo name is error
	if echo.Name == "error" {
		return nil, errors.New("return error just for test")
	}
	// append request in echoList and return back
	var echoList []*pb.Echo
	echoList = append(echoList, &pb.Echo{
		Id:          echo.Id,
		Name:        echo.Name,
		Description: echo.Description,
		Status:      echo.Status,
		CreateDate:  echo.CreateDate,
		UpdateDate:  echo.UpdateDate,
	})
	response := &pb.EchoList{
		Api:      "GetEchoAPI",
		EchoList: echoList,
	}
	
	return response, nil
}
