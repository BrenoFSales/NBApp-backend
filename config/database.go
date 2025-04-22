package config

import (
	"fmt"
	"log"
	"nbapp/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("nbapp.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	db.AutoMigrate(&models.User{}, &models.Team{})

	DB = db
}

// Essa função apenas foi usada para inserir os times no banco...
func InsertTeams() {
	teamsData := []models.Team{
		{Conference: "East", Division: "Southeast", City: "Atlanta", Name: "Hawks", FullName: "Atlanta Hawks", Abbreviation: "ATL"},
		{Conference: "East", Division: "Atlantic", City: "Boston", Name: "Celtics", FullName: "Boston Celtics", Abbreviation: "BOS"},
		{Conference: "East", Division: "Atlantic", City: "Brooklyn", Name: "Nets", FullName: "Brooklyn Nets", Abbreviation: "BKN"},
		{Conference: "East", Division: "Southeast", City: "Charlotte", Name: "Hornets", FullName: "Charlotte Hornets", Abbreviation: "CHA"},
		{Conference: "East", Division: "Central", City: "Chicago", Name: "Bulls", FullName: "Chicago Bulls", Abbreviation: "CHI"},
		{Conference: "East", Division: "Central", City: "Cleveland", Name: "Cavaliers", FullName: "Cleveland Cavaliers", Abbreviation: "CLE"},
		{Conference: "West", Division: "Southwest", City: "Dallas", Name: "Mavericks", FullName: "Dallas Mavericks", Abbreviation: "DAL"},
		{Conference: "West", Division: "Northwest", City: "Denver", Name: "Nuggets", FullName: "Denver Nuggets", Abbreviation: "DEN"},
		{Conference: "East", Division: "Central", City: "Detroit", Name: "Pistons", FullName: "Detroit Pistons", Abbreviation: "DET"},
		{Conference: "West", Division: "Pacific", City: "Golden State", Name: "Warriors", FullName: "Golden State Warriors", Abbreviation: "GSW"},
		{Conference: "West", Division: "Southwest", City: "Houston", Name: "Rockets", FullName: "Houston Rockets", Abbreviation: "HOU"},
		{Conference: "East", Division: "Central", City: "Indiana", Name: "Pacers", FullName: "Indiana Pacers", Abbreviation: "IND"},
		{Conference: "West", Division: "Pacific", City: "LA", Name: "Clippers", FullName: "LA Clippers", Abbreviation: "LAC"},
		{Conference: "West", Division: "Pacific", City: "Los Angeles", Name: "Lakers", FullName: "Los Angeles Lakers", Abbreviation: "LAL"},
		{Conference: "West", Division: "Southwest", City: "Memphis", Name: "Grizzlies", FullName: "Memphis Grizzlies", Abbreviation: "MEM"},
		{Conference: "East", Division: "Southeast", City: "Miami", Name: "Heat", FullName: "Miami Heat", Abbreviation: "MIA"},
		{Conference: "East", Division: "Central", City: "Milwaukee", Name: "Bucks", FullName: "Milwaukee Bucks", Abbreviation: "MIL"},
		{Conference: "West", Division: "Northwest", City: "Minnesota", Name: "Timberwolves", FullName: "Minnesota Timberwolves", Abbreviation: "MIN"},
		{Conference: "West", Division: "Southwest", City: "New Orleans", Name: "Pelicans", FullName: "New Orleans Pelicans", Abbreviation: "NOP"},
		{Conference: "East", Division: "Atlantic", City: "New York", Name: "Knicks", FullName: "New York Knicks", Abbreviation: "NYK"},
		{Conference: "West", Division: "Northwest", City: "Oklahoma City", Name: "Thunder", FullName: "Oklahoma City Thunder", Abbreviation: "OKC"},
		{Conference: "East", Division: "Southeast", City: "Orlando", Name: "Magic", FullName: "Orlando Magic", Abbreviation: "ORL"},
		{Conference: "East", Division: "Atlantic", City: "Philadelphia", Name: "76ers", FullName: "Philadelphia 76ers", Abbreviation: "PHI"},
		{Conference: "West", Division: "Pacific", City: "Phoenix", Name: "Suns", FullName: "Phoenix Suns", Abbreviation: "PHX"},
		{Conference: "West", Division: "Northwest", City: "Portland", Name: "Trail Blazers", FullName: "Portland Trail Blazers", Abbreviation: "POR"},
		{Conference: "West", Division: "Pacific", City: "Sacramento", Name: "Kings", FullName: "Sacramento Kings", Abbreviation: "SAC"},
		{Conference: "West", Division: "Southwest", City: "San Antonio", Name: "Spurs", FullName: "San Antonio Spurs", Abbreviation: "SAS"},
		{Conference: "East", Division: "Atlantic", City: "Toronto", Name: "Raptors", FullName: "Toronto Raptors", Abbreviation: "TOR"},
		{Conference: "West", Division: "Northwest", City: "Utah", Name: "Jazz", FullName: "Utah Jazz", Abbreviation: "UTA"},
		{Conference: "East", Division: "Southeast", City: "Washington", Name: "Wizards", FullName: "Washington Wizards", Abbreviation: "WAS"},
	}

	result := DB.Create(&teamsData)
	if result.Error != nil {
		log.Fatalf("Falha ao inserir o dataset inicial dos times da NBA: %v", result.Error)
	}

	fmt.Printf("Inseridos %d times com sucesso!\n", result.RowsAffected)
}
