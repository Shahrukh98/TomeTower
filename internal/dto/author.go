package dto

import (
	"tometower/internal/entity"
	"tometower/pkg/util"
)

type CreateAuthorDto struct {
	Name        string `json:"name"`
	PhotoUrl    string `json:"photoUrl"`
	Nationality string `json:"nationality"`
	DateOfBirth string `json:"dateOfBirth"`
	DateOfDeath string `json:"dateOfDeath"`
}

type AuthorDataDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhotoUrl    string `json:"photoUrl"`
	Nationality string `json:"nationality"`
	DateOfBirth string `json:"dateOfBirth"`
	DateOfDeath string `json:"dateOfDeath"`
}

func MapCreateAuthorDtoToAuthor(createAuthorDto *CreateAuthorDto) (*entity.Author, error) {
	var author entity.Author

	author.Name = createAuthorDto.Name
	author.PhotoUrl = createAuthorDto.PhotoUrl
	author.Nationality = createAuthorDto.Nationality

	err := util.IsDateValid(createAuthorDto.DateOfBirth)
	if err != nil {
		return nil, err
	}
	author.DateOfBirth = createAuthorDto.DateOfBirth

	if len(createAuthorDto.DateOfDeath) == 0 {
		author.DateOfDeath = ""
	} else {
		err := util.IsDateValid(createAuthorDto.DateOfDeath)
		if err != nil {
			return nil, err
		}
		author.DateOfDeath = createAuthorDto.DateOfDeath
	}
	return &author, nil
}

func MapAuthorToAuthorDataDto(author entity.Author) (AuthorDataDto, error) {
	var authorDataDto AuthorDataDto

	authorDataDto.Id = author.Id
	authorDataDto.Name = author.Name
	authorDataDto.PhotoUrl = author.PhotoUrl
	authorDataDto.Nationality = author.Nationality
	authorDataDto.DateOfBirth = author.DateOfBirth
	authorDataDto.DateOfDeath = author.DateOfDeath
	return authorDataDto, nil
}
