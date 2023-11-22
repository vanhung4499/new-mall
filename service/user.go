package service

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"new-mall/config"
	"new-mall/constant"
	"new-mall/pkg/utils"
	"new-mall/pkg/utils/ctl"
	"new-mall/pkg/utils/email"
	"new-mall/pkg/utils/upload"
	"new-mall/repository/db/dao"
	"new-mall/repository/db/model"
	"new-mall/types"
	"sync"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

// UserRegister user registration
func (s *UserSrv) UserRegister(ctx context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
	if exist {
		err = errors.New("user already exists")
		return
	}
	user := &model.User{
		NickName: req.NickName,
		UserName: req.UserName,
		Status:   model.Active,
		Money:    constant.UserInitMoney,
	}
	// Encrypt password
	if err = user.SetPassword(req.Password); err != nil {
		utils.Logger.Error(err)
		return
	}
	// Default avatar is local
	user.Avatar = constant.UserDefaultAvatarLocal
	if config.Config.System.UploadModel == constant.UploadModelOss {
		// If the configuration uses OSS, use URL as the default avatar
		user.Avatar = constant.UserDefaultAvatarOss
	}

	// Create user
	err = userDao.CreateUser(user)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	return
}

// UserLogin user login function
func (s *UserSrv) UserLogin(ctx context.Context, req *types.UserLoginReq) (res interface{}, err error) {
	var user *model.User
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if !exist {
		utils.Logger.Error(err)
		return nil, errors.New("user does not exist")
	}

	if !user.CheckPassword(req.Password) {
		return nil, errors.New("incorrect username/password")
	}

	accessToken, refreshToken, err := utils.GenerateToken(user.ID, req.UserName)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}

	userRes := &types.UserInfoRes{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.AvatarURL(),
		CreateAt: user.CreatedAt.Unix(),
	}

	res = &types.UserTokenData{
		User:         userRes,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return
}

// UserInfoUpdate update user information
func (s *UserSrv) UserInfoUpdate(ctx context.Context, req *types.UserInfoUpdateReq) (resp interface{}, err error) {
	// Find the user
	u, _ := ctl.GetUserInfo(ctx)
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(u.Id)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}

	if req.NickName != "" {
		user.NickName = req.NickName
	}

	err = userDao.UpdateUserById(u.Id, user)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}

	return
}

// UserAvatarUpload update avatar
func (s *UserSrv) UserAvatarUpload(ctx context.Context, file *multipart.FileHeader, req *types.UserServiceReq) (res interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	uId := u.Id
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}

	oss := upload.NewOss()
	filePath, _, err := oss.UploadFile(file)

	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}

	user.Avatar = filePath
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}

	return
}

// SendEmail send email
func (s *UserSrv) SendEmail(ctx context.Context, req *types.SendEmailServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	var address string
	token, err := utils.GenerateEmailToken(u.Id, req.OperationType, req.Email, req.Password)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	sender := email.NewEmailSender()
	address = config.Config.Email.ValidEmail + token
	mailText := fmt.Sprintf(constant.EmailOperationMap[req.OperationType], address)
	if err = sender.Send(mailText, req.Email, "FanOneMall"); err != nil {
		utils.Logger.Error(err)
		return
	}

	return
}

// Valid validate content
func (s *UserSrv) Valid(ctx context.Context, req *types.ValidEmailServiceReq) (res interface{}, err error) {
	var userId uint
	var email string
	var password string
	var operationType uint
	// Validate token
	if req.Token == "" {
		err = errors.New("token does not exist")
		utils.Logger.Error(err)
		return
	}
	claims, err := utils.ParseEmailToken(req.Token)
	if err != nil {
		utils.Logger.Error(err)
		return
	} else {
		userId = claims.UserID
		email = claims.Email
		password = claims.Password
		operationType = claims.OperationType
	}

	// Get user information
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	switch operationType {
	case constant.EmailOperationBinding:
		user.Email = email
	case constant.EmailOperationNoBinding:
		user.Email = ""
	case constant.EmailOperationUpdatePassword:
		err = user.SetPassword(password)
		if err != nil {
			err = errors.New("password encryption error")
			return
		}
	default:
		return nil, errors.New("operation does not match")
	}

	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	res = &types.UserInfoRes{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.AvatarURL(),
		CreateAt: user.CreatedAt.Unix(),
	}

	return
}

// UserInfoShow show user information
func (s *UserSrv) UserInfoShow(ctx context.Context, req *types.UserInfoShowReq) (res interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
	user, err := dao.NewUserDao(ctx).GetUserById(u.Id)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
	res = &types.UserInfoRes{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.AvatarURL(),
		CreateAt: user.CreatedAt.Unix(),
	}

	return
}

func (s *UserSrv) UserFollow(ctx context.Context, req *types.UserFollowingReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	err = dao.NewUserDao(ctx).FollowUser(u.Id, req.Id)

	return
}

func (s *UserSrv) UserUnFollow(ctx context.Context, req *types.UserUnFollowingReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error(err)
		return nil, err
	}
	err = dao.NewUserDao(ctx).UnFollowUser(u.Id, req.Id)

	return
}
