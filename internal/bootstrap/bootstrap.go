package bootstrap

import (
	"centr_rosta/internal/config"
	"centr_rosta/internal/domain/usecase/admin"
	"centr_rosta/internal/domain/usecase/admin/admin_user"
	adminpl "centr_rosta/internal/domain/usecase/admin/personal_lesson"
	"centr_rosta/internal/domain/usecase/auth"
	"centr_rosta/internal/domain/usecase/lesson"
	hand "centr_rosta/internal/handler"
	handadmin "centr_rosta/internal/handler/admin"
	handadminus "centr_rosta/internal/handler/admin/admin_user"
	handadminpl "centr_rosta/internal/handler/admin/personal_lesson"
	handauth "centr_rosta/internal/handler/auth"
	handlesson "centr_rosta/internal/handler/lesson"
	handmiddleware "centr_rosta/internal/handler/middleware"
	jwts "centr_rosta/internal/infra/jwt"
	"centr_rosta/internal/infra/middleware"
	"centr_rosta/internal/infra/pass_hash"
	"fmt"

	pg "centr_rosta/internal/infra/postgres"
	re "centr_rosta/internal/infra/redis"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type infraData struct {
	session  *re.SessionRepository
	jwt      *jwts.ServiceJwt
	passHash *pass_hash.PassHash
	validate *middleware.ValidateMiddleWare
}

type repoData struct {
	repoUser            *pg.UserRepository
	repoTransaction     *pg.TransactionRepository
	repoLesson          *pg.LessonRepository
	repoPersonalLessons *pg.PersonalLessonRepository
}

type useCaseData struct {
	useCaseAuth                 auth.UseCaseAuth
	useCaseAdmin                admin.UseCaseAdmin
	useCaseLesson               lesson.UseCaseLesson
	useCaseAdminUser            admin_user.UseCaseAdminUser
	useCaseAdminPersonalLessons adminpl.UseCaseAdminPersonalLesson
}

type handlerData struct {
	handler                     *hand.Handler
	handlerMiddleware           *handmiddleware.Middleware
	handlerAuth                 *handauth.HandlerAuth
	handlerAdmin                *handadmin.HandlerAdmin
	handlerLesson               *handlesson.HandlerLesson
	handlerAdminUser            *handadminus.AdminUserHandler
	handlerAdminPersonalLessons *handadminpl.AdminPersonalLessonHandler
}

var infra = &infraData{}
var repo = &repoData{}
var useCase = &useCaseData{}
var handler = &handlerData{}

func Bootstrap() (rdb *redis.Client, h hand.Handler) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Env.RedisHost, config.Env.RedisPort),
		Password: config.Env.RedisPass,
		DB:       0,
	})
	db := pg.Connect()

	infraInit(rdb)
	repositoryInit(db)
	useCaseInit()
	handlerInit()

	h = *handler.handler

	return
}

func infraInit(rdb *redis.Client) {
	infra.session = re.NewRepositorySession(rdb)
	infra.jwt = jwts.NewServiceJwt([]byte(config.Env.JwtSecret))
	infra.passHash = pass_hash.NewPassHash()
	infra.validate = middleware.NewValidateMiddleWare(*infra.session, *infra.jwt)
}

func repositoryInit(db *gorm.DB) {
	repo.repoUser = pg.NewUserRepository(db)
	repo.repoTransaction = pg.NewTransactionRepository(db)
	repo.repoLesson = pg.NewLessonRepository(db)
	repo.repoPersonalLessons = pg.NewPersonalLessonRepository(db)
}

func useCaseInit() {
	useCase.useCaseAuth = auth.NewUseCaseAuth(repo.repoUser, infra.session, infra.jwt, infra.passHash, infra.validate)
	useCase.useCaseLesson = lesson.NewUseCaseLesson(repo.repoLesson, infra.session, infra.validate)
	useCase.useCaseAdmin = admin.NewUseCaseAdmin(repo.repoTransaction, infra.validate)
	useCase.useCaseAdminUser = admin_user.NewUseCaseAdminUser(repo.repoUser, infra.validate, infra.passHash)
	useCase.useCaseAdminPersonalLessons = adminpl.NewUseCaseAdminPersonalLesson(repo.repoPersonalLessons, infra.validate)
}

func handlerInit() {
	handler.handlerAuth = handauth.NewHandlerAuth(useCase.useCaseAuth)
	handler.handlerAdmin = handadmin.NewHandlerAdmin(useCase.useCaseAdmin)
	handler.handlerLesson = handlesson.NewHandlerLesson(useCase.useCaseLesson)
	handler.handlerAdminUser = handadminus.NewAdminUserHandler(useCase.useCaseAdminUser)
	handler.handlerAdminPersonalLessons = handadminpl.NewAdminPersonalLessonHandler(useCase.useCaseAdminPersonalLessons)

	handler.handlerMiddleware = handmiddleware.NewMiddleware()
	handler.handler = hand.NewHandler(*handler.handlerAuth, *handler.handlerAdmin, *handler.handlerLesson, *handler.handlerAdminUser, *handler.handlerAdminPersonalLessons, *handler.handlerMiddleware)
}
