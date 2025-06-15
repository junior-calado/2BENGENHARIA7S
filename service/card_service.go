package service

import (
	"2BENGENHARIA7S/model"
)

func GetCards() ([]model.Card, error) {
	var cards []model.Card
	if err := DB.Find(&cards).Error; err != nil {
		return nil, err
	}
	return cards, nil
}

func AddCard(card model.Card) (model.Card, error) {
	if err := DB.Create(&card).Error; err != nil {
		return model.Card{}, err
	}
	return card, nil
}

func GetCardByID(id int) (model.Card, error) {
	var card model.Card
	if err := DB.First(&card, id).Error; err != nil {
		return model.Card{}, err
	}
	return card, nil
}

func UpdateCard(card model.Card) (model.Card, error) {
	if err := DB.Save(&card).Error; err != nil {
		return model.Card{}, err
	}
	return card, nil
}

func DeleteCard(id int) error {
	return DB.Delete(&model.Card{}, id).Error
} 