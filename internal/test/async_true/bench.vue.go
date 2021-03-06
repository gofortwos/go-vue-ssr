// Code generated by go-vue-ssr: https://github.com/zbysir/go-vue-ssr
// src_hash:6ac29ed62023a955b75810801f2955b0

package async

import (
	"strings"
)

type _ strings.Builder

func (r *Render) Component_bench(options *Options) *PromiseGroup {
	scope := extendScope(r.Global.Scope, options.Props)
	options.Directives.Exec(r, options)
	_ = scope
	g := &PromiseGroup{}
	g.AppendGroup(r.tag("div", true, &Options{
		PropsClass: map[string]interface{}{"a": true},
		Class:      []string{"b"},
		Slots: map[string]NamedSlotFunc{"default": func(props Props) *PromiseGroup {
			g := &PromiseGroup{}
			g.Append(PromiseString("<span" + mixinClass(nil, []string{"d"}, map[string]interface{}{"c": true}) + mixinAttr(nil, nil, map[string]interface{}{"a": 1}) + ">"))
			g.Append(PromiseString(interfaceToStr(scope.Get("data", "msg"))))
			g.Append(PromiseString("</span>"))
			g.AppendGroup(func() *PromiseGroup {
				if interfaceToBool(scope.Get("a")) {
					g.Append(PromiseString("<div>"))

					g.Append(PromiseString("</div>"))
				} else {
					for index, item := range interface2Slice(scope.Get("a")) {
						g.AppendGroup(func(xscope *Scope) *PromiseGroup {
							scope := extendScope(xscope, map[string]interface{}{
								"$index": index,
								"id":     item,
							})
							g := &PromiseGroup{}
							g.Append(PromiseString("<div>"))

							g.Append(PromiseString("</div>"))

							return g
						}(scope))
					}

				}
				return nil
			}())
			g.AppendGroup(
			for index, item := range interface2Slice(scope.Get("data", "c")) {
				g.AppendGroup(func(xscope *Scope) *PromiseGroup {
					scope := extendScope(xscope, map[string]interface{}{
						"$index": index,
						"item":   item,
					})
					g := &PromiseGroup{}
					g.Append(PromiseString("<div>"))
					g.AppendGroup(r.Component_async(&Options{
						Slots: map[string]NamedSlotFunc{"default": func(props Props) *PromiseGroup {
							g := &PromiseGroup{}
							g.AppendGroup(r.Component_bench(&Options{
								Props: map[string]interface{}{"data": scope.Get("item")},
								Slots: map[string]NamedSlotFunc{"default": func(props Props) *PromiseGroup { g := &PromiseGroup{}
									return g }},
								P:     options,
								Scope: scope,
							}))

							return g
						}},
						P:     options,
						Scope: scope,
					}))

					g.Append(PromiseString("</div>"))
					return g
				}(scope))
			}
			)

			return g
		}},
		P:     options,
		Scope: scope,
	}))
	return g
}
