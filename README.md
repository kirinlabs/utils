utils
=====
Utils is a go language development toolkit that mainly includes functions such as string, slice, map, time, type judgment, type conversion, and data compression.


### Installing utils
    go get github.com/kirinlabs/utils

### How to use utils?


## Public

```go
  import "github.com/kirinlabs/uitls"

  uitls.Empty(arg interface{}) bool
  uitls.Export(arg interface{}) string
  uitls.Json(arg interface{}) string
  uitls.Decode(s string,arg ...interface{})
  
  uitls.Type(arg interface{}) string
  uitls.IsInt(arg interface{}) bool
  uitls.IsInt64(arg interface{}) bool
  uitls.IsFloat64(arg interface{}) bool
  uitls.IsString(arg interface{}) bool
  uitls.IsTime(arg interface{}) bool
  uitls.IsBool(arg interface{}) bool
  uitls.IsSlice(arg interface{}) bool

```

## String

```go
  import "github.com/kirinlabs/uitls/str"

  s := "hello github"
  str.Before(s, "github")
  str.BeforeLast(s, "github")
  str.After(s, "git")
  str.AfterLast(s, "git")
  str.StartsWith(s, "hel")
  str.EndsWitdh(s, "b")
  str.RuneIndex(s,"github")
  str.Substr(s, 2, 3)
  str.Char(s)
  str.Escape(s)
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
  sli.InInterface(v interface{}, sl []interface{}) bool
  sli.Slice(iface interface{}) ([]interface{}, error)
  sli.Unique(list *[]string) []string
  sli.UniqueInt(list *[]int) []int
  sli.UniqueInt64(list *[]int64) []int64
  sli.Chunk(slice []interface{}, size int) (chunk [][]interface{})
  sli.Range(start, end, step int64) (intslice []int64)
  sli.Pad(slice []interface{}, size int, val interface{}) []interface{}
  sli.Shuffle(slice []interface{}) []interface{}
  sli.Diff(slice1, slice2 []string) (diffslice []string)
  sli.Intersect(slice1, slice2 []string) (interSlice []string)
  sli.Reverse(list *[]string) []string
  sli.ReverseInt(list *[]int) []int
  sli.ReverseInt64(list *[]int64) []int64
  sli.ReverseFloat(list *[]float64) []float64
  sli.Rand(a []string) (b string)
  sli.RandInt(a []int) (b int)
  sli.RandInt64(a []int64) (b int64)
  sli.Sum(i []int64) (s int64)
  sli.Range(start, end, step int64) (intslice []int64)
  sli.JoinInt(s []int,sep ...string) string
  sli.SplitInt(str string, sep ...string) ([]int, error)
```

## Convert

```go
  import "github.com/kirinlabs/uitls/convert"

  convert.Int(v interface{}) (int64, error)
  convert.Float(v interface{}) (float64, error)
  convert.Bool(bs []byte) (bool, error)
  convert.Bytes(buf []byte, val reflect.Value) (b []byte, ok bool)
  convert.String(iface interface{}) string
  convert.Kind(vv reflect.Value, tp reflect.Type) (interface{}, error)

```


## Datetime

```go
  import "github.com/kirinlabs/uitls/datetime"

  datetime.Gmtime() string
  datetime.Localtime() string
  datetime.Strtotime(s string, args ...string) (time.Time, error)
  datetime.String(format ...string) string
  datetime.Year(t ...time.Time) int
  datetime.Month(t ...time.Time) int
  datetime.Day(t ...time.Time) int
  datetime.YearDay(t ...time.Time) int
  datetime.Hour(t ...time.Time) int
  datetime.Minute(t ...time.Time) int
  datetime.Second(t ...time.Time) int
  datetime.Millisecond() int64
  datetime.Microsecond() int64
  datetime.Nanosecond() int64
  datetime.Timestamp(args ...string) int64

```

## Encrypt

```go
  import "github.com/kirinlabs/uitls/encrypt"

  encrypt.Md5(s string) string
  encrypt.Sha1(s string) string
  encrypt.Sha256(s string) string
  encrypt.Sha512(s string) string
  encrypt.Base64Encode(s string) string
  encrypt.Base64Decode(s string) string
```

## sys

```go
  import "github.com/kirinlabs/uitls/sys"

  sys.RealPath(f string) string
  sys.IsExists(path string) bool
  sys.IsFile(f string) bool
  sys.IsDir(p string) bool

```