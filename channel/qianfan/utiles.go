package qianfan

import (
	"context"

	"github.com/mitchellh/mapstructure"

)

// 转换任意对象成 map
func dumpToMap(input interface{}) (map[string]interface{}, error) {
	target := map[string]interface{}{}
	err := mapstructure.Decode(input, &target)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func runWithContext(ctx context.Context, fn func()) error {
	c := make(chan struct{}, 1)
	go func() {
		fn()
		c <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-c:
		return nil
	}
}