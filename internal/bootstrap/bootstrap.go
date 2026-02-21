package bootstrap

import (
	"centr_rosta/internal/config"
	rhandler "centr_rosta/internal/handler"
	authhandler "centr_rosta/internal/handler/auth"
	sessionrepository "centr_rosta/internal/infra/session"
	"centr_rosta/internal/repository"
	authrepository "centr_rosta/internal/repository/user"
	authservice "centr_rosta/internal/usecase/auth"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type cacheData struct {
	session sessionrepository.RepositorySession
}

type repoData struct {
	repoAuth authrepository.RepositoryUser
}

type serviceData struct {
	useCaseAuth authservice.UseCaseAuth
}

type handlerData struct {
	handler     rhandler.Handler
	handlerAuth authhandler.HandlerAuth
}

var cache = &cacheData{}
var repo = &repoData{}
var serv = &serviceData{}
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
	repo.repoAuth = authrepository.NewRepositoryUser(db)
}

func serviceInit(repo *repoData, cache *cacheData, serv *serviceData) {
	serv.useCaseAuth = authservice.NewService(repo.repoAuth, cache.session)
}

func handlerInit(serv *serviceData, handler *handlerData) {
	handler.handlerAuth = authhandler.NewHandlerAuth(serv.useCaseAuth)

	handler.handler = rhandler.NewHandler(handler.handlerAuth)
}
