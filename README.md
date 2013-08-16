# cachemap

## A memory-friendly map with a user-defined persistence store.


## Usage

    go get github.com/steakknife/cachemap


```go
import "github.com/steakknife/cachemap"



type Store something

// define all methods of cachemap.Persistent



s := Store{}

m := NewCacheMap(&s)

err := m.Add("key", "value", 3, 5)
v, err := m.Get("key") // v = "value"
err = m.Delete("key")

err = m.Add("big", somethingbig, 3, 1024*1024*1024)
err = m.Evict("big") // force something big out of RAM

m.DropCaches() // drop all cached items


s2 := Store{}
m2 := NewLazyCacheMap(&s2) // faster adds, but a goroutine in the background Evict()s things

```
