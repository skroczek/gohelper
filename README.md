Go Helper
=========

This repository provides some handy helper function. Generics are heavily used for this purpose.

Provided Functions
==================

`boolean` Functions
-------------------

### `boolean.In`

The `In` function is a shorthand for multiple OR conditions.

Example:

```go
package main

import (
	"fmt"
	"github.com/skroczek/gohelper/boolean"
)

func main() {
    v := "1"
    // Alternative: if v == "2" || v == "3" ...
	if boolean.In(v, "2", "3" /*, ... */) {
		fmt.Printf("The first argument occurs at least once more in the remaining ones.")
	} else {
		fmt.Printf("The first argument does not occur again")
    }
}
```

### `boolean.NotIn`

The `NotIn` function is a shorthand for multiple AND NOT EQUALS conditions.

Example:

```go
package main

import (
	"fmt"
	"github.com/skroczek/gohelper/boolean"
)

func main() {
    v := "1"
    // Alternative: if v != "2" && v != "3" ...
	if boolean.NotIn(v, "2", "3" /*, ... */) {
		fmt.Printf("The first argument does not occur again")
	} else {
		fmt.Printf("The first argument occurs at least once more in the remaining ones.")
    }
}
```


`diff` Functions
----------------

### `diff.Diff`

The `Diff` function returns the elements of two slices that have been added or removed. The first slice is considered the
"old" and the second the "new".

### `diff.Added`

The `Added` function returns the elements of two slices that occur in the **second** slice but not in the **first**.

### `diff.Removed`

The `Removed` function returns the elements of two slices that occur in the **first** slice but not in the **second**.

`error` Functions
-----------------

### `error.Panic`

Triggers a panic when err != nil

### `error.Must`

Triggers a panic if the err != nil, otherwise only the first return value is returned. The second (error) return
value is not returned.

### `error.Must2`

Triggers a panic if err != nil, otherwise only the first two return values are returned. The third (error) return
value is not returned.

`interfaces`
------------

### `interfaces.Equaler`

The Equaler interface allows for the general comparison of two (same) types. It is mainly used for the `diff.Diff`
function.

Example:

```go
package main

import "fmt"

type String string

func (s String) Equal(v String) bool {
	return s == v
}

func main() {

	a := String("a")
	b := String("b")

	if a.Equal(b) {
		fmt.Printf("%s == %s", a, b)
	}else{
		fmt.Printf("%s !== %s", a, b)
    }// Output: a !== b
}
```

`ptr`
-----

### `ptr.Ptr`

The function `ptr.Prt` is a shortcut for creating a pointer on the fly. Useful mainly for native types like int,
float64, string, etc.

```go
package main

import "github.com/skroczek/gohelper/ptr"

func main() {
	_ = struct {
		a *int
		b *string
		c *float32
		d *float64
		//...
	}{
		a: ptr.Ptr(1),
		b: ptr.Ptr("a"),
		c: ptr.Ptr(float32(21)),
        d: ptr.Ptr(23.42),
        // ...
	}
}
```

Test `assert` Functions
-----------------------

### `assert.Equal`

Tests whether two values are equal. If the test fails, an error is generated with a corresponding message. Under the
hood, the function `reflect.DeepEqual` is used.

### `assert.NotEqual`

Tests whether two values are not equal. If the test fails, an error is generated with a corresponding message. Under the
hood, the function `reflect.DeepEqual` is used.

License
=======

Copyright 2023 Sebastian Kroczek <me@xbug.de>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.