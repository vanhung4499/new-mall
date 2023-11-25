package services

import (
	"context"
	"errors"
	"mime/multipart"
	"new-mall/internal/global"
	"new-mall/internal/models"
	"new-mall/internal/repositories"
	"new-mall/internal/types"
	"new-mall/pkg/common"
	"new-mall/pkg/component/hasher"
	"new-mall/pkg/component/salt"
	"new-mall/pkg/component/tokenprovider"
	"new-mall/pkg/component/tokenprovider/jwt"
	"new-mall/pkg/component/uploadprovider"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

// RegisterUser user registration
func (s *UserService) RegisterUser(ctx context.Context, data *models.UserCreate) error {
	user, err := s.UserRepo.FindWithCondition(ctx, map[string]interface{}{"email": data.Email})

	if errors.Is(err, common.RecordNotFound) {

		data.Salt = salt.GenSalt(50)
		data.Password = hasher.Hash(data.Password + data.Salt)
		data.Role = common.RoleUser

		data.Avatar = common.UserDefaultAvatarLocal
		if global.CONFIG.System.UploadModel == common.UploadModelS3 {
			// If the configuration uses OSS, use URL as the default avatar
			data.Avatar = common.UserDefaultAvatarOss
		}

		// Create user
		if err = s.UserRepo.Create(ctx, data); err != nil {
			return common.ErrCannotCreateEntity(models.UserEntityName, err)
		}
	} else if user != nil {
		return models.ErrEmailExisted
	} else {
		return err
	}

	return nil
}

// LoginUser user login function
func (s *UserService) LoginUser(ctx context.Context, data *types.UserLoginReq) (*types.UserTokenRes, error) {
	user, err := s.UserRepo.FindWithCondition(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, models.ErrUsernameOrPasswordInvalid
	}

	hashedPassword := hasher.Hash(data.Password + user.Salt)

	if user.Password != hashedPassword {
		return nil, models.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.ID,
		Role:   user.Role.String(),
	}

	tokenprovider := jwt.NewTokenJWTProvider(global.CONFIG.EncryptSecret.JwtSecret)

	accessToken, err := tokenprovider.Generate(payload, global.CONFIG.Token.AccessTokenExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := tokenprovider.Generate(payload, global.CONFIG.Token.RefreshTokenExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if err != nil {
		return nil, err
	}

	res := types.UserTokenRes{
		User:         user,
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
	}

	return &res, nil
}

// UpdateUser update user information
func (s *UserService) UpdateUser(ctx context.Context, id uint, data *models.UserUpdate) error {
	if err := s.UserRepo.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}

// UploadUserAvatar update avatar
func (s *UserService) UploadUserAvatar(ctx context.Context, userID uint, file *multipart.FileHeader) error {

	user, err := s.UserRepo.FindWithCondition(ctx, map[string]interface{}{"id": userID})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrEntityNotFound(models.UserEntityName, err)
		}
		return err
	}

	oss := uploadprovider.NewUploadProvider()
	filePath, _, err := oss.UploadFile(ctx, file)

	if err != nil {
		return models.ErrCannotSaveFile(err)
	}

	user.Avatar = filePath
	err = s.UserRepo.Save(ctx, user)
	if err != nil {
		return common.ErrCannotUpdateEntity(models.UserEntityName, err)
	}

	return nil
}

// GetUser show user information
func (s *UserService) GetUser(ctx context.Context, id uint) (*models.User, error) {

	user, err := s.UserRepo.FindWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return nil, common.ErrEntityNotFound(models.UserEntityName, err)
		}
		return nil, err
	}

	return user, nil
}
