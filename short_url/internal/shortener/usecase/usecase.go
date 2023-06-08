package usecase

import (
	"ShortURL/internal/logging"
	shortener2 "ShortURL/internal/shortener"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

type useCase struct {
	repo   shortener2.Repo
	logger *logging.Logger

	uRLHash  map[string]string
	InMemory bool
}

const (
	hashLength = 10
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
)

const (
	emptyURL        = "empty URL"
	URLDoesNotExist = "URL does not exist"
	exist           = "URL all ready exist"
)

func (uc *useCase) MakeURLShorter(ctx context.Context, url string) (string, error) {
	if url == "" {
		return "", fmt.Errorf(emptyURL)
	}
	shortURL, err := uc.generateUniqueHash(ctx, url)
	if err != nil {
		if err.Error() != exist {
			return "", err
		} else {
			return shortURL, nil
		}
	}
	if uc.InMemory {
		uc.uRLHash[shortURL] = url
	} else {
		if err = uc.repo.AddShortURL(ctx, url, shortURL); err != nil {
			return "", err
		}
	}
	return shortURL, nil
}

func (uc *useCase) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	var url string
	if uc.InMemory {
		var ok bool
		url, ok = uc.uRLHash[shortURL]
		if !ok {
			err := fmt.Errorf(URLDoesNotExist)
			return "", err
		}
		return url, nil
	} else {
		var err error
		url, err = uc.repo.GetURL(ctx, shortURL)
		if err != nil {
			return "", err
		}
		if url == "" {
			err = fmt.Errorf(URLDoesNotExist)
			return "", err
		}
	}
	return url, nil
}

func (uc *useCase) generateUniqueHash(ctx context.Context, url string) (string, error) {
	for counter := 0; ; counter++ {
		strWithCounter := fmt.Sprintf("%s%d", url, counter)
		hash := uc.generateHash(strWithCounter)
		if uc.InMemory {
			retURL, ok := uc.uRLHash[hash]
			if !ok {
				return hash, nil
			}
			if retURL == url {
				return hash, fmt.Errorf(exist)
			}
		} else {
			retURL, err := uc.repo.GetURL(ctx, hash)
			if err != nil {
				return "", err
			}
			if retURL == "" {
				return hash, nil
			}
			if retURL == url {
				return hash, fmt.Errorf(exist)
			}
		}
	}
}

func (uc *useCase) generateHash(str string) string {
	hashBytes := sha256.Sum256([]byte(str))
	hashString := base64.RawURLEncoding.EncodeToString(hashBytes[:])
	hash := uc.filterCharacters(hashString, charset)[:hashLength]
	return hash
}

func (uc *useCase) filterCharacters(str string, allowedChars string) string {
	var filtered string
	for _, char := range str {
		if strings.ContainsRune(allowedChars, char) {
			filtered += string(char)
		}
	}
	return filtered
}

func NewUseCase(repo shortener2.Repo, log *logging.Logger, inMemory bool) shortener2.UseCase {
	uc := &useCase{
		repo:     repo,
		logger:   log,
		uRLHash:  map[string]string{},
		InMemory: inMemory,
	}
	return uc
}
