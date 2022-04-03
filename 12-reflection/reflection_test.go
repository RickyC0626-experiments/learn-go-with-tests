package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Person"},
			[]string{"Person"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Person", "Earth"},
			[]string{"Person", "Earth"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Person", 30},
			[]string{"Person"},
		},
		{
			"Nested fields",
			Person{
				"Person",
				Profile{30, "Earth"},
			},
			[]string{"Person", "Earth"},
		},
		{
			"Pointers to things",
			&Person{
				"Person",
				Profile{30, "Earth"},
			},
			[]string{"Person", "Earth"},
		},
		{
			"Slices",
			[]Profile{
				{30, "Earth"},
				{33, "Mars"},
			},
			[]string{"Earth", "Mars"},
		},
		{
			"Arrays",
			[2]Profile{
				{30, "Earth"},
				{33, "Mars"},
			},
			[]string{"Earth", "Mars"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{30, "Earth"}
			aChannel <- Profile{33, "Mars"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Earth", "Mars"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{30, "Earth"}, Profile{33, "Mars"}
		}

		var got []string
		want := []string{"Earth", "Mars"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
