package client

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/cloud-hashing-inspire/pkg/message/const"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-inspire"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get inspire connection: %v", err)
	}
	defer conn.Close()

	cli := npool.NewCloudHashingInspireClient(conn)

	return fn(_ctx, cli)
}

func CreateInvitationCode(ctx context.Context, in *npool.UserInvitationCode) (*npool.UserInvitationCode, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingInspireClient) (cruder.Any, error) {
		resp, err := cli.CreateUserInvitationCode(ctx, &npool.CreateUserInvitationCodeRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create user invitation code: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create user invitation code: %v", err)
	}
	return info.(*npool.UserInvitationCode), nil
}
