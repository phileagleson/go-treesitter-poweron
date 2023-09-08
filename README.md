# go-treesitter-poweron
---
Go bindings for [tree-sitter-poweron](https://github.com/phileagleson/tree-sitter-poweron).

## Usage 

Create a parser: 
```go
import (
    sitter "github.com/smacker/go-tree-sitter"
    "github.com/phileagleson/go-treesitter-poweron"
)

parser := sitter.NewParser()
parser.SetLanguage(poweron.GetLanguage())
```

Parse some code:
```go
sourceCode := []byte("TARGET=ACCOUNT")
tree := parser.Parse(nil, sourceCode)
```

Inspect the syntax tree:
```go
n := tree.RootNode()
fmt.Println(n) //(source_file (target_division (record_type)))
child := n.NamedChild(0)
fmt.Println(child.Type()) // target_division
fmt.Println(child.StartByte()) // 0
fmt.Println(child.EndByte()) // 13 
```

