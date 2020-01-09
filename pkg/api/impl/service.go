package impl

import (
	"context"
	api "github.com/tianhanfangyan/grpc-app/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
}

func NewServer() api.StudentServiceServer {
	return &server{}
}

// 添加学生信息
func (s *server) AddStudent(ctx context.Context, args *api.AddStudentArgs) (*api.AddStudentReply, error) {

	if args.Stu.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "'id' not exists")
	}
	if args.Stu.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "'name' not exists")
	}
	if args.Stu.Sex == "" {
		return nil, status.Error(codes.InvalidArgument, "'sex' not exists")
	}
	if args.Stu.Age <= 0 {
		return nil, status.Error(codes.InvalidArgument, "'age' not exists or invalid age")
	}

	uid, err := s.addStudent(ctx, args.Stu)
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	return &api.AddStudentReply{Uid: uid}, nil
}

// 得到学生信息
func (s *server) GetStudent(ctx context.Context, args *api.GetStudentArgs) (*api.GetStudentReply, error) {

	if args.Uid == "" {
		return nil, status.Error(codes.InvalidArgument, "'uid' not exists")
	}

	stu, err := s.getStudent(ctx, args.Uid)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &api.GetStudentReply{Stu: stu}, nil
}

// 得到所有学生信息
func (s *server) GetAllStudent(ctx context.Context, args *api.GetAllStudentArgs) (*api.GetAllStudentReply, error) {

	stus, err := s.getAllStudent(ctx)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &api.GetAllStudentReply{Stus: stus}, nil
}

// 更新学生信息
func (s *server) UpdateStudent(ctx context.Context, args *api.UpdateStudentArgs) (*api.UpdateStudentReply, error) {

	if args.Uid == "" {
		return nil, status.Error(codes.InvalidArgument, "'uid' not exists")
	}
	if args.Stu.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "'id' not exists")
	}
	if args.Stu.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "'name' not exists")
	}
	if args.Stu.Sex == "" {
		return nil, status.Error(codes.InvalidArgument, "'sex' not exists")
	}
	if args.Stu.Age <= 0 {
		return nil, status.Error(codes.InvalidArgument, "'age' not exists or invalid age")
	}

	stu, err := s.updateStudent(ctx, args.Uid, args.Stu)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.UpdateStudentReply{Stu: stu}, nil
}

// 删除学生信息
func (s *server) DeleteStudent(ctx context.Context, args *api.DeleteStudentArgs) (*api.DeleteStudentReply, error) {

	reply, err := s.deleteStudent(ctx, args.Uid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.DeleteStudentReply{Reply: reply}, nil
}
