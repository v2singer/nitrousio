package v1

import (
	"context"
	"database/sql"
	"fmt"
	v1 "go_grpc/api/proto/v1"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

type toDoServiceServer struct {
	db *sql.DB
	v1.UnsafeToDoServiceServer
}

func NewToDoServiceServer(db *sql.DB) v1.ToDoServiceServer {
	return &toDoServiceServer{db: db}
}

func (s *toDoServiceServer) mustEmbedUnimplementedToDoServiceServer() {}

func (s *toDoServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Error(codes.Unimplemented, "unsupported API version: "+apiVersion)
		}
	}
	return nil
}

func (s *toDoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "connect database failed: "+err.Error())
	}
	return c, nil
}

func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := s.connect(ctx)

	if err != nil {
		return nil, err
	}

	defer c.Close()

	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "args error: "+err.Error())
	}
	// TODO: SQLi error, use gorm pl
	sqlStr := "INSERT INTO ToDo(`Title`, `Description`, `Reminder`) VALUES(?,?,?)"
	res, err := c.ExecContext(ctx, sqlStr, req.ToDo.Title, req.ToDo.Description, reminder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "add todo failed: "+err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "get newest id failed: "+err.Error())

	}
	return &v1.CreateResponse{Api: req.Api, Id: id}, nil
}

func (s *toDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := s.connect(ctx)

	if err != nil {
		return nil, err
	}
	defer c.Close()

	// TODO: SQLi error
	sqlStr := "SELECT `ID`, `Title`, `Description` FROM ToDo WHERE `ID`=?"
	rows, err := c.QueryContext(ctx, sqlStr, req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "query failed: "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to get data: "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ID='%d' not found", req.Id))
	}
	var td v1.ToDo
	if err := rows.Scan(&td.Id, &td.Title, &td.Description); err != nil {
		return nil, status.Error(codes.Unknown, "scan data failed: "+err.Error())
	}
	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("multi data for ID='%d'", req.Id))
	}
	return &v1.ReadResponse{Api: req.Api, ToDo: &td}, nil
}

func (s *toDoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder arg error")

	}
	// TODO: SQLi
	sqlStr := "UPDATE ToDo SET `Title`=?, `Description`=?, `Reminder`=? WHERE `ID`=?"
	res, err := c.ExecContext(ctx, sqlStr,
		req.ToDo.Title,
		req.ToDo.Description, reminder, req.ToDo.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "update failed: "+err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed affected line update error: "+err.Error())
	}
	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ID=%d not found", req.ToDo.Id))
	}
	return &v1.UpdateResponse{Api: req.Api, Updated: req.ToDo.Id}, nil
}

func (s *toDoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	sqlStr := "DELETE FROM ToDo where `ID`=?"
	res, err := c.ExecContext(ctx, sqlStr, req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "delete failed: "+err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update error: "+err.Error())
	}
	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ID=%d not found", req.Id))
	}
	return &v1.DeleteResponse{Api: req.Api, Deleted: req.Id}, nil
}

func (s *toDoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	c, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	sqlStr := "SELECT `ID`, `TiTle`, `Descrption` FROM ToDo"
	rows, err := c.QueryContext(ctx, sqlStr)
	if err != nil {
		return nil, status.Error(codes.Unknown, "query failed: "+err.Error())
	}
	defer rows.Close()

	todoList := []*v1.ToDo{}
	for rows.Next() {
		td := new(v1.ToDo)
		if err := rows.Scan(&td.Id, &td.Title, &td.Description); err != nil {
			return nil, status.Error(codes.Unknown, "parse all failed: "+err.Error())
		}
		todoList = append(todoList, td)
	}
	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "query all failed "+err.Error())
	}
	return &v1.ReadAllResponse{Api: req.Api, ToDos: todoList}, nil
}
