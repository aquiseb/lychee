package helpers

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

func IDToUint(i graphql.ID) (uint, error) {
	r, err := strconv.ParseInt(string(i), 10, 32)
	if err != nil {
		return 0, errors.Wrap(err, "GqlIDToUint")
	}

	return uint(r), nil
}

func Int32P(i uint) *int32 {
	r := int32(i)
	return &r
}

func BoolP(b bool) *bool {
	return &b
}

func IDP(id uint) *graphql.ID {
	r := graphql.ID(fmt.Sprint(id))
	return &r
}

func CompareUnorderedSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	a_copy := make([]string, len(a))
	b_copy := make([]string, len(b))

	copy(a_copy, a)
	copy(b_copy, b)

	sort.Strings(a_copy)
	sort.Strings(b_copy)

	return reflect.DeepEqual(a_copy, b_copy)
}
