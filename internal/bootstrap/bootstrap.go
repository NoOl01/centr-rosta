package bootstrap

import (
	"centr_rosta/internal/config"
	rhandler "centr_rosta/internal/handler"
	adminhandler "centr_rosta/internal/handler/admin"
	authhandler "centr_rosta/internal/handler/auth"
	sessionrepository "centr_rosta/internal/infra/session"
	"centr_rosta/internal/repository"
	"centr_rosta/internal/repository/lesson"
	"centr_rosta/internal/repository/transaction"
	"centr_rosta/internal/repository/user"
	adminservice "centr_rosta/internal/usecase/admin"
	authservice "centr_rosta/internal/usecase/auth"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type cacheData struct {
	session sessionrepository.RepositorySession
}

type repoData struct {
	repoUser        user.RepositoryUser
	repoLesson      lesson.RepositoryLesson
	repoTransaction transaction.RepositoryTransaction
}

type useCaseData struct {
	useCaseAuth  authservice.UseCaseAuth
	useCaseAdmin adminservice.UseCaseAdmin
}

type handlerData struct {
	handler      rhandler.Handler
	handlerAuth  authhandler.HandlerAuth
	handlerAdmin adminhandler.HandlerAdmin
}

var cache = &cacheData{}
var repo = &repoData{}
var serv = &useCaseData{}
var handler = &handlerData{}

func Bootstrap() (rdb *redis.Client, h rhandler.Handler) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Env.RedisHost, config.Env.RedisPort),
		Password: config.Env.RedisPass,
		DB:       0,
	})
	db := repository.Connect()

	cacheInit(rdb, cache)
	repositoryInit(db, repo)
	serviceInit(repo, cache, serv)
	handlerInit(serv, handler)

	h = handler.handler

	return
}

func cacheInit(rdb *redis.Client, cache *cacheData) {
	cache.session = sessionrepository.NewRepositorySession(rdb)
}

func repositoryInit(db *gorm.DB, repo *repoData) {
	repo.repoUser = user.NewRepositoryUser(db)
	repo.repoLesson = lesson.NewRepositoryLesson(db)
	repo.repoTransaction = transaction.NewRepositoryTransaction(db)
}

func serviceInit(repo *repoData, cache *cacheData, useCase *useCaseData) {
	useCase.useCaseAuth = authservice.NewService(repo.repoUser, cache.session)
	useCase.useCaseAdmin = adminservice.NewUseCaseAdmin(repo.repoTransaction, cache.session)
}

func handlerInit(useCase *useCaseData, handler *handlerData) {
	handler.handlerAuth = authhandler.NewHandlerAuth(useCase.useCaseAuth)
	handler.handlerAdmin = adminhandler.NewHandlerAdmin(useCase.useCaseAdmin)

	handler.handler = rhandler.NewHandler(handler.handlerAuth, handler.handlerAdmin)
}
