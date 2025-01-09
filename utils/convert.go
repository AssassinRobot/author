package convert

import (
	"strconv"

	"github.com/AssassinRobot/author/graph/model"
	dbModel "github.com/AssassinRobot/author/internal/model"
)

func ConvertAuthorWithoutBooks(dbAuthor *dbModel.Author) *model.Author {
	graphAuthor := new(model.Author)

	graphAuthor.ID = IntToString(dbAuthor.ID)
	graphAuthor.Name = dbAuthor.Name
	graphAuthor.Born = dbAuthor.Born
	graphAuthor.Died = dbAuthor.Died

	return graphAuthor
}

func ConvertAuthorsWithoutBooks(dbAuthors []*dbModel.Author) []*model.Author {
	var graphAuthors = []*model.Author{}
	for _, dbAuthor := range dbAuthors {
		graphAuthor := ConvertAuthorWithoutBooks(dbAuthor)
		graphAuthors = append(graphAuthors, graphAuthor)
	}

	return graphAuthors
}

func ConvertAuthorWithBooks(dbAuthor *dbModel.Author) *model.Author {
	graphAuthor := ConvertAuthorWithoutBooks(dbAuthor)
	graphBooks := ConvertBooksWithoutAuthors(dbAuthor.Books)

	graphAuthor.Books = graphBooks

	return graphAuthor
}

func ConvertAuthorsWithBooks(dbAuthors []*dbModel.Author) []*model.Author {
	var graphAuthors = []*model.Author{}
	for _, dbAuthor := range dbAuthors {
		graphAuthor := ConvertAuthorWithBooks(dbAuthor)
		graphAuthors = append(graphAuthors, graphAuthor)
	}

	return graphAuthors
}

func ConvertBooksWithoutAuthors(dbBooks []*dbModel.Book) []*model.Book {
	var graphBooks = []*model.Book{}
	for _, dbBook := range dbBooks {
		graphBook := ConvertBookWithoutAuthors(dbBook)
		graphBooks = append(graphBooks, graphBook)
	}

	return graphBooks
}

func ConvertBookWithoutAuthors(dbBook *dbModel.Book) *model.Book {
	graphBook := new(model.Book)
	graphBook.ID = IntToString(dbBook.ID)
	graphBook.Name = dbBook.Name
	graphBook.Pages = int32(dbBook.Pages)
	graphBook.Publication = int32(dbBook.Publication)

	graphGenres := ConvertGenresWithoutBooks(dbBook.Genres)
	graphBook.Genres = graphGenres

	graphLanguage := ConvertLanguageWithoutBooks(dbBook.Language)
	graphBook.Language = graphLanguage

	return graphBook
}

func ConvertBooksWithAuthors(dbBooks []*dbModel.Book) []*model.Book {
	var graphBooks = []*model.Book{}
	for _, dbBook := range dbBooks {
		graphBook := ConvertBookWithAuthors(dbBook)
		graphBooks = append(graphBooks, graphBook)
	}

	return graphBooks
}

func ConvertBookWithAuthors(dbBook *dbModel.Book) *model.Book {
	graphBook := ConvertBookWithoutAuthors(dbBook)

	var graphAuthors = []*model.Author{}
	for _, dbAuthor := range dbBook.Authors {
		graphAuthor := ConvertAuthorWithoutBooks(dbAuthor)
		graphAuthors = append(graphAuthors, graphAuthor)
	}
	graphBook.Authors = graphAuthors

	return graphBook
}

func ConvertGenresWithoutBooks(dbGenres []*dbModel.Genre) []*model.Genre {
	graphGenres := []*model.Genre{}
	for _, dbGenre := range dbGenres {
		graphGenre := ConvertGenreWithoutBooks(dbGenre)
		graphGenres = append(graphGenres, graphGenre)
	}

	return graphGenres
}

func ConvertGenreWithoutBooks(dbGenre *dbModel.Genre) *model.Genre {
	graphGenre := new(model.Genre)
	graphGenre.ID = IntToString(dbGenre.ID)
	graphGenre.Name = dbGenre.Name

	return graphGenre
}

func ConvertGenresWithBooks(dbGenres []*dbModel.Genre) []*model.Genre {
	graphGenres := []*model.Genre{}
	for _, dbGenre := range dbGenres {
		graphGenre := ConvertGenreWithBooks(dbGenre)
		graphGenres = append(graphGenres, graphGenre)
	}

	return graphGenres
}

func ConvertGenreWithBooks(dbGenre *dbModel.Genre) *model.Genre {
	graphGenre := ConvertGenreWithoutBooks(dbGenre)

	graphBooks := ConvertBooksWithAuthors(dbGenre.Books)
	graphGenre.Books = graphBooks

	return graphGenre
}

func ConvertLanguageWithoutBooks(dbLanguage *dbModel.Language) *model.Language {
	graphLanguage := new(model.Language)

	graphLanguage.ID = IntToString(dbLanguage.ID)
	graphLanguage.Name = dbLanguage.Name

	return graphLanguage
}

func ConvertLanguageWithBooks(dbLanguage *dbModel.Language) *model.Language {
	graphLanguage := ConvertLanguageWithoutBooks(dbLanguage)
	
	graphBooks := ConvertBooksWithAuthors(dbLanguage.Books)
	graphLanguage.Books = graphBooks

	return graphLanguage
}

func ConvertLanguagesWithBooks(dbLanguages []*dbModel.Language) []*model.Language {
	graphLanguages := []*model.Language{}
	for _, dbLanguage := range dbLanguages {
		graphLanguage := ConvertLanguageWithBooks(dbLanguage)
		graphLanguages = append(graphLanguages, graphLanguage)
	}

	return graphLanguages
}

func StringToInt(ID string) int {
	intID, err := strconv.Atoi(ID)
	if err != nil {
		panic(err)
	}

	return intID
}

func IntToString(ID int) string {
	stringID := strconv.Itoa(ID)
	return stringID
}
