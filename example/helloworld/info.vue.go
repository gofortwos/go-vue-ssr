// Code generated by go-vue-ssr: https://github.com/bysir-zl/go-vue-ssr
// src_hash:535087cd1e2031e7772d0d62e5390830

package main

func (r *Render) Component_info(options *Options) string {
	this := extendMap(r.Prototype, options.Props)
	_ = this
	return r.Tag("div", true, &Options{
		Style: map[string]string{"text-align": "center"},
		Slot: map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string {
			return "<p style=\"padding: 10px 0; \"" + mixinAttr(nil, nil, map[string]interface{}{"height": interfaceAdd(lookInterface(this, "height"), 1)}) + ">" + interfaceToStr(lookInterface(this, "slogan"), true) + "</p><img" + mixinAttr(nil, map[string]string{"alt": "todo logo", "height": "50px"}, map[string]interface{}{"src": lookInterface(this, "logo")}) + "></img>"
		}},
		P:    options,
		Data: this,
	})
}
