package main

import (
	"context"
	"fmt"

	"github.com/leaper-one/2SOMEone/service"
	"github.com/leaper-one/2SOMEone/util"

	pb "github.com/leaper-one/2someone-proto/gen/grpc/bubble-box/note/golang"
)

const (
	SUCCESS = 200
	FAIL    = 500
)

var (
	config      = loadConfig("./config.yaml")
	dbc         = util.OpenDB("./user.db")
	noteService = service.NewNoteService(dbc)
)

type NoteService struct {
	pb.UnimplementedNoteServiceServer
}

/*
 * 根据 note_id 访问投稿
 * jwt needed in metadata
 */
func (n *NoteService) GetNotes(ctx context.Context, in *pb.GetNotesRequest) (*pb.GetNotesResponse, error) {
	fmt.Printf("GetNotes: %v\n", in)
	user_id, err := util.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	notes, count, err := noteService.GetNotes(ctx, in.Offset, in.Limit, in.IsSender, user_id)
	if err != nil {
		return &pb.GetNotesResponse{Code: FAIL, Msg: err.Error()}, err
	}

	var pbNotes []*pb.Note
	for k, v := range notes {
		pbNotes[k].CreatedAt = v.CreatedAt.GoString()
		pbNotes[k].UpdatedAt = v.UpdatedAt.GoString()
		pbNotes[k].NoteId = v.NoteID
		pbNotes[k].Title = v.Title
		pbNotes[k].Content = v.Content
	}
	return &pb.GetNotesResponse{Code: SUCCESS, Msg: "success.", Count: count, Notes: pbNotes}, nil
}
