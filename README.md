# Tiny Proc

> Super simple process execution for Go

## Usage

```go
import tp "github.com/rajatsharma/tiny-proc"

tp.Proc([]string{"go", "mod", "init", "github.com/rajatsharma/" + project}, &cwd)
```
