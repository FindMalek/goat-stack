package handler

import (
	"context"
	"fmt"
	"goat/internal/db"
	goatv1 "goat/proto/goat/v1"
	"goat/proto/goat/v1/goatv1connect"
	"time"

	"connectrpc.com/connect"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type goatHandler struct {
	db *sqlx.DB
}

type VoteCount struct {
	Count int64 `db:"count"`
}

func NewGoatHandler() *goatHandler {
	return &goatHandler{
		db: db.New(),
	}
}

func (h *goatHandler) Vote(ctx context.Context, req *connect.Request[goatv1.VoteRequest]) (*connect.Response[goatv1.VoteResponse], error) {

	tx := h.db.MustBegin()

	var value string
	switch req.Msg.Vote {
	case goatv1.Vote_YES:
		value = "yes"
		break
	case goatv1.Vote_NO:
		value = "no"
		break

	default:
		break
	}
	r := tx.MustExec("INSERT INTO vote (timestamp, vote) VALUES ($1, $2)", time.Now().Unix(), value)
	fmt.Println(r.LastInsertId())
	tx.Commit()
	res := connect.NewResponse(&goatv1.VoteResponse{
		Success: true,
	})

	return res, nil
}

func (h *goatHandler) GetVotes(context.Context, *connect.Request[goatv1.GetVotesRequest]) (*connect.Response[goatv1.GetVotesResponse], error) {

	totalCount := VoteCount{}
	error := h.db.Get(&totalCount, "SELECT count(*) as count FROM vote")
	if error != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get votes"))
	}
	yesCount := VoteCount{}
	error = h.db.Get(&yesCount, "SELECT count(*) as count FROM vote where vote = 'yes'")
	if error != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get votes"))
	}
	fmt.Println("yes", yesCount.Count)
	noCount := VoteCount{}
	error = h.db.Get(&noCount, "SELECT count(*) as count FROM vote where vote = 'no'")
	if error != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get votes"))
	}
	maybeCount := VoteCount{}
	error = h.db.Get(&maybeCount, "SELECT count(*) as count FROM vote where vote = 'maybe'")
	if error != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get votes"))
	}

	res := connect.NewResponse(&goatv1.GetVotesResponse{
		Yes:   yesCount.Count,
		No:    noCount.Count,
	})
	return res, nil
}

func RegisterConnect(r *chi.Mux) {
	goatServer := NewGoatHandler()

	path, handler := goatv1connect.NewGoatServiceHandler(goatServer)

	r.Group(func(r chi.Router) {
		r.Mount(path, handler)
	})
}
