package problem

import (
	g "OnlineJudge/app/global"
	"OnlineJudge/app/internal/model/problem"
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SProblem struct{}

var insProblem SProblem

func (s *SProblem) GetProblem(ctx context.Context, problemId int) (*problem.Problem, error) {
	problemSubject := &problem.Problem{}
	err := g.MysqlDB.WithContext(ctx).
		Table("problem").
		Select("*").
		Where("id=?", problemId).
		First(problemSubject).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			g.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user"),
			)
			return nil, fmt.Errorf("internal err")
		}
	}
	return problemSubject, nil
}

func (s *SProblem) GetProblemList(ctx context.Context) ([]*problem.Problem, error) {
	list := make([]*problem.Problem, 0)
	err := g.MysqlDB.WithContext(ctx).
		Table("problem").
		Select("*").
		Find(&list).Error

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			g.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user"),
			)
			return nil, err
		}
	}
	return list, nil
}

func (s *SProblem) SaveProblem(ctx context.Context, problemSubject *problem.Problem) error {
	err := g.MysqlDB.WithContext(ctx).Table("problem").Create(problemSubject).Error
	return err
}

func (s *SProblem) UpdateProblemInfo(ctx context.Context, problemSubject *problem.Problem) error {
	err := g.MysqlDB.WithContext(ctx).Table("problem").Where("id=?", problemSubject.Id).Updates(problemSubject).Error
	return err
}

// CheckProblemDoesExist 确认题目确实存在
func (s *SProblem) CheckProblemDoesExist(ctx context.Context, id int) error {
	var problemSubject = &problem.Problem{}
	err := g.MysqlDB.WithContext(ctx).Table("problem").Where("id=?", id).First(problemSubject).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			g.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "problem"),
			)
			return fmt.Errorf("internal err")
		} else {
			return fmt.Errorf("problem is not exist")
		}
	}
	return nil
}
