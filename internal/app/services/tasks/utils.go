package tasks

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/bogi-lyceya-44/common/pkg/utils"
)

type idWithDependencies utils.Pair[int64, []int64]

func checkForSelfDependency(pairs []idWithDependencies) error {
	selfDependent := utils.Filter(
		pairs,
		func(pair idWithDependencies) bool {
			return slices.Contains(pair.Second, pair.First)
		},
	)

	if len(selfDependent) > 0 {
		ids := utils.Map(
			pairs,
			func(pair idWithDependencies) string {
				return strconv.Itoa(int(pair.First))
			},
		)

		return fmt.Errorf(
			"found self dependent tasks: %v",
			strings.Join(ids, ","),
		)
	}

	return nil
}
