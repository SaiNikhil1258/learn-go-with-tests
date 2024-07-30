package hello

import "testing"

// T for Testing and B for Benchmark
func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Nikhil", "")
		want := "Hello, Nikhil"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say, 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Hello in spanish", func(t *testing.T) {
		got := Hello("Nikhil", "Spanish")
		want := "Hola, Nikhil"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Hello in French", func(t *testing.T) {
		got := Hello("Nikhil", "French")
		want := "Bounjour, Nikhil"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	/*
	   if we use the Helper function then if the function fails the line number reported
	   will be in our function call rather than inside our test helper.
	*/
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*
Write a Test
Make the compiler pass
Run the test, see that it fails and check the error message in meaningful
Write enough code to make the test pass
Refactor
*/
