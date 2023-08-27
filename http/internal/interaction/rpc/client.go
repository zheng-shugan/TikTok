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

func init() {
	relationConf := config.C().Apps.Relation
	Addr := fmt.Sprintf("%s:%s", relationConf.Host, relationConf.Port)
	conn, err := grpc.Dial(Addr, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return
	}

	interactionClient = ___interaction.NewInteractionClient(conn)
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
