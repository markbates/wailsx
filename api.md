# Using the API

## The `API` type

<figure id="api-type">

<go doc="-short github.com/markbates/wailsx.API"></go>

<figcaption>The `github.com/markbates/wailsx.API` type</figcaption>

</figure>

## `Nil` API Calls

<figure id="nil-api-snip">

<code src="doc_test.go" snippet="nil-api"></code>

<figcaption>Calling methods on a `nil` `API`.</figcaption>

</figure>

<figure id="nil-api-test">

<go test="-v -run Test_Nil_API_Call"></go>

<figcaption>Running the test for calling methods on a `nil` `API`.</figcaption>

</figure>

<figure id="nil-api-test-prod">

<go test="-v -run Test_Nil_API_Call -tags wails" exit="1"></go>

<figcaption>Running the test for calling methods on a `nil` `API` in production mode.</figcaption>

</figure>
