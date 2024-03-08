# Wailsx - A Testable, Idiomatic Wrapper for Wails

I love [Wails](https://wails.io) and have been using it to create some great applications. Unfortunately, the Wails `v2` API is not very testable or idiomatic. Wailsx is a wrapper around the Wails API that makes it easier to test and use in a more idiomatic way.

## The API Interface

The <godoc>github.com/markbates/wailsx#API</godoc> interface, <ref>api-int</ref> declares an idiomatic interface for the <godoc>github.com/wailsapp/wails/v2/pkg/runtime</godoc> package.

<figure id="api-int" type="listing">
<code src="wailsrun/api.go"></code>
<figcaption>The <godoc>github.com/markbates/wailsx#API</godoc> interface</figcaption>
</figure>

### Error Handling

In particular, the <godoc>github.com/markbates/wailsx#API</godoc> improves error handling by adding error returns to the methods that previously returned nothing. For example, the <godoc>github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu</godoc>, <ref id="runtime-menuset"></ref>, method now returns an error, <ref id="wailsx-menuset"></ref>.

<figure id="runtime-menuset" type="listing">

<go doc="github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu"></go>

<figcaption>The <godoc>github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu</godoc> method</figcaption>

</figure>

<figure id="wailsx-menuset" type="listing">

<go doc="github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu"></go>

<figcaption>The <godoc>github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu</godoc> method</figcaption>

</figure>
