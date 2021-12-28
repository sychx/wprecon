<p align="center" ><img alt="Golpher Ninja by Takuya Ueda" src="https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/ninja.png"></p>

<h1 align="center">WPRecon (Wordpress Recon)</h1>
<p align="center">
  <a href="https://lgtm.com/projects/g/blackcrw/wprecon/alerts/"><img alt="Total alerts" src="https://img.shields.io/lgtm/alerts/g/blackcrw/wprecon.svg?logo=lgtm&logoWidth=18"/></a>
  <img alt="GitHub go.mod Go version (branch)" src="https://img.shields.io/github/go-mod/go-version/blackcrw/wprecon/master?label=Go&logo=go">
</p>

```
WPRecon, is a tool for the recognition of vulnerabilities and blackbox information for wordpress.

We can use wprecon to recognize the versions of plugins, themes, and wordpress core, in addition to counting users, and waf (web application firewall).

The purpose of this tool is just to help developers find possible loopholes in their systems/wordpress sites.
```

<p align="center">
  <a href="https://github.com/blackcrw/wprecon/wiki/Compile-and-Install">Build & Install</a> •
  <a href="https://github.com/blackcrw/wprecon/blob/master/LICENSE">License</a> • 
  <a href="https://github.com/blackcrw">Author</a>
</p>

<h4>
  <p align="center">
    <code>v2.3.2a</code>
  </p>
</h4>
<br>
 
<h3 align="center">🔨 <code>Build</code></h3>

```
For you to compile wprecon you will need to have the golang compiler installed.
And for that you will access the official website of golang and will download and install it.

- https://golang.org/dl/

Once downloaded and installed you will download wprecon directly from github with the command:
```

1. `go get github.com/blackcrw/wprecon`
2. `cd $(go env GOPATH)/src/github.com/blackcrw/wprecon`
3. `make build`
4. `make install`

<h3 align="center">⚠️ <code>Warning</code></h3>

```
wprecon does not have any kind of connection to the site: wprecon.com
```

<h3 align="center">🚀 <code>Running</code></h3>

<img align="center" alt="wprecon -u https://xxxx.com/ --agressive-mode --random-gent" src="https://i.imgur.com/zyfINsx.png">