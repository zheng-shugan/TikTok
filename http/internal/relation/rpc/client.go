package rpc

import (
	"context"
	"fmt"
	config "github.com/sunflower10086/TikTok/http/config"
	___relation "github.com/sunflower10086/TikTok/relation/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var relationClient ___relation.RelationClient

func Init() {
	relationConf := config.C().Apps.Relation
	if relationConf == nil {
		log.Println("relationConf2为空")
	} else {
		log.Println("relationConf2非空")
	}
	Addr := fmt.Sprintf("%s:%s", relationConf.Host, relationConf.Port)
	conn, err := grpc.Dial(Addr, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return
	}

	relationClient = ___relation.NewRelationClient(conn)
}

func RelationAction(ctx context.Context, request *___relation.ActionRequest) error {
	_, err := relationClient.Action(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func FollowList(ctx context.Context, request *___relation.FollowListRequest) (*___relation.FollowListResponse, error) {
	resp, err := relationClient.FollowList(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func FollowerList(ctx context.Context, request *___relation.FollowerListRequest) (*___relation.FollowerListResponse, error) {
	resp, err := relationClient.FollowerList(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func FriendList(ctx context.Context, request *___relation.FriendListRequest) (*___relation.FriendListResponse, error) {
	resp, err := relationClient.FriendList(ctx, request)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func MessageChat(ctx context.Context, request *___relation.MessageChatRequest) (*___relation.MessageChatResponse, error) {
	resp, err := relationClient.MessageChat(ctx, request)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func MessageAction(ctx context.Context, request *___relation.MessageActionRequest) (*___relation.MessageActionResponse, error) {
	resp, err := relationClient.MessageAction(ctx, request)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
