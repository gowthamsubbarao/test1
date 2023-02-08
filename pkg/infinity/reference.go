package infinity

import (
	"context"
	"errors"
	"strings"

	"github.com/yesoreyeram/grafana-infinity-datasource/pkg/models"
)

func UpdateQueryWithReferenceData(ctx context.Context, query models.Query, settings models.InfinitySettings) (models.Query, error) {
	if query.Source == "reference" {
		for _, item := range settings.ReferenceData {
			if strings.EqualFold(item.Name, query.RefName) {
				query.Source = "inline"
				query.Data = item.Data
				return query, nil
			}
		}
		return query, errors.New("error getting reference data. Either empty or not defined")
	}
	return query, nil
}
