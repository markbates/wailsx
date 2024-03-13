# Wailsx - A Testable, Idiomatic Wrapper for Wails

I love [Wails](https://wails.io) and have been using it to create some great applications. Unfortunately, the Wails `v2` API is not very testable or idiomatic. Wailsx is a wrapper around the Wails API that makes it easier to test and use in a more idiomatic way.

<toc></toc>

## Installation

Wailsx is a Go module and can be installed with `go get`.

<figure id="wailsx-installation">

```bash
go get github.com/markbates/wailsx
```

<figcaption>Installing `github.com/markbates/wailsx` with `go get`.</figcaption>

</figure>

Once imported, you can use the `wailsx` package in your application.

<figure id="new-api">

<go doc="-short github.com/markbates/wailsx.NewAPI"></go>

<figcaption>The <godoc>github.com/markbates/wailsx#NewAPI</godoc> function</figcaption>

</figure>

<include src="wailsrun/wailsrun.md"></include>
<include src="clipx/clipx.md"></include>
<include src="dialogx/dialogx.md"></include>
<include src="eventx/eventx.md"></include>
<include src="logx/logx.md"></include>
<include src="menux/menux.md"></include>
<include src="statedata/statedata.md"></include>
<include src="windowx/windowx.md"></include>
<include src="api.md"></include>
