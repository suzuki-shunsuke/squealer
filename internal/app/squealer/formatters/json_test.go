package formatters

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/owenrumney/squealer/internal/app/squealer/match"
)

func TestJsonFormatterOutput(t *testing.T) {
	trans := []match.Transgression{createTestTransgression("Joe Bloggs", "joe@bloggs.com", "2001-01-01", "abcd123456efg")}

	plainText, _ := JsonFormatter{}.PrintTransgressions(trans, false)
	assert.Equal(t, `{
  "transgressions": [
    {
      "content": "password=Password1234",
      "filename": "/config.yml",
      "secret_hash": "sdjn34rf32fds",
      "match_string": "Password1234",
      "committer": {
        "name": "Joe Bloggs",
        "email": "joe@bloggs.com"
      },
      "commit_hash": "abcd123456efg",
      "committed": "2001-01-01",
      "exclude_rule": ""
    }
  ]
}`, plainText)

	redacted, _ := JsonFormatter{}.PrintTransgressions(trans, true)
	assert.Equal(t, `{
  "transgressions": [
    {
      "content": "password=REDACTED",
      "filename": "/config.yml",
      "secret_hash": "sdjn34rf32fds",
      "match_string": "Password1234",
      "committer": {
        "name": "Joe Bloggs",
        "email": "joe@bloggs.com"
      },
      "commit_hash": "abcd123456efg",
      "committed": "2001-01-01",
      "exclude_rule": ""
    }
  ]
}`, redacted)
}
