package directive

import (
	"context"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/99designs/gqlgen/graphql"
	"github.com/hizzuu/grpc-example-bff/gen/graph/gqlerr"
)

func (d *directive) Constraint(
	ctx context.Context,
	obj interface{},
	next graphql.Resolver,
	label string,
	notEmpty *bool,
	notBlank *bool,
	pattern *string,
	min *int64,
	max *int64,
) (interface{}, error) {
	i, _ := next(ctx)
	switch v := i.(type) {
	case *int64:
		if v != nil {
			if err := checkInt(*v, label, min, max); err != nil {
				return nil, err
			}
		}
	case int64:
		if err := checkInt(v, label, min, max); err != nil {
			return nil, err
		}
	case *string:
		if v != nil {
			if err := checkString(*v, label, notEmpty, notBlank, pattern, min, max); err != nil {
				return nil, err
			}
		}
	case string:
		if err := checkString(v, label, notEmpty, notBlank, pattern, min, max); err != nil {
			return nil, err
		}
	}
	return i, nil
}

func checkInt(v int64, label string, min *int64, max *int64) error {
	if min != nil && v < *min {
		return gqlerr.UserInputError("%sは%d以上で入力してください。", label, *min)
	}
	if max != nil && v > *max {
		gqlerr.UserInputError("%sは%d以下で入力してください。", label, *max)
	}
	return nil
}

func checkString(v string, label string, notEmpty *bool, notBlank *bool, pattern *string, min *int64, max *int64) error {
	if notEmpty != nil && v == "" {
		return gqlerr.UserInputError("%sを入力してください。", label)
	}
	if notBlank != nil {
		if strings.TrimSpace(v) == "" {
			return gqlerr.UserInputError("%sを入力してください。", label)
		}
	}
	if min != nil && utf8.RuneCountInString(v) < int(*min) {
		return gqlerr.UserInputError("%sは%d以上で入力してください。", label, *min)
	}
	if max != nil && utf8.RuneCountInString(v) > int(*max) {
		return gqlerr.UserInputError("%sは%d文字以下で入力してください。", label, *max)
	}
	if pattern != nil {
		if ok, _ := regexp.MatchString(*pattern, v); !ok {
			return gqlerr.UserInputError("%sのフォーマットが間違っています。")
		}
	}
	return nil
}
