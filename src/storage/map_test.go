package storage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMapDB(t *testing.T) {
	require := require.New(t)

	actual := NewMapDB()

	require.NotNil(actual)
}

func TestMapDBSet(t *testing.T) {
	t.Run("Empty Key", func(t *testing.T) {
		require := require.New(t)
		db := NewMapDB()

		actualErr := db.Set("", "value")

		require.NotNil(actualErr)

		actualStr := actualErr.Error()
		expectedStr := fmt.Sprintf(fieldNotSetErr, "key")

		require.Equal(expectedStr, actualStr)
	})

	t.Run("Empty Value", func(t *testing.T) {
		require := require.New(t)
		db := NewMapDB()

		actualErr := db.Set("key", "")

		require.NotNil(actualErr)

		actualStr := actualErr.Error()
		expectedStr := fmt.Sprintf(fieldNotSetErr, "value")

		require.Equal(expectedStr, actualStr)
	})
}

func TestMapDBGet(t *testing.T) {
	t.Parallel()
	t.Run("Empty Key", func(t *testing.T) {
		require := require.New(t)
		db := NewMapDB()

		_, actualErr := db.Get("")

		require.NotNil(actualErr)

		actualStr := actualErr.Error()
		expectedStr := fmt.Sprintf(fieldNotSetErr, "key")

		require.Equal(expectedStr, actualStr)
	})

	t.Run("Non Set", func(t *testing.T) {
		require := require.New(t)
		db := NewMapDB()
		key := "key"

		_, actualErr := db.Get(key)

		require.NotNil(actualErr)

		actualStr := actualErr.Error()
		expectedStr := fmt.Sprintf(noValueFoundErr, key)

		require.Equal(expectedStr, actualStr)

	})

	t.Run("Get", func(t *testing.T) {
		require := require.New(t)
		db := NewMapDB()
		key := "key"
		value := "value"

		err := db.Set(key, value)

		require.Nil(err)
	})
}
