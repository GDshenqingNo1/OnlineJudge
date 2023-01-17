package user

import (
	g "OnlineJudge/app/global"
	"OnlineJudge/app/internal/model/user"
	"OnlineJudge/utils/jwt"
	"context"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"math/rand"
	"strconv"

	"time"
)

type SUser struct{}

var insUser = SUser{}

func (s *SUser) CheckUserIsExist(ctx context.Context, username string) error {
	user := &user.User{}
	err := g.MysqlDB.WithContext(ctx).Table("user").Select("username").Where("username = ?", username).First(user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			g.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user"),
			)
			return fmt.Errorf("internal err")
		}
	} else {
		return fmt.Errorf("username already exist")
	}
	return nil
}
func (s *SUser) EncryptPassword(password string) string {
	d := sha3.Sum224([]byte(password))
	return hex.EncodeToString(d[:])
}

func (s *SUser) CreateUser(ctx context.Context, userSubject *user.User) error {
	err := g.MysqlDB.WithContext(ctx).Table("user").Create(userSubject).Error
	return err
}

func (s *SUser) CheckPassword(ctx context.Context, user *user.User) error {
	err := g.MysqlDB.WithContext(ctx).
		Table("user").
		Where(&user).
		First(user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			g.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user"),
			)
			return fmt.Errorf("internal err")
		} else {
			return fmt.Errorf("invalid username or password")
		}
	}

	return nil
}

func (s *SUser) GenerateToken(ctx context.Context, user *user.User) (string, error) {
	config := g.Config.Middleware.Jwt
	j := jwt.NewJWT(&jwt.Config{
		SecretKey:   config.SecretKey,
		ExpiresTime: config.ExpiresTime,
		BufferTime:  config.BufferTime,
		Issuer:      config.Issuer})
	claims := j.CreateClaims(&jwt.BaseClaims{
		Id:         user.Id,
		Username:   user.Username,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	})
	tokenString, err := j.GenerateToken(&claims)
	if err != nil {
		g.Logger.Error("generate token failed.", zap.Error(err))
		return "", fmt.Errorf("internal err")
	}
	err = g.Rdb.Set(ctx, fmt.Sprintf("jwt:%d", user.Id), tokenString, time.Duration(config.ExpiresTime)*time.Second).Err()
	if err != nil {
		g.Logger.Error("set redis cache failed.",
			zap.Error(err), zap.String("key", "jwt:[id]"),
			zap.Int64("id", user.Id),
		)
		return "", fmt.Errorf("internal err")
	}
	return tokenString, nil

}

func (s *SUser) SendCode(ctx context.Context, mail string) error {
	rand.Seed(time.Now().UnixNano())
	code := strconv.Itoa(rand.Intn(100000))
	g.Rdb.Set(ctx, mail, code, time.Second*60)
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress("1960441553@qq.com", "MyOnlineJudge"))
	m.SetHeader("To", mail)
	m.SetHeader("Subject", "注册验证码已发送")
	m.SetBody("text/html", "您的验证码：<b>"+code+"</b>")
	d := gomail.NewDialer("smtp.qq.com", 587, "1960441553", "uddnhcmjxzmsgchf")
	err := d.DialAndSend(m)
	return err
}

func (s *SUser) CheckMailIsExist(ctx context.Context, mail string) error {
	var userSubject = &user.User{}
	err := g.MysqlDB.WithContext(ctx).
		Table("user").
		Where("mail=?", mail).
		First(userSubject).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			g.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user"),
			)
			return fmt.Errorf("internal err")
		}
	} else {
		return fmt.Errorf("mail is already signed")
	}
	return nil
}
func (s *SUser) CheckCode(ctx context.Context, mail string, code string) error {
	cmd := g.Rdb.Get(ctx, mail)
	err := cmd.Err()
	if err != nil {
		g.Logger.Error("get code from rdb error", zap.Error(err))
		return fmt.Errorf("internal err")
	}
	if code != cmd.Val() {
		return fmt.Errorf("incorrect code")
	}
	return nil

}
