package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/senbox/services-management/internal/domain/repository"
	"github.com/senbox/services-management/internal/domain/usecase"
	"github.com/senbox/services-management/internal/infrastructure/database"
	infraRepo "github.com/senbox/services-management/internal/infrastructure/repository"
	httpInterface "github.com/senbox/services-management/internal/interface/http"
	"github.com/senbox/services-management/internal/interface/http/handler"
	"github.com/senbox/services-management/internal/interface/middleware"
	"github.com/senbox/services-management/pkg/config"
	"github.com/senbox/services-management/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

// Container holds all application dependencies
type Container struct {
	Config              *config.Config
	Logger              *logger.Logger
	MongoDB             *mongo.Database
	ServiceGroupRepo    repository.ServiceGroupRepository
	ServiceRepo         repository.ServiceRepository
	ServiceGroupUseCase usecase.ServiceGroupUseCase
	ServiceUseCase      usecase.ServiceUseCase
	ServiceGroupHandler *handler.ServiceGroupHandler
	ServiceHandler      *handler.ServiceHandler
	AuditMiddleware     *middleware.AuditMiddleware
	App                 *fiber.App
}

// NewContainer initializes all application dependencies
func NewContainer() (*Container, error) {
	c := &Container{}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	c.Config = cfg

	// Initialize logger
	appLogger, err := logger.NewLogger("logs")
	if err != nil {
		return nil, err
	}
	c.Logger = appLogger
	c.Logger.Info("Logger initialized successfully")

	// Initialize database
	if err := c.initDatabase(); err != nil {
		return nil, err
	}

	// Initialize repositories
	c.initRepositories()

	// Initialize use cases
	c.initUseCases()

	// Initialize handlers
	c.initHandlers()

	// Initialize middlewares
	c.initMiddlewares()

	// Setup router
	c.setupRouter()

	return c, nil
}

// initDatabase initializes MongoDB connection
func (c *Container) initDatabase() error {
	c.Logger.Info("Connecting to MongoDB database")
	mongoDB, err := database.NewMongoConnection(database.MongoConfig{
		Host:     c.Config.MongoDB.Host,
		Port:     c.Config.MongoDB.Port,
		User:     c.Config.MongoDB.User,
		Password: c.Config.MongoDB.Password,
		DBName:   c.Config.MongoDB.DBName,
	})
	if err != nil {
		return err
	}
	c.MongoDB = mongoDB
	c.Logger.Info("MongoDB connection established")
	return nil
}

// initRepositories initializes all repositories
func (c *Container) initRepositories() {
	c.ServiceGroupRepo = infraRepo.NewServiceGroupRepositoryMongo(c.MongoDB)
	c.ServiceRepo = infraRepo.NewServiceRepositoryMongo(c.MongoDB)
}

// initUseCases initializes all use cases
func (c *Container) initUseCases() {
	c.ServiceGroupUseCase = usecase.NewServiceGroupUseCase(c.ServiceGroupRepo)
	c.ServiceUseCase = usecase.NewServiceUseCase(c.ServiceRepo)
}

// initHandlers initializes all HTTP handlers
func (c *Container) initHandlers() {
	c.ServiceGroupHandler = handler.NewServiceGroupHandler(c.ServiceGroupUseCase)
	c.ServiceHandler = handler.NewServiceHandler(c.ServiceUseCase)
}

// initMiddlewares initializes all middlewares
func (c *Container) initMiddlewares() {
	c.AuditMiddleware = middleware.NewAuditMiddleware(c.Logger)
}

// setupRouter sets up the Fiber application with routes
func (c *Container) setupRouter() {
	c.App = httpInterface.SetupRouter(
		c.ServiceGroupHandler,
		c.ServiceHandler,
		c.AuditMiddleware,
	)
}
