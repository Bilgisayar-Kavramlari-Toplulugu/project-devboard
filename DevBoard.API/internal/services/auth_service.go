package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"project-devboard/internal/config"
	domain "project-devboard/internal/domain/entities"
	"project-devboard/internal/dtos"
	"project-devboard/internal/repository"
	"project-devboard/pkg/apperrors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(req dtos.LoginRequest) (*domain.User, *TokenPair, error)
	Signup(req dtos.SignupRequest, actorID uuid.UUID) (*TokenPair, error)
	Refresh(refreshToken string) (*TokenPair, error)
	Logout(refreshToken string) error
	GetMe(userID uuid.UUID) (*domain.User, error)
	ForgotPassword(email string) error
	ResetPassword(token, newPassword string) error
}

type authService struct {
	db         *gorm.DB
	userRepo   repository.UserRepository
	jwtService JWTService
	cfg        *config.Config
}

func NewAuthService(db *gorm.DB, userRepo repository.UserRepository, jwtService JWTService, cfg *config.Config) AuthService {
	return &authService{db: db, userRepo: userRepo, jwtService: jwtService, cfg: cfg}
}

func (s *authService) Login(req dtos.LoginRequest) (*domain.User, *TokenPair, error) {
	user, err := s.userRepo.GetByEmailWithRoles(req.Email)
	if err != nil {
		return nil, nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}
	if user == nil {
		return nil, nil, apperrors.New(apperrors.Unauthorized, apperrors.ErrInvalidCredentials)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, nil, apperrors.New(apperrors.Unauthorized, apperrors.ErrInvalidCredentials)
	}

	role := primaryRoleNameFromRoles(user.UserRoles)

	tokenPair, err := s.jwtService.GenerateTokenPair(user.Id, role)
	if err != nil {
		return nil, nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	if err := s.jwtService.CreateSession(user.Id, tokenPair.RefreshToken); err != nil {
		return nil, nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	user.Password = ""
	return user, tokenPair, nil
}

func (s *authService) Signup(req dtos.SignupRequest, actorID uuid.UUID) (*TokenPair, error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, tx.Error)
	}

	txUserRepo := s.userRepo.WithTx(tx)

	existing, err := txUserRepo.GetByEmail(req.Email)
	if err != nil {
		tx.Rollback()
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}
	if existing != nil {
		tx.Rollback()
		return nil, apperrors.New(apperrors.Conflict, apperrors.ErrUserAlreadyExists)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	newUserID := uuid.New()
	if actorID == uuid.Nil {
		actorID = newUserID
	}

	newUser := &domain.User{
		Id:               newUserID,
		Email:            req.Email,
		Password:         string(hashed),
		Firstname:        req.Firstname,
		Lastname:         req.Lastname,
		PhoneNumber:      req.PhoneNumber,
		IsEmailValidated: false,
		BaseEntity: domain.BaseEntity{
			IsActive:       true,
			CreatedBy:      actorID,
			LastModifiedBy: actorID,
		},
	}

	if err := txUserRepo.Create(newUser); err != nil {
		tx.Rollback()
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	role, err := ensureDefaultRole(tx, actorID)
	if err != nil {
		tx.Rollback()
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	userRole := domain.UserRole{
		UserId: newUser.Id,
		RoleId: role.Id,
		BaseEntity: domain.BaseEntity{
			IsActive:       true,
			CreatedBy:      actorID,
			LastModifiedBy: actorID,
		},
	}
	if err := tx.Create(&userRole).Error; err != nil {
		tx.Rollback()
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	tokenPair, err := s.jwtService.GenerateTokenPair(newUser.Id, role.Name)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	if err := s.jwtService.CreateSession(newUser.Id, tokenPair.RefreshToken); err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	return tokenPair, nil
}

func (s *authService) Refresh(refreshToken string) (*TokenPair, error) {
	claims, err := s.jwtService.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, apperrors.New(apperrors.Unauthorized, apperrors.ErrInvalidToken)
	}

	tokenPair, err := s.jwtService.RefreshTokens(refreshToken, claims.UserID)
	if err != nil {
		return nil, apperrors.New(apperrors.Unauthorized, apperrors.ErrInvalidToken)
	}

	oldHash := sha256.Sum256([]byte(refreshToken))
	_ = s.jwtService.RevokeSession(hex.EncodeToString(oldHash[:]))

	if err := s.jwtService.CreateSession(claims.UserID, tokenPair.RefreshToken); err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	return tokenPair, nil
}

func (s *authService) Logout(refreshToken string) error {
	hash := sha256.Sum256([]byte(refreshToken))
	if err := s.jwtService.RevokeSession(hex.EncodeToString(hash[:])); err != nil {
		return apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}
	return nil
}

func (s *authService) GetMe(userID uuid.UUID) (*domain.User, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}
	if user == nil {
		return nil, apperrors.New(apperrors.NotFound, apperrors.ErrUserNotFound)
	}
	user.Password = ""
	return user, nil
}

func (s *authService) ForgotPassword(email string) error {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}
	if user == nil {
		return nil
	}

	if err := s.db.Model(&domain.PasswordResetToken{}).
		Where("user_id = ? AND used_at IS NULL", user.Id).
		Update("used_at", time.Now()).Error; err != nil {
		return apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}
	rawToken := hex.EncodeToString(tokenBytes)

	hash := sha256.Sum256([]byte(rawToken))
	tokenHash := hex.EncodeToString(hash[:])

	resetToken := &domain.PasswordResetToken{
		Id:        uuid.New(),
		UserId:    user.Id,
		Token:     tokenHash,
		ExpiresAt: time.Now().Add(time.Duration(s.cfg.PasswordResetExpireMinutes) * time.Minute),
		CreatedOn: time.Now(),
	}

	if err := s.db.Create(resetToken).Error; err != nil {
		return apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	return nil
}

func (s *authService) ResetPassword(token, newPassword string) error {
	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])

	var resetToken domain.PasswordResetToken
	if err := s.db.Where("token = ?", tokenHash).First(&resetToken).Error; err != nil {
		return apperrors.New(apperrors.Unauthorized, apperrors.ErrInvalidToken)
	}

	if !resetToken.IsValid() {
		return apperrors.New(apperrors.Unauthorized, apperrors.ErrInvalidToken)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, tx.Error)
	}

	if err := tx.Model(&domain.User{}).Where("id = ?", resetToken.UserId).
		Update("password", string(hashed)).Error; err != nil {
		tx.Rollback()
		return apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	now := time.Now()
	if err := tx.Model(&resetToken).Update("used_at", &now).Error; err != nil {
		tx.Rollback()
		return apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	if err := tx.Commit().Error; err != nil {
		return apperrors.Wrap(apperrors.InternalError, apperrors.ErrInternalServer, err)
	}

	return nil
}

func primaryRoleNameFromRoles(userRoles []domain.UserRole) string {
	if len(userRoles) > 0 && userRoles[0].Role.Name != "" {
		return userRoles[0].Role.Name
	}
	return "Developer"
}

func ensureDefaultRole(tx *gorm.DB, actorID uuid.UUID) (*domain.Role, error) {
	var role domain.Role
	err := tx.Where("name = ?", "Developer").First(&role).Error
	if err == nil {
		return &role, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	role = domain.Role{
		Name: "Developer",
		BaseEntity: domain.BaseEntity{
			IsActive:       true,
			CreatedBy:      actorID,
			LastModifiedBy: actorID,
		},
	}
	if err := tx.Create(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
