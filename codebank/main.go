package main

import (
	"database/sql"
	"fmt"
	"github.com/imgabreuw/codebank/infrastructure/kafka"
	"github.com/imgabreuw/codebank/infrastructure/repository"
	"github.com/imgabreuw/codebank/infrastructure/server"
	"github.com/imgabreuw/codebank/usecase"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db := setupDb()
	defer db.Close()

	producer := setupKafkaProducer()
	processTransactionUseCase := setupTransactionUsaCase(db, producer)
	serveGrpc(processTransactionUseCase)
}

func setupKafkaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer("host.docker.internal.9094")

	return producer
}

func setupTransactionUsaCase(db *sql.DB, producer kafka.KafkaProducer) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	useCase.KafkaProducer = producer

	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Erro connection to database")
	}

	return db
}

func serveGrpc(processTransactionUseCase usecase.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessTransactionUseCase = processTransactionUseCase
	fmt.Println("Rodando gRPC Server")
	grpcServer.Serve()
}

//creditCard := domain.NewCreditCard()
//
//creditCard.Number = "1234"
//creditCard.Name = "Wesley"
//creditCard.ExpirationMonth = 2021
//creditCard.ExpirationMonth = 7
//creditCard.CVV = 123
//creditCard.Limit = 1000
//creditCard.Balance = 0
//
//repo := repository.NewTransactionRepositoryDb(db)
//err := repo.CreateCreditCard(*creditCard)
//
//if err != nil {
//fmt.Println(err)
//}
