package service

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/TanatipWa/go-echo-server/pkg/proto/kbtg/kbtg1000"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_echoServiceServerImpl_Query(t *testing.T) {
	timestampRun := timestamppb.Now()
	type args struct {
		ctx  context.Context
		echo *pb.Echo
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.EchoList
		wantErr bool
	}{
		// Add test here
		{
			"Nil case",
			args{context.Background(), nil},
			nil,
			false,
		},
		{
			"Error name case",
			args{context.Background(), &pb.Echo{Name: "error"}},
			nil,
			true,
		},
		{
			"Sample echo case",
			args{context.Background(), &pb.Echo{
				Id:          "id",
				Name:        "name",
				Description: "desc",
				Status:      "status",
				CreateDate:  timestampRun,
				UpdateDate:  timestampRun,
			}},
			&pb.EchoList{
				Api: "GetEchoAPI",
				EchoList: []*pb.Echo{&pb.Echo{
					Id:          "id",
					Name:        "name",
					Description: "desc",
					Status:      "status",
					CreateDate:  timestampRun,
					UpdateDate:  timestampRun,
				}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &echoServiceServerImpl{}
			got, err := i.Query(tt.args.ctx, tt.args.echo)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() got = %v, want %v", got, tt.want)
			}
		})
	}
}
