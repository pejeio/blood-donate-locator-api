package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/types"
	log "github.com/sirupsen/logrus"
)

type PageLimitOffset struct {
	Limit  int
	Offset int
}

const (
	DefaultLimit  = 10
	DefaultOffset = 0
)

func GetPaginationQueryParams(c *fiber.Ctx) (*PageLimitOffset, error) {
	q := new(types.PaginationRequest)
	err := c.QueryParser(q)
	if err != nil {
		log.Println(err)
	}

	if q.Limit == "" {
		q.Limit = strconv.Itoa(DefaultLimit)
	}
	if q.Offset == "" {
		q.Offset = strconv.Itoa(DefaultOffset)
	}

	intOffset, err := strconv.Atoi(q.Offset)
	if err != nil {
		return nil, err
	}
	intLimit, err := strconv.Atoi(q.Limit)
	if err != nil {
		return nil, err
	}

	return &PageLimitOffset{
		Limit:  intLimit,
		Offset: intOffset,
	}, nil
}
