package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisements/models"
)

// ContainsAdvertisementByTitle проверяет, что в списке есть айтем с указанным названием
func ContainsAdvertisementByTitle(t *testing.T, items []models.Advertisement, title string) {
	t.Helper()

	titles := make([]string, 0, len(items))
	for _, item := range items {
		titles = append(titles, item.Title)
	}

	require.Contains(t, titles, title, "Среди результатов поиска должно быть объявление %q", title)
}
