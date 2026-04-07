package bootstrap

import (
	"centr_rosta/internal/config"
	"centr_rosta/internal/domain/usecase/admin"
	"centr_rosta/internal/domain/usecase/admin/admin_user"
	"centr_rosta/internal/domain/usecase/auth"
	"centr_rosta/internal/domain/usecase/lesson"
	validateus "centr_rosta/internal/domain/usecase/validate"
	hand "centr_rosta/internal/handler"
	handadmin "centr_rosta/internal/handler/admin"
	handadminus "centr_rosta/internal/handler/admin/admin_user"
	handauth "centr_rosta/internal/handler/auth"
	handlesson "centr_rosta/internal/handler/lesson"
	"centr_rosta/internal/handler/middleware"
	jwts "centr_rosta/internal/infra/jwt"
	"centr_rosta/internal/infra/pass_hash"
	"fmt"

	pg "centr_rosta/internal/infra/postgres"
	re "centr_rosta/internal/infra/redis"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type cacheData struct {
	session *re.SessionRepository
}

type repoData struct {
	repoUser        *pg.UserRepository
	repoTransaction *pg.TransactionRepository
	repoLesson      *pg.LessonRepository
}

type useCaseData struct {
	useCaseAuth      auth.UseCaseAuth
	useCaseAdmin     admin.UseCaseAdmin
	useCaseLesson    lesson.UseCaseLesson
	useCaseAdminUser admin_user.UseCaseAdminUser
}

type handlerData struct {
	handler          *hand.Handler
	handlerAuth      *handauth.HandlerAuth
	handlerAdmin     *handadmin.HandlerAdmin
	handlerLesson    *handlesson.HandlerLesson
	handlerAdminUser *handadminus.AdminUserHandler
}

type middlewareData struct {
	middleware *middleware.Middleware
}

type jwtData struct {
	jwt *jwts.ServiceJwt
}

type passHashData struct {
	passHash *pass_hash.PassHash
}

var validate validateus.Validate
var cache = &cacheData{}
var repo = &repoData{}
var serv = &useCaseData{}
var handler = &handlerData{}
var mw = &middlewareData{}
var jwt = &jwtData{}
var passHash = &passHashData{}

func Bootstrap() (rdb *redis.Client, h hand.Handler) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Env.RedisHost, config.Env.RedisPort),
		Password: config.Env.RedisPass,
		DB:       0,
	})
	db := pg.Connect()

	jwtInit([]byte(config.Env.JwtSecret))
	passHashInit()
	middlewareInit()
	cacheInit(rdb, cache)
	repositoryInit(db, repo)
	useCaseInit(repo, cache, serv, jwt, passHash)
	handlerInit(serv, handler)

	h = *handler.handler

	return
}

func jwtInit(secret []byte) {
	jwt.jwt = jwts.NewServiceJwt(secret)
}

func passHashInit() {
	passHash.passHash = pass_hash.NewPassHash()
}

func middlewareInit() {
	mw.middleware = middleware.NewMiddleware()
}

func cacheInit(rdb *redis.Client, cache *cacheData) {
	cache.session = re.NewRepositorySession(rdb)
}

func repositoryInit(db *gorm.DB, repo *repoData) {
	repo.repoUser = pg.NewUserRepository(db)
	repo.repoTransaction = pg.NewTransactionRepository(db)
	repo.repoLesson = pg.NewLessonRepository(db)
}

func useCaseInit(repo *repoData, cache *cacheData, useCase *useCaseData, jwt *jwtData, passHash *passHashData) {
	validate = validateus.NewValidate(cache.session, jwt.jwt)

	useCase.useCaseAuth = auth.NewUseCaseAuth(repo.repoUser, cache.session, jwt.jwt, passHash.passHash, validate)
	useCase.useCaseAdmin = admin.NewUseCaseAdmin(repo.repoTransaction, validate)
	useCase.useCaseLesson = lesson.NewUseCaseLesson(repo.repoLesson, cache.session, validate)
	useCase.useCaseAdminUser = admin_user.NewUseCaseAdminUser(repo.repoUser, validate, passHash.passHash)
}

func handlerInit(useCase *useCaseData, handler *handlerData) {
	handler.handlerAuth = handauth.NewHandlerAuth(useCase.useCaseAuth)
	handler.handlerAdmin = handadmin.NewHandlerAdmin(useCase.useCaseAdmin)
	handler.handlerLesson = handlesson.NewHandlerLesson(useCase.useCaseLesson)
	handler.handlerAdminUser = handadminus.NewAdminUserHandler(useCase.useCaseAdminUser)

	handler.handler = hand.NewHandler(*handler.handlerAuth, *handler.handlerAdmin, *handler.handlerLesson, *handler.handlerAdminUser, *mw.middleware)
}
