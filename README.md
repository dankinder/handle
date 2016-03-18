# handle

Handle go(golang) errors in simple scripts with an easy one-liner. Package handle
provides a dead-simple function to panic on error. *Not* for use in production
software where the typical method of check `if err != nil` is the right thing to do.

Writing Go involves many error checks:

```go
func CopyFile(dstName, srcName string) {
	src, err := os.Open(srcName)
	if err != nil {
		// Handle this somehow...
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		// Handle this somehow...
	}
	defer dst.Close()

	written, err := io.Copy(dst, src)
	if err != nil {
		// Handle this somehow...
	}
}
```

For production software this is absolutely essential and despite being verbose
it is great design. But it hurts for simple scripting tasks where we would
prefer that it simply panic, as python/java/etc would.

So let's just tell it to handle it for us:

```go
	func CopyFile(dstName, srcName string) {
		src, err := os.Open(srcName)
		handle.Err(err)
		defer src.Close()

		dst, err := os.Create(dstName)
		handle.Err(err)
		defer dst.Close()

		written, err := io.Copy(dst, src)
		handle.Err(err)
	}
```

That is the purpose of this simple package.
