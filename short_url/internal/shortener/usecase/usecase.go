package usecase

import (
	"ShortURL/internal/logging"
	"ShortURL/internal/shortener"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

type useCase struct {
	repo   shortener.Repo
	logger *logging.Logger
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
	if err = uc.repo.AddShortURL(ctx, url, shortURL); err != nil {
		return "", err
	}
	return shortURL, nil
}

func (uc *useCase) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	url, err := uc.repo.GetURL(ctx, shortURL)
	if err != nil {
		return "", err
	}
	if url == "" {
		err = fmt.Errorf(URLDoesNotExist)
		return "", err
	}
	return url, nil
}

func (uc *useCase) generateUniqueHash(ctx context.Context, url string) (string, error) {
	for counter := 0; ; counter++ {
		strWithCounter := fmt.Sprintf("%s%d", url, counter)
		hash := uc.generateHash(strWithCounter)
		if len(hash) < hashLength {
			continue
		}
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

func (uc *useCase) generateHash(str string) string {
	hashBytes := sha256.Sum256([]byte(str))
	hashString := base64.RawURLEncoding.EncodeToString(hashBytes[:])
	tmpHash := uc.filterCharacters(hashString, charset)
	if len(tmpHash) < hashLength {
		return ""
	}
	hash := tmpHash[:hashLength]
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

func NewUseCase(repo shortener.Repo, log *logging.Logger, inMemory bool) shortener.UseCase {
	uc := &useCase{
		repo:   repo,
		logger: log,
	}
	return uc
}
