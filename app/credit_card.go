package app

import (
	"github.com/scmn-dev/core/db"
	"github.com/scmn-dev/core/model"
)

// CreateCreditCard creates a new credit card and saves it to the store
func CreateCreditCard(s db.Store, dto *model.CreditCardDTO, schema string) (*model.CreditCard, error) {
	rawModel := model.ToCreditCard(dto)
	encModel := EncryptModel(rawModel)

	createdCreditCard, err := s.CreditCards().Create(encModel.(*model.CreditCard), schema)
	if err != nil {
		return nil, err
	}

	return createdCreditCard, nil
}

// UpdateCreditCard updates the credit card with the dto and applies the changes in the store
func UpdateCreditCard(s db.Store, creditCard *model.CreditCard, dto *model.CreditCardDTO, schema string) (*model.CreditCard, error) {
	rawModel := model.ToCreditCard(dto)
	encModel := EncryptModel(rawModel).(*model.CreditCard)

	creditCard.CardName = encModel.CardName
	creditCard.CardholderName = encModel.CardholderName
	creditCard.Type = encModel.Type
	creditCard.Number = encModel.Number
	creditCard.VerificationNumber = encModel.VerificationNumber
	creditCard.ExpiryDate = encModel.ExpiryDate

	updatedCreditCard, err := s.CreditCards().Update(creditCard, schema)
	if err != nil {
		return nil, err
	}

	return updatedCreditCard, nil
}
