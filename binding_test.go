package poweron_test

import (
  "context"
  "testing"
  sitter "github.com/smacker/go-tree-sitter"
  "github.com/phileagleson/go-treesitter-poweron"
  "github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
  assert := assert.New(t)
  n, err := sitter.ParseCtx(context.Background(), []byte("TARGET=ACCOUNT"),poweron.GetLanguage())
  assert.NoError(err)
  assert.Equal(
    "(source_file (target_division (record_type)))",
    n.String(),
    )
}
