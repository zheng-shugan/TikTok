package rpc

import (
	"context"
	"fmt"
	config "github.com/sunflower10086/TikTok/http/config"
	___interaction "github.com/sunflower10086/TikTok/interaction/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var interactionClient ___interaction.InteractionClient

func Init() {
	relationConf := config.C().Apps.Relation
	if relationConf == nil {
		log.Println("relationConf1为空")
	} else {
		log.Println("relationConf1非空")
	}
	Addr := fmt.Sprintf("%s:%s", relationConf.Host, relationConf.Port)
	conn, err := grpc.Dial(Addr, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("rpc连接失败: ", err)
		return
	} else {
		log.Println("rpc连接成功")
	}

	interactionClient = ___interaction.NewInteractionClient(conn)

	if interactionClient != nil {
		log.Println("interactionClient非空")
	} else {
		log.Println("interactionClient为空")
	}
}

func FavoriteAction(ctx context.Context, req *___interaction.FavoriteActionReq, opts ...grpc.CallOption) error {
	_, err := interactionClient.FavoriteAction(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func FavoriteList(ctx context.Context, req *___interaction.FavoriteListReq, opts ...grpc.CallOption) (*___interaction.FavoriteListResp, error) {
	resp, err := interactionClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CommentAction(ctx context.Context, req *___interaction.CommentActionReq, opts ...grpc.CallOption) (*___interaction.CommentActionResp, error) {
	resp, err := interactionClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CommentList(ctx context.Context, req *___interaction.CommentListReq, opts ...grpc.CallOption) (*___interaction.CommentListResp, error) {
	resp, err := interactionClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
