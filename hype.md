# Wailsx - A Testable, Idiomatic Wrapper for Wails

I love [Wails](https://wails.io) and have been using it to create some great applications. Unfortunately, the Wails `v2` API is not very testable or idiomatic. Wailsx is a wrapper around the Wails API that makes it easier to test and use in a more idiomatic way.

<toc></toc>

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

## Protecting Wails API Calls

Wailsx uses Go build tags to protect the Wails API calls from being called in a production environment. The <godoc>github.com/markbates/wailsx/wailsrun#API</godoc> interface, <ref>api-int</ref>, is implemented in two different files, <ref>prod-calls</ref> and <ref>wailsx-calls</ref>.

The `wailsrun/wailscalls_prod.go` file, <ref>prod-calls</ref>, is only built when any of the following builds are provided: `wails || dev || desktop || production`. This file contains the actual Wails API calls and most returned errors are `nil`.

<figure id="prod-calls" type="listing">

<code src="wailsrun/wailscalls_prod.go" snippet="BrowserOpenURL"></code>

<figcaption>Production Wails API calls: `wailsrun/wailscalls_prod.go`</figcaption>

</figure>

In all other environments, such as testing, the `wailsrun/wailscalls.go` file, <ref>wailsx-calls</ref>, is built in all environments and contains the Wailsx API calls. The Wailsx API calls are then used to call the Wails API calls in the `development` environment.

<figure id="wailsx-calls" type="listing">

<code src="wailsrun/wailscalls.go" snippet="BrowserOpenURL"></code>

<figcaption>Stubbed Wails API calls: `wailsrun/wailscalls.go`</figcaption>

</figure>

In these environments all of the Wails API calls will return the `ErrNotAvailable` error, <ref>err-not-implemented</ref>.

<figure id="err-not-implemented" type="listing">

<go doc="github.com/markbates/wailsx/wailsrun.ErrNotAvailable"></go>

<figcaption>The `ErrNotAvailable` error</figcaption>

</figure>

### Testing Invalid Wails API Calls

With the help of Go build tags, any direct calls made to the Wails API, _outside_ of a running Wails application, will return the <godoc>github.com/markbates/wailsx/wailsrun.ErrNotAvailable</godoc> error. This allows for testing of the Wails API calls in a non-Wails environment.

In the test seen in <ref>test-api</ref> we are making a direct call to the Wails API and checking the error returned. The test passes when the error returned is `ErrNotAvailable`.

<figure id="test-api" type="listing">

<code src="docs/examples/api_calls/api_calls_test.go"></code>

<figcaption>Testing the `BrowserOpenURL` method</figcaption>

</figure>

When running the tests outside of a Wails application, the `BrowserOpenURL` method will return the `ErrNotAvailable` error, <ref>err-not-implemented</ref>.

<figure id="err-not-implemented" type="listing">

<go src="docs/examples/api_calls" test="-v"></go>

<figcaption>Testing the `BrowserOpenURL` method output.</figcaption>

</figure>

If the tests are run in a Wails application, using one of the known build tags, the `BrowserOpenURL` method will call the actual Wails API method, <ref>prod-calls</ref>. The result is a call to <godoc>log.Fatal</godoc> because we don't have a valid Wails context.

<figure id="log-fatal" type="listing">

<go src="docs/examples/api_calls" test="-v -tags wails" exit="1"></go>

<figcaption>Testing the `BrowserOpenURL` method output in `production`.</figcaption>

</figure>

<include src="eventx/eventx.md"></include>
