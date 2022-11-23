package repository

import "context"

// モックする対象のinterface
type Repository interface {
	// なにかIDを指定してデータを持ってくるメソッド
	Get(ctx context.Context, id string) (data string, err error)
}
