// Code generated by go-vue-ssr: https://github.com/zbysir/go-vue-ssr
// src_hash:1d768bf541d8155ccea0ceb89de5d279

package async

import (
	"strings"
)

type _ strings.Builder

type Promise interface {
	Result() string
}

type PromiseString string

func (p PromiseString) Result() string {
	return string(p)
}

type PromiseFunc func() string

func (p PromiseFunc) Result() string {
	return p()
}

type PromiseGroup []Promise

func (p PromiseGroup) Join() string {
	b := strings.Builder{}
	for _, v := range p {
		b.WriteString(v.Result())
	}

	return b.String()
}

func (r *Render) Component_bench(options *Options) PromiseGroup {
	scope := extendScope(r.Global.Scope, options.Props)
	options.Directives.Exec(r, options)
	_ = scope
	var g PromiseGroup
	g = append(g, r.tag("div", true, &Options{
		PropsClass: map[string]interface{}{"a": true},
		Class:      []string{"b"},
		Slots: map[string]NamedSlotFunc{"default": func(props Props) PromiseGroup {
			var g PromiseGroup
			g = append(g, PromiseString("<span" + mixinClass(nil, []string{"d"}, map[string]interface{}{"c": true}) + mixinAttr(nil, nil, map[string]interface{}{"a": 1}) + ">\n        " + interfaceToStr(scope.Get("data", "msg"), true) + "\n    </span>" ))

			g = append(g, r.Component_async(&Options{
				Slots: map[string]NamedSlotFunc{"default": func(props Props) PromiseGroup {
					return func() PromiseGroup {
						var b PromiseGroup

						for index, item := range interface2Slice(scope.Get("data", "c")) {
							b  = append(b, func(xscope *Scope) PromiseGroup {
								scope := extendScope(xscope, map[string]interface{}{
									"$index": index,
									"item":   item,
								})

								var g PromiseGroup
								g = append(g, PromiseString("<div>"))
								g = append(g, r.Component_bench(&Options{
									Props: map[string]interface{}{"data": scope.Get("item")},
									Slots: map[string]NamedSlotFunc{},
									P:     options,
									Scope: scope,
								})...)
								g = append(g, PromiseString("</div>"))
								return g
							}(scope)...)
						}

						return b
					}()
				}},
				P:     options,
				Scope: scope,
			})...)

			return g
		}},
		P:     options,
		Scope: scope,
	})...)

	return g
}