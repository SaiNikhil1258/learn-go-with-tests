package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}
type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Nikhil"},
			[]string{"Nikhil"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Nikhil", "London"},
			[]string{"Nikhil", "London"},
		},
		{
			"struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Nikhil", 23},
			[]string{"Nikhil"},
		},
		{
			"nested fields",
			Person{
				"Nikhil",
				Profile{23, "London"},
			},
			[]string{"Nikhil", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Nikhil",
				Profile{23, "London"},
			},
			[]string{"Nikhil", "London"},
		},
		{
			"slices",
			[]Profile{
				{23, "London"},
				{24, "Bangalore"},
			},
			[]string{"London", "Bangalore"},
		},
		{
			"arrays",
			[2]Profile{
				{23, "London"},
				{24, "Bangalore"},
			},
			[]string{"London", "Bangalore"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
		t.Run("with Maps", func(t *testing.T) {
			aMap := map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			}
			var got []string
			Walk(aMap, func(input string) {
				got = append(got, input)
			})
			assertContains(t, got, "Moo")
			assertContains(t, got, "Baa")
		})
		t.Run("with Channels", func(t *testing.T) {
			aChannel := make(chan Profile)
			go func() {
				aChannel <- Profile{23, "London"}
				aChannel <- Profile{24, "Bangalore"}
				close(aChannel)
			}()
			var got []string
			want := []string{"London", "Bangalore"}
			Walk(aChannel, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
		t.Run("with function", func(t *testing.T) {
			aFunction := func() (Profile, Profile) {
				return Profile{23, "London"}, Profile{24, "Bangalore"}
			}
			var got []string
			want := []string{"London", "Bangalore"}
			Walk(aFunction, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
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
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
