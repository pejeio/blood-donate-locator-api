package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pejeio/blood-donate-locator-api/internal/types"
	"github.com/rs/zerolog/log"
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

	// Parse query parameters into q
	if err := c.QueryParser(q); err != nil {
		log.Error().Err(err)
		return nil, err
	}

	// Set default values if Limit or Offset are empty
	if q.Limit == "" {
		q.Limit = strconv.Itoa(DefaultLimit)
	}
	if q.Offset == "" {
		q.Offset = strconv.Itoa(DefaultOffset)
	}

	// Parse Limit and Offset to integers
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
