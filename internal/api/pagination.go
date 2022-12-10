package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/types"
)

type PageLimitOffset struct {
	Limit  int
	Offset int
}

func GetPaginationQueryParams(c *fiber.Ctx) *PageLimitOffset {
	q := new(types.PaginationRequest)
	c.QueryParser(q)

	if q.Limit == "" {
		q.Limit = "10"
	}
	if q.Offset == "" {
		q.Offset = "0"
	}

	intOffset, _ := strconv.Atoi(q.Offset)
	intLimit, _ := strconv.Atoi(q.Limit)

	return &PageLimitOffset{
		Limit:  intLimit,
		Offset: intOffset,
	}
}
