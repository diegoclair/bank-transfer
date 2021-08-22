package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/diegoclair/bank-transfer/domain/entity"
	"github.com/diegoclair/bank-transfer/infra/auth"
	"github.com/diegoclair/go_utils-lib/v2/resterrors"
	"github.com/labstack/gommon/log"
)

const (
	tokenExpirationTime        = 30 * time.Minute
	wrongLogin          string = "Document or password are wrong"
)

type authService struct {
	svc *Service
}

func newAuthService(svc *Service) AuthService {
	return &authService{
		svc: svc,
	}
}

func (s *authService) Login(appContext context.Context, documentNumber, password string) (retVal entity.Authentication, err error) {

	log.Info("Login: Process Started")
	defer log.Info("Login: Process Finished")

	encryptedDocumentNumber, err := s.svc.cipher.Encrypt(documentNumber)
	if err != nil {
		log.Error("Login: ", err)
		return retVal, err
	}

	log.Info("testeeeee", s.svc.dm)
	user, err := s.svc.dm.MySQL().User().GetUserByDocument(encryptedDocumentNumber)
	if err != nil {
		log.Error("Login: Invalid document number received: ", err)
		return retVal, resterrors.NewUnauthorizedError(wrongLogin)
	}
	log.Info("Login: user_id: ", user.ID)
	log.Info("Login: user_uuid: ", user.UUID)
	log.Info("Login: user_name: ", user.Name)

	hasher := md5.New()
	hasher.Write([]byte(password))
	pass := hex.EncodeToString(hasher.Sum(nil))

	if pass != user.Password {
		return retVal, resterrors.NewUnauthorizedError(wrongLogin)
	}

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(tokenExpirationTime)

	claims := &entity.TokenData{}
	claims.UserID = user.UUID
	claims.LoggedIn = true
	claims.IssuedAt = issuedAt.Unix()
	claims.ExpiresAt = expiresAt.Unix()

	newToken, err := auth.GenerateToken(s.svc.cfg.App.Auth, claims)

	if err != nil {
		log.Error("Login: ", err)
		return retVal, err
	}

	retVal.Token = newToken
	retVal.ServerTime = issuedAt.Unix()
	retVal.ValidTime = expiresAt.Unix()

	return retVal, nil
}

func (s *authService) GetDataToken(context context.Context) (retVal *entity.TokenData, err error) {

	token := context.Value(auth.ContextTokenKey)

	jwtToken, ok := token.(*jwt.Token)
	if !ok {
		return retVal, resterrors.NewUnauthorizedError("Invalid data claims")
	}

	claims, ok := jwtToken.Claims.(*entity.TokenData)
	if !ok {
		return retVal, resterrors.NewUnauthorizedError("Invalid data claims")
	}

	return claims, nil
}
