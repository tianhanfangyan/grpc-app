package impl

import (
	"context"
	"errors"
	"github.com/tianhanfangyan/grpc-app/pkg/api"
	"strconv"
	"sync"
	"time"
)

var (
	StudentMap sync.Map
)

func init() {
	stu := api.Student{}
	stu.Id = 1
	stu.Name = "ben"
	stu.Age = 20
	stu.Sex = "man"
	StudentMap.Store("user_1", &stu)
}

func (s *server) addStudent(ctx context.Context, ss *api.Student) (string, error) {
	uid := "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	StudentMap.Store(uid, ss)
	return uid, nil
}

func (s *server) getStudent(ctx context.Context, uid string) (*api.Student, error) {
	var stu api.Student

	if s, ok := StudentMap.Load(uid); ok {
		stu.Id = s.(*api.Student).Id
		stu.Name = s.(*api.Student).Name
		stu.Age = s.(*api.Student).Age
		stu.Sex = s.(*api.Student).Sex
		return &stu, nil
	}
	return nil, errors.New("Student not exists")
}

func (s *server) getAllStudent(ctx context.Context) ([]*api.Student, error) {
	var DataList []*api.Student

	//Range
	//遍历sync.Map, 要求输入一个func作为参数
	f := func(k, v interface{}) bool {
		//这个函数的入参、出参的类型都已经固定，不能修改
		//可以在函数体内编写自己的代码，调用map中的k,v
		DataList = append(DataList, v.(*api.Student))
		return true
	}
	StudentMap.Range(f)

	return DataList, nil
}

func (s *server) updateStudent(ctx context.Context, uid string, ss *api.Student) (*api.Student, error) {
	if s, ok := StudentMap.Load(uid); ok {
		if ss.Name != "" {
			s.(*api.Student).Name = ss.Name
		}
		if ss.Age != 0 {
			s.(*api.Student).Age = ss.Age
		}
		if ss.Sex != "" {
			s.(*api.Student).Sex = ss.Sex
		}

		return s.(*api.Student), nil
	}
	return nil, errors.New("User Not Exist")
}

func (s *server) deleteStudent(ctx context.Context, uid string) (*api.Reply, error) {
	var reply api.Reply
	StudentMap.Delete(uid)

	reply.Msg = "delete successsful"
	reply.Status = 200

	return &reply, nil
}
