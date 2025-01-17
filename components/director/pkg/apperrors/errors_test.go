package apperrors

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestIsNotFoundError(t *testing.T) {
	notFoundError := &notFoundError{}
	wrappedNotFoundError := errors.Wrap(notFoundError, "wrapped text")
	multiWrappedNotFoundError := errors.Wrap(wrappedNotFoundError, "multi wrapped")
	testErr := errors.New("test")

	testCases := []struct {
		Name           string
		Error          error
		expectedResult bool
	}{
		{
			Name:           "Unwrapped NotFound error",
			Error:          notFoundError,
			expectedResult: true,
		},
		{
			Name:           "Wrapped NotFound error",
			Error:          wrappedNotFoundError,
			expectedResult: true,
		},
		{
			Name:           "Multi wrapped NotFound error",
			Error:          multiWrappedNotFoundError,
			expectedResult: true,
		},
		{
			Name:           "Different error",
			Error:          testErr,
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, IsNotFoundError(testCase.Error))
		})
	}
}

func TestIsNotUnique(t *testing.T) {
	notUniqueError := &notUniqueError{}
	wrappedNotUniqueError := errors.Wrap(notUniqueError, "wrapped text")
	multiWrappedNotUniqueError := errors.Wrap(wrappedNotUniqueError, "multi wrapped")
	testErr := errors.New("test")

	testCases := []struct {
		Name           string
		Error          error
		expectedResult bool
	}{
		{
			Name:           "Unwrapped NotUnique error",
			Error:          notUniqueError,
			expectedResult: true,
		},
		{
			Name:           "Wrapped NotUnique error",
			Error:          wrappedNotUniqueError,
			expectedResult: true,
		},
		{
			Name:           "Multi wrapped NotUnique error",
			Error:          multiWrappedNotUniqueError,
			expectedResult: true,
		},
		{
			Name:           "Different error",
			Error:          testErr,
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, IsNotUnique(testCase.Error))
		})
	}
}

func TestKeyDoesNotExistError(t *testing.T) {
	keyDoesNotExistError := &keyDoesNotExistError{}
	wrappedKeyDoesNotExistError := errors.Wrap(keyDoesNotExistError, "wrapped text")
	multiWrappedKeyDoesNotExistError := errors.Wrap(wrappedKeyDoesNotExistError, "multi wrapped")
	testErr := errors.New("test")

	testCases := []struct {
		Name           string
		Error          error
		expectedResult bool
	}{
		{
			Name:           "Unwrapped KeyDoesNotExist error",
			Error:          keyDoesNotExistError,
			expectedResult: true,
		},
		{
			Name:           "Wrapped KeyDoesNotExist error",
			Error:          wrappedKeyDoesNotExistError,
			expectedResult: true,
		},
		{
			Name:           "Multi wrapped KeyDoesNotExist error",
			Error:          multiWrappedKeyDoesNotExistError,
			expectedResult: true,
		},
		{
			Name:           "Different error",
			Error:          testErr,
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, IsKeyDoesNotExist(testCase.Error))
		})
	}
}

func TestInvalidStringCastError(t *testing.T) {
	invalidStringCastError := &invalidStringCastError{}
	wrappedInvalidStringCastError := errors.Wrap(invalidStringCastError, "wrapped text")
	multiWrappedInvalidStringCastError := errors.Wrap(wrappedInvalidStringCastError, "multi wrapped")
	testErr := errors.New("test")

	testCases := []struct {
		Name           string
		Error          error
		expectedResult bool
	}{
		{
			Name:           "Unwrapped InvalidStringCast error",
			Error:          invalidStringCastError,
			expectedResult: true,
		},
		{
			Name:           "Wrapped InvalidStringCast error",
			Error:          wrappedInvalidStringCastError,
			expectedResult: true,
		},
		{
			Name:           "Multi wrapped InvalidStringCast error",
			Error:          multiWrappedInvalidStringCastError,
			expectedResult: true,
		},
		{
			Name:           "Different error",
			Error:          testErr,
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, IsInvalidCast(testCase.Error))
		})
	}
}
