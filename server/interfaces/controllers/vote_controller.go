package controllers

import (
	"github.com/kindaidensan/UMR/proto/vote"
	"google.golang.org/grpc"
	"context"
	"log"
)

type VoteConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type VoteController struct {
	client vote.VoteClient
}

func NewVoteController(config VoteConfig) (*VoteController, error) {
	conn, err := grpc.Dial(config.Host+":"+config.Port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := vote.NewVoteClient(conn)
	return &VoteController {
		client: client,
	}, nil
}

func (controller *VoteController) Create(c Context) {
	createRequest := vote.CreateRequest{}
	c.Bind(&createRequest)
	userid, exist := c.Get("userid")
	if !exist {
		c.JSON(500, NewMsg("ユーザーIDエラー"))
		return
	}
	createRequest.Owner = userid.(string)
	res, err := controller.client.Create(context.TODO(), &createRequest)
	if err != nil {
		c.JSON(500, NewMsg(err.Error()))
		return
	}
	c.JSON(200, res)
}

func (controller *VoteController) Vote(c Context) {
	voteRequest := vote.VoteRequest{}
	c.Bind(&voteRequest)
	userid, exist := c.Get("userid")
	if !exist {
		c.JSON(500, NewMsg("ユーザーIDエラー"))
		return
	}
	voteRequest.Userid = userid.(string)
	log.Print(voteRequest.Id)	
	log.Print(voteRequest.Userid)	
	log.Print(voteRequest.Agree)	
	res, err := controller.client.Vote(context.TODO(), &voteRequest)
	if err != nil {
		c.JSON(500, NewMsg(err.Error()))
		return
	}
	c.JSON(200, res)	
}

func (controller *VoteController) Get(c Context) {
	userid, exist := c.Get("userid")
	if !exist {
		c.JSON(500, NewMsg("ユーザーIDエラー"))
		return
	}
	getRequest := vote.GetRequest {
		Userid: userid.(string),
	}
	res, err := controller.client.Get(context.TODO(), &getRequest)
	if err != nil {
		c.JSON(500, NewMsg(err.Error()))
		return
	}
	c.JSON(200, res)
}
