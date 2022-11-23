package gateway

import (
	"context"
	"fmt"
)

// Repository interfaceの実装
type SampelRepository struct{}

// 本来ならRDBとかに接続するコードを書くことになるが、今回は適当に文字列を返す。
func (r *SampelRepository) Get(ctx context.Context, id string) (data string, err error) {
	return fmt.Sprintf("data of %s", id), nil
}
