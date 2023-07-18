package gooption

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Run("Option is Some()", func(t *testing.T) {
		s := Some(1)
		r := s.Get()
		assert.Equal(t, r, 1)
	})

	t.Run("Option is None()", func(t *testing.T) {
		// Make sure we catch if we panic (we should panic here)
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		s := None[int]()
		r := s.Get()
		assert.Equal(t, r, 2)
	})
}

func TestGetOrElse(t *testing.T) {
	t.Run("Option is Some()", func(t *testing.T) {
		s := Some(1)
		r := s.GetOrElse(2)
		assert.Equal(t, r, 1)
	})

	t.Run("Option is None()", func(t *testing.T) {
		s := None[int]()
		r := s.GetOrElse(2)
		assert.Equal(t, r, 2)
	})
}

func TestIsDefined(t *testing.T) {
	t.Run("Option is Some()", func(t *testing.T) {
		s := Some(1)
		r := s.IsDefined()
		assert.True(t, r)
	})

	t.Run("Option is None()", func(t *testing.T) {
		s := None[int]()
		r := s.IsDefined()
		assert.False(t, r)
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("Option is Some()", func(t *testing.T) {
		s := Some(1)
		r := s.IsEmpty()
		assert.False(t, r)
	})

	t.Run("Option is None()", func(t *testing.T) {
		s := None[int]()
		r := s.IsEmpty()
		assert.True(t, r)
	})
}

func TestMap(t *testing.T) {
	addTwo := func(a int) int {
		return a + 2
	}
	t.Run("Option is Some()", func(t *testing.T) {
		s := Some(1)

		r := Map(s, addTwo)
		assert.Equal(t, r, Some(3))
	})

	t.Run("Option is None()", func(t *testing.T) {
		s := None[int]()
		r := Map(s, addTwo)
		assert.Equal(t, r, None[int]())
	})
}

func TestFlatMap(t *testing.T) {
	addTwoOnlyIfDivisibleByTwo := func(a int) Option[int] {
		if a%2 == 0 {
			return Some(a + 2)
		} else {
			return None[int]()
		}
	}
	t.Run("Option is Some()", func(t *testing.T) {
		s := Some(2)

		r := FlatMap(s, addTwoOnlyIfDivisibleByTwo)
		assert.Equal(t, r, Some(4))
	})

	t.Run("Option is None()", func(t *testing.T) {
		s := None[int]()
		r := FlatMap(s, addTwoOnlyIfDivisibleByTwo)
		assert.Equal(t, r, None[int]())
	})
}

type testType struct {
	Name  string         `json:"name"`
	Phone Option[string] `json:"phone"`
}

func TestMarshalJSON(t *testing.T) {
	t.Run("Option is Some()", func(t *testing.T) {
		s := Some(1)
		json, err := s.MarshalJSON()

		assert.Empty(t, err)
		assert.Equal(t, string(json), "1")
	})

	t.Run("Option is None()", func(t *testing.T) {
		s := None[int]()
		json, err := s.MarshalJSON()
		assert.Empty(t, err)
		assert.Empty(t, json)
	})

	t.Run("Complex Type with none value", func(t *testing.T) {

		tType := testType{
			Name:  "Calvin",
			Phone: None[string](),
		}
		bytes, err := json.Marshal(tType)
		expectedJson := `{"name":"Calvin", "phone": null}`
		assert.Empty(t, err)
		assert.JSONEq(t, string(bytes), expectedJson)
	})

	t.Run("Complex Type with some value", func(t *testing.T) {

		tType := testType{
			Name:  "Calvin",
			Phone: Some("111-111-1111"),
		}
		bytes, err := json.Marshal(tType)

		expectedJson := `{"name":"Calvin", "phone": "111-111-1111"}`
		assert.Empty(t, err)
		assert.JSONEq(t, string(bytes), expectedJson)
	})
}

func TestUnmarshalJSON(t *testing.T) {

	t.Run("Complex Type with none value", func(t *testing.T) {

		expectedJson := `{"name":"Calvin", "phone": null}`
		var gotType testType
		err := json.Unmarshal([]byte(expectedJson), &gotType)

		assert.Empty(t, err)
		expectedType := testType{
			Name:  "Calvin",
			Phone: None[string](),
		}
		assert.Equal(t, expectedType, gotType)
	})

	t.Run("Complex Type with some value", func(t *testing.T) {

		expectedJson := `{"name":"Calvin", "phone": "111-111-1111"}`
		var gotType testType
		err := json.Unmarshal([]byte(expectedJson), &gotType)

		assert.Empty(t, err)
		expectedType := testType{
			Name:  "Calvin",
			Phone: Some("111-111-1111"),
		}
		assert.Equal(t, expectedType, gotType)
	})
}
