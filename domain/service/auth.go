package service

import (
	"crypto/md5"
	"encoding/hex"

	"time"

	"github.com/diegoclair/bank-transfer/domain/entity"
	"github.com/diegoclair/bank-transfer/infra/auth"
	"github.com/diegoclair/bank-transfer/util/errors"
	"github.com/diegoclair/go_utils-lib/v2/resterrors"
	"github.com/labstack/gommon/log"
	"github.com/twinj/uuid"
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

func (s *authService) CreateUser(user entity.User) (err error) {

	log.Info("CreateUser: Process Started")
	defer log.Info("CreateUser: Process Finished")

	user.DocumentNumber, err = s.svc.cipher.Encrypt(user.DocumentNumber)
	if err != nil {
		log.Error("CreateUser: ", err)
		return err
	}

	_, err = s.svc.dm.MySQL().User().GetUserByDocument(user.DocumentNumber)
	if err != nil && !errors.SQLNotFound(err.Error()) {
		log.Error("CreateUser: ", err)
		return err
	} else if err == nil {
		log.Error("CreateUser: The document number is already in use")
		return resterrors.NewConflictError("The document number is already in use")
	}

	hasher := md5.New()
	hasher.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hasher.Sum(nil))

	user.UUID = uuid.NewV4().String()

	err = s.svc.dm.MySQL().User().CreateUser(user)
	if err != nil {
		log.Error("CreateUser: ", err)
		return err
	}

	return nil
}

func (s *authService) Login(documentNumber, password string) (retVal entity.Authentication, err error) {

	log.Info("Login: Process Started")
	defer log.Info("Login: Process Finished")

	encryptedDocumentNumber, err := s.svc.cipher.Encrypt(documentNumber)
	if err != nil {
		log.Error("Login: ", err)
		return retVal, err
	}

	user, err := s.svc.dm.MySQL().User().GetUserByDocument(encryptedDocumentNumber)
	if err != nil {
		log.Error("Login: ", err)
		return retVal, resterrors.NewUnauthorizedError(wrongLogin)
	}

	log.Info("Login: user_id: ", user.ID)
	log.Info("Login: user_uuid: ", user.UUID)
	log.Info("Login: user_name: ", user.Name)

	hasher := md5.New()
	hasher.Write([]byte(password))
	pass := hex.EncodeToString(hasher.Sum(nil))

	if pass != user.Password {
		log.Error("Login: wrong password")
		return retVal, resterrors.NewUnauthorizedError(wrongLogin)
	}

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(tokenExpirationTime)

	claims := &entity.TokenData{}
	claims.UserID = user.UUID
	claims.LoggedIn = true
	claims.IssuedAt = issuedAt.Unix()
	claims.ExpiresAt = expiresAt.Unix()
	claims.Issuer = "ST-BANK-TRANSFER"

	newToken, err := auth.GenerateToken(s.svc.cfg.App.Auth, claims)
	if err != nil {
		log.Error("Login: error to generate token: ", err)
		return retVal, err
	}

	retVal.Token = newToken
	retVal.ServerTime = issuedAt.Unix()
	retVal.ValidTime = expiresAt.Unix()

	return retVal, nil
}
