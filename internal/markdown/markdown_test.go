package markdown

import (
	"testing"
)

func TestToHTML(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "bold",
			input: "**hello**",
			want:  "<strong>hello</strong>",
		},
		{
			name:  "italic",
			input: "*hello*",
			want:  "<em>hello</em>",
		},
		{
			name:  "bold italic",
			input: "***hello***",
			want:  "<strong><em>hello</em></strong>",
		},
		{
			name:  "strikethrough",
			input: "~~hello~~",
			want:  "<s>hello</s>",
		},
		{
			name:  "inline code",
			input: "use `fmt.Println`",
			want:  "use <code>fmt.Println</code>",
		},
		{
			name:  "code block",
			input: "```\nfmt.Println()\n```",
			want:  "<pre><code>fmt.Println()</code></pre>",
		},
		{
			name:  "code block with language",
			input: "```go\nfmt.Println()\n```",
			want:  "<pre><code>fmt.Println()</code></pre>",
		},
		{
			name:  "inline code preserves formatting chars",
			input: "`**not bold**`",
			want:  "<code>**not bold**</code>",
		},
		{
			name:  "markdown link",
			input: "[click here](https://example.com)",
			want:  `<a href="https://example.com">click here</a>`,
		},
		{
			name:  "bare URL",
			input: "check https://example.com for details",
			want:  `check <a href="https://example.com">https://example.com</a> for details`,
		},
		{
			name:  "blockquote",
			input: "> quoted text",
			want:  "<blockquote>quoted text</blockquote>",
		},
		{
			name:  "newlines",
			input: "line one\nline two",
			want:  "line one<br>line two",
		},
		{
			name:  "html escape",
			input: "<script>alert('xss')</script>",
			want:  "&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;",
		},
		{
			name:  "combined formatting",
			input: "**bold** and *italic* with `code`",
			want:  "<strong>bold</strong> and <em>italic</em> with <code>code</code>",
		},
		{
			name:  "plain text unchanged",
			input: "just plain text",
			want:  "just plain text",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "code block preserves formatting",
			input: "```\n**not bold** *not italic*\n```",
			want:  "<pre><code>**not bold** *not italic*</code></pre>",
		},
		{
			name:  "multi-line blockquote",
			input: "> line one\n> line two",
			want:  "<blockquote>line one</blockquote><br><blockquote>line two</blockquote>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := ToHTML(tt.input)
			if got != tt.want {
				t.Errorf("\ninput: %q\nwant:  %q\ngot:   %q", tt.input, tt.want, got)
			}
		})
	}
}
