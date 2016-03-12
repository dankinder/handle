/*
Package handle provides a dead-simple function to panic on error.

Writing Go involves many error checks:

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

For production software this is absolutely essential and despite being verbose
it is great design. But it hurts for simple scripting tasks where we would
prefer that it simply panic, as python/java/etc would.

So let's just tell it to handle it for us:

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

That is the purpose of this simple package.
*/
package handle

// Err is the key function of the handle package. If err is nil, it does
// nothing. If it is not nil, it panics.
func Err(err error) {
	if err != nil {
		panic(err)
	}
}
