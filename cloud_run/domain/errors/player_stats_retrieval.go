package errors

import "fmt"

type PlayerStatsRetrievalError struct {
	Err error
}

func NewPlayerStatsRetrievalError(err error) *PlayerStatsRetrievalError {
	return &PlayerStatsRetrievalError{
		Err: err,
	}
}

func (e *PlayerStatsRetrievalError) Error() string {
	return fmt.Sprintf("Player stats retrieval error: %s", e.Err)
}

func (e *PlayerStatsRetrievalError) Unwrap() error {
	return e.Err
}
