package models

import "go.mongodb.org/mongo-driver/v2/bson"

// files belonging to the same package can share the functions, variables, and types defined in that package.
// Only exported items (starting with capital letter) can be accessed from other packages.
// Unexported items (starting with lowercase letter) are only accessible within the same package. They are private to the package.

type Genre struct {
	GenreID   int    `bson:"genre_id" json:"genre_id" validate:"required"`
	GenreName string `bson:"genre_name" json:"genre_name" validate:"required,min=2,max=100"`
}

type Ranking struct {
	RankingValue int    `bson:"ranking_value" json:"ranking_value" validate:"required"`
	RankingName  string `bson:"ranking_name" json:"ranking_name" validate:"oneof=Excellent Good Okay Bad Terrible"`
}

type Movie struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ImdbID     string        `bson:"imdb_id" json:"imdb_id" validate:"required"`
	Title      string        `bson:"title" json:"title" validate:"required,min=2,max=500"`
	PosterPath string        `bson:"poster_path" json:"poster_path" validate:"required,url"`
	YoutubeID  string        `bson:"youtube_id" json:"youtube_id" validate:"required"`
	Genre      []Genre       `bson:"genre" json:"genre" validate:"required,dive"`
	// this dive keyword ensures that the Genre struct will also be validated
	AdminReview string    `bson:"admin_review" json:"admin_review" validate:"required"`
	Ranking     []Ranking `bson:"ranking" json:"ranking" validate:"required"`
}
