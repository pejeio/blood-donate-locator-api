package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/types"
)

type PageLimitOffset struct {
	Page   int
	Limit  int
	Offset int
}

func GetPaginationQueryParams(c *fiber.Ctx) *PageLimitOffset {
	q := new(types.PaginationRequest)
	c.QueryParser(q)

	if q.Limit == "" {
		q.Limit = "10"
	}
	if q.Page == "" {
		q.Page = "1"
	}

	intPage, _ := strconv.Atoi(q.Page)
	intLimit, _ := strconv.Atoi(q.Limit)

	offset := (intPage - 1) * intLimit
	return &PageLimitOffset{
		Page:   intPage,
		Limit:  intLimit,
		Offset: offset,
	}
}
