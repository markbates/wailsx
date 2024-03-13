# Using the API

While the <godoc>github.com/markbates/wailsx/wailsrun</godoc> package can be used directly, it is recommended to use the <godoc>github.com/markbates/wailsx.API</godoc> type to create a testable, idiomatic wrapper around the Wails API.

## The `API` type

The <godoc>github.com/markbates/wailsx#API</godoc> type is a wrapper around the <godoc>github.com/markbates/wailsx/wailsrun</godoc> package. By default, if the <godoc>github.com/markbates/wailsx#API</godoc> type is `nil`, or `zero` (i.e. `&API{}`), all methods will be will be mapped directly to the `wailsrun` package. This allows you to use the <godoc>github.com/markbates/wailsx#API</godoc> type in your application without having to worry about the <godoc>github.com/markbates/wailsx#API</godoc> being `nil`.

<figure id="api-type">

<go doc="-short github.com/markbates/wailsx.API"></go>

<figcaption>The <godoc>github.com/markbates/wailsx#API</godoc> type</figcaption>

</figure>

The <godoc>github.com/markbates/wailsx#NewAPI</godoc> function can be used to create a new <godoc>github.com/markbates/wailsx#API</godoc> type. This function will populate the <godoc>github.com/markbates/wailsx#API</godoc> type with implementations of its embedded interfaces. For example, using <godoc>github.com/markbates/wailsx/eventx#NewManager</godoc> to create a new <godoc>github.com/markbates/wailsx/eventx#Manager</godoc> that will fill the needed <godoc>github.com/markbates/wailsx/eventx#EventManager</godoc> in the <godoc>github.com/markbates/wailsx#API</godoc> type.

<figure id="new-api">

<go doc="-short github.com/markbates/wailsx.NewAPI"></go>

<figcaption>The <godoc>github.com/markbates/wailsx#NewAPI</godoc> function</figcaption>

</figure>

## `Nil` API Calls

<figure id="nil-api-snip">

<code src="doc_test.go" snippet="nil-api"></code>

<figcaption>Calling methods on a `nil` <godoc>github.com/markbates/wailsx#API</godoc>.</figcaption>

</figure>

<figure id="nil-api-test">

<go test="-v -run Test_Nil_API_Call"></go>

<figcaption>Running the test for calling methods on a `nil` <godoc>github.com/markbates/wailsx#API</godoc>.</figcaption>

</figure>

<figure id="nil-api-test-prod">

<go test="-v -run Test_Nil_API_Call -tags wails" exit="1"></go>

<figcaption>Running the test for calling methods on a `nil` <godoc>github.com/markbates/wailsx#API</godoc> in production mode.</figcaption>

</figure>

## `Nop` API Calls

<figure id="nop-api">

<go doc="-short github.com/markbates/wailsx.NopAPI"></go>

<figcaption>The <godoc>github.com/markbates/wailsx#NopAPI</godoc> function</figcaption>

</figure>

<figure id="nop-api-snip">

<code src="doc_test.go" snippet="nop-api"></code>

<figcaption>Calling methods on a `nop` <godoc>github.com/markbates/wailsx#API</godoc>.</figcaption>

</figure>
