package model

type Card struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ManaCost    string  `json:"manaCost"`
	Type        string  `json:"type"`
	Rarity      string  `json:"rarity"`
	Set         string  `json:"set"`
	Power       *int    `json:"power,omitempty"`
	Toughness   *int    `json:"toughness,omitempty"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
} 