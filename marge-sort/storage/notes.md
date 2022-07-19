Breaking things is educational:

1. If I return like `return conquer(devide(items[:len(items)/2]), devide(items[len(items)/2:]))`
things go work. I dont' have to init two variables Hudai :) 

2. If I comment the following return, It actually stack overflow. Because there is no bound to stop.

```	if len(items) < 2 {
		fmt.Println("Inside IF?", len(items) < 2)
		// return items
	}
```
Like : 
`runtime: goroutine stack exceeds 1000000000-byte limit
runtime: sp=0xc0200e0388 stack=[0xc0200e0000, 0xc0400e0000]
fatal error: stack overflow

runtime stack:
runtime.throw({0x496444?, 0x50e840?})
        /usr/local/go/src/runtime/panic.go:992 +0x71
runtime.newstack()`