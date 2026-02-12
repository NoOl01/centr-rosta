package bootstrap

import (
	"centr_rosta/internal/config"
	rhandler "centr_rosta/internal/handler"
	authhandler "centr_rosta/internal/handler/auth"
	sessionrepository "centr_rosta/internal/infra/session"
	"centr_rosta/internal/repository"
	authrepository "centr_rosta/internal/repository/auth"
	authservice "centr_rosta/internal/service/auth"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type cacheData struct {
	session sessionrepository.RepositorySession
}

type repoData struct {
	repoAuth authrepository.RepositoryAuth
}

type serviceData struct {
	serviceAuth authservice.ServiceAuth
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
	repo.repoAuth = authrepository.NewRepositoryAuth(db)
}

func serviceInit(repo *repoData, cache *cacheData, serv *serviceData) {
	serv.serviceAuth = authservice.NewService(repo.repoAuth, cache.session)
}

func handlerInit(serv *serviceData, handler *handlerData) {
	handler.handlerAuth = authhandler.NewHandlerAuth(serv.serviceAuth)

	handler.handler = rhandler.NewHandler(handler.handlerAuth)
}
