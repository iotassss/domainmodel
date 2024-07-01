package domain

import "fmt"

type ValidationError struct {
	Msg string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.Msg)
}

type NotFoundError struct {
	Msg string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found error: %s", e.Msg)
}

type ServerError struct {
	Msg string
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("server error: %s", e.Msg)
}

type ConflictError struct {
	Msg string
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("conflict error: %s", e.Msg)
}

// 現在400、404、500のエラーを返すことができる
// TODO: 他のエラーコードも追加する
