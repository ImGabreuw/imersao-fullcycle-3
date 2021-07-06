package usecase

import (
	"encoding/json"
	"github.com/imgabreuw/codebank/domain"
	"github.com/imgabreuw/codebank/dto"
	"github.com/imgabreuw/codebank/infrastructure/kafka"
)

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
	KafkaProducer         kafka.KafkaProducer
}

func NewUseCaseTransaction(transactionRepository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{
		TransactionRepository: transactionRepository,
	}
}

func (useCaseTransaction UseCaseTransaction) ProcessTransaction(transactionDto dto.Transaction) (domain.Transaction, error) {
	creditCard := useCaseTransaction.hydrateCreditCard(transactionDto)

	creditCardBalanceAndLimit, err := useCaseTransaction.TransactionRepository.GetCreditCard(*creditCard)

	if err != nil {
		return domain.Transaction{}, err
	}

	creditCard.ID = creditCardBalanceAndLimit.ID
	creditCard.Limit = creditCardBalanceAndLimit.Limit
	creditCard.Balance = creditCardBalanceAndLimit.Balance

	transaction := useCaseTransaction.newTransaction(transactionDto, creditCardBalanceAndLimit)
	transaction.ProcessAndValidate(creditCard)

	err = useCaseTransaction.TransactionRepository.SaveTransaction(*transaction, *creditCard)

	if err != nil {
		return domain.Transaction{}, err
	}

	transactionDto.ID = transaction.ID
	transactionDto.CreatedAt = transaction.CreatedAt

	transactionJson, err := json.Marshal(transactionDto)

	if err != nil {
		return domain.Transaction{}, err
	}

	err = useCaseTransaction.KafkaProducer.Publish(
		string(transactionJson),
		"payment",
	)

	if err != nil {
		return domain.Transaction{}, err
	}

	return *transaction, nil
}

func (useCaseTransaction UseCaseTransaction) hydrateCreditCard(transactionDto dto.Transaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()

	creditCard.Name = transactionDto.Name
	creditCard.Number = transactionDto.Number
	creditCard.ExpirationMonth = transactionDto.ExpirationMonth
	creditCard.ExpirationYear = transactionDto.ExpirationYear
	creditCard.ExpirationYear = transactionDto.ExpirationYear
	creditCard.CVV = transactionDto.CVV

	return creditCard
}

func (useCaseTransaction UseCaseTransaction) newTransaction(transactionDto dto.Transaction, creditCard domain.CreditCard) *domain.Transaction {
	transaction := domain.NewTransaction()

	transaction.CreditCardId = transactionDto.ID
	transaction.Amount = transactionDto.Amount
	transaction.Store = transactionDto.Store
	transaction.Description = transactionDto.Description
	transaction.CreatedAt = transactionDto.CreatedAt

	return transaction
}
