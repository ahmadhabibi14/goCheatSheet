/*
Context merupakan sebuah data yang isinya membawa value, sinyal cancel, sinyal timeout
dan sinyal deadline. Context biasanya dibuat per request (misal setiap ada request masuk ke server
melalui http request), context digunakan untuk mempermudah kita meneruskan value, dan sinyal antar proses.
*/

package cheatsheet

import (
	"context"
	"fmt"
	"testing"
)

func doSomething(ctx context.Context) {
	fmt.Printf("Do something: name's value is %s\n", ctx.Value("name"))
}

func TestContext(t *testing.T) {
	// parent context
	ctx := context.Background()

	// child context
	ctx = context.WithValue(ctx, "name", "Habi")

	doSomething(ctx)
}
