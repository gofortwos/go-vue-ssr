// Code generated by go-vue-ssr: https://github.com/zbysir/go-vue-ssr
// src_hash:6124f6827f08c66a298ec09d87e99bda

package async

import (
	"strings"
)

type _ strings.Builder

func (r *Render) Component_async(options *Options) PromiseGroup {
	scope := extendScope(r.Global.Scope, options.Props)
	options.Directives.Exec(r, options)
	_ = scope

	r.Wg.Add(1)
	s := make(chan string, 1)
	go func() {
		s <- r.Component_slot(&Options{
			Slots: map[string]NamedSlotFunc{},
			P:     options,
			Scope: scope,
		}).Join()
		r.Wg.Done()
	}()

	var f PromiseFunc = func() string {
		return <-s
	}

	return PromiseGroup{f}
}