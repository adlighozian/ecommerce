package main

import (
	"log"
	"time"
	"user-consumer-go/config"
	"user-consumer-go/handler"
	"user-consumer-go/helper/logging"
	"user-consumer-go/package/db"
	"user-consumer-go/package/rmq"
	"user-consumer-go/repository"
	"user-consumer-go/service"
)

func main() {
	config, errConf := config.LoadConfig()
	if errConf != nil {
		log.Fatalf("load config err:%s", errConf)
	}

	logger := logging.New(config.Debug)

	sqlDB, errDB := db.NewGormDB(config.Debug, config.Database.Driver, config.Database.URL)
	if errDB != nil {
		logger.Fatal().Err(errDB).Msg("db failed to connect")
	}
	logger.Debug().Msg("db connected")

	rmq, errRmq := rmq.NewRabbitMQ(config.RabbitMQ.URL)
	if errRmq != nil {
		logger.Fatal().Err(errRmq).Msg("rabbitmq failed to connect")
	}
	logger.Debug().Msg("rabbitmq connected")

	defer func() {
		errDBC := sqlDB.Close()
		if errDBC != nil {
			logger.Fatal().Err(errDBC).Msg("db failed to closed")
		}
		logger.Debug().Msg("db closed")

		errRmqC := rmq.Shutdown()
		if errRmqC != nil {
			logger.Fatal().Err(errRmqC).Msg("rabbitmq failed to closed")
		}
		logger.Debug().Msg("rabbitmq closed")
	}()

	userRepo := repository.NewUserRepository(sqlDB.SQLDB)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(rmq, logger, userSvc)

	time.Sleep(5 * time.Second)

	var forever chan struct{}

	userHandler.Create()

	logger.Debug().Msg("[*] To exit press CTRL+C")
	<-forever
}
