// package route
package http

import (
	"lurcury/types"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestSignbatch(t *testing.T) {
	// write some test
	r := gofight.New()

}

func TestBasicHelloWorld(t *testing.T) {
	r := gofight.New()
	core_arg := &types.CoreStruct{}

	r.GET("/hello").
		SetDebug(true).
		// Run(Server(core_arg), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		Run(Server(core_arg), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, "Hello World", r.Body.String())
			// assert.Equal(t, http.StatusOK, r.Code)
		})
}

// func TestRouter(t *testing.T) {
// type args struct {
// coreStruct *types.CoreStruct
// }
// tests := []struct {
// name string
// args args
// }{
// // TODO: Add test cases.
// }
// for _, tt := range tests {
// t.Run(tt.name, func(t *testing.T) {
// Router(tt.args.coreStruct)
// })
// }
// }
