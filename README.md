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

## Slice

```go
  import "github.com/kirinlabs/uitls/sli"

  sli.InSlice(v interface{}, s interface{}) bool
  sli.InIfaceSlice(v interface{}, sl []interface{}) bool
  sli.Slice(iface interface{}) ([]interface{}, error)    //Support []string,[]int,[]int64,[]float64
  sli.Unique(list *[]string) []string
  sli.UniqueInt(list *[]int) []int
  sli.UniqueInt64(list *[]int64) []int64
  sli.Chunk(slice []interface{}, size int) (chunk [][]interface{})
  sli.Range(start, end, step int64) (intslice []int64)
  sli.Pad(slice []interface{}, size int, val interface{}) []interface{}
  sli.Shuffle(slice []interface{}) []interface{}
  sli.Diff(slice1, slice2 []string) (diffslice []string)
  sli.Intersect(slice1, slice2 []string) (interSlice []string)
  sli.Rand(a []string) (b string)
  sli.RandInt(a []int) (b int)
  sli.RandInt64(a []int64) (b int64)
  sli.Sum(i []int64) (s int64)
  sli.Range(start, end, step int64) (intslice []int64)
```