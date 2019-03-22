utils
=====
Utils is a go language development toolkit that mainly includes functions for string, slice, map, time, type judgment, type conversion, data compression, etc.


### Installing utils
    go get github.com/kirinlabs/utils

### How to use utils?


## String

```go
  import "github.com/kirinlabs/uitls/str"

  s := "hello github"
  str.Substr(s, 2, 3)
  str.Char(s)
  str.Escape(s)
  str.Before(s, "github")
  str.BeforeLast(s, "github")
  str.After(s, "git")
  str.AfterLast(s, "git")
  str.Contians(s, "hub")
  str.StartsWith(s, "hel")
  str.EndsWitdh(s, "b")
  str.String(i interface{})
  str.Length(s)
```