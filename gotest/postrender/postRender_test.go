package postrender_test

import (
	"bytes"
	"gotest/postrender"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var (
		aPost = postrender.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	// t.Run("convert a single post heading into HTML", func(t *testing.T) {
	// 	buf := bytes.Buffer{} //one of a writer type
	// 	err := postrender.Render(&buf, aPost)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	got := buf.String()
	// 	want := `<h1>hello world</h1>`
	// 	if got != want {
	// 		t.Errorf("got '%s' want '%s'", got, want)
	// 	}
	// })

	/*v2.0*/

	// t.Run("it converts a single post into HTML", func(t *testing.T) {
	// 	buf := bytes.Buffer{}
	// 	err := postrender.Render(&buf, aPost)

	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	got := buf.String()
	// 	want := `<h1>hello world</h1><p>This is a description</p>Tags: <ul><li>go</li><li>tdd</li></ul>`

	// 	if got != want {
	// 		t.Errorf("got '%s' want '%s'", got, want)
	// 	}
	// })
	/*using approval test*/

	postParser, err := postrender.NewPostParser()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := postParser.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}

// benchmark to see the effect of parsing every time we run a test versus arsing once.
func BenchmarkRender(b *testing.B) {
	var (
		aPost = postrender.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postParser, err := postrender.NewPostParser()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postParser.Render(io.Discard, aPost) //Discard is a Writer on which all Write calls succeed without doing anything.
	}
}
