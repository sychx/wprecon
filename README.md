<p align="center" ><img alt="Golpher Ninja by Takuya Ueda" src="https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/ninja.png"></p>

<h1 align="center">WPrecon (Wordpress Recon)</h1>
<p align="center">
  <img alt="Build" src="https://github.com/blackbinn/wprecon/workflows/build/badge.svg">
  <img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/m/blackbinn/wprecon">
  <img alt="GitHub go.mod Go version (branch)" src="https://img.shields.io/github/go-mod/go-version/blackbinn/wprecon/master?label=Go&logo=go">

  <p align="center">
    Hello! Welcome. Wprecon (Wordpress Recon), is a vulnerability recognition tool in CMS Wordpress, 100% developed in Go.
  </p>
</p> 

<p align="center">
  <a href="#features">Features</a> •
  <a href="https://github.com/blackbinn/wprecon/wiki/Compile-and-Install">Compile & Install</a> •
  <a href="https://github.com/blackbinn/wprecon/blob/master/LICENSE">License</a> • 
  <a href="https://github.com/blackbinn">Author</a>
</p>

<h3><p align="center">Version: 1.6.5 alpha</p></h3>
<h2 align="center">⚠️ Warning</h2>
<p align="center">
I recently found out that my tool has the same name as a website ... which has the same "subject" as my tool.
But I already warn you that this wprecon that you see on github has nothing to do with this site.
I don't have a website yet!
</p>
 
<h2 align="center">✨ Features</h2>

All that are already marked, are already on wprecon.
But since they are not, it is for the reason that they are still in development.

- [x] Detection WAF (Passive & Aggressive)
- [x] Random User Agent
- [x] Tor Proxy
- [x] Enumerator (Users, Plugins & Themes)
- [x] Enumerator Version (Plugins, Themes & WordPress)
- [x] Brute Force (xml-rpc & wp-login)
- [x] Scripts
- [x] Vulnerability Version Checking (Plugins) **(Beta)**

<h2 align="center">🔨 Build</h2>

For you to compile **wprecon** you will need to have the golang compiler installed.
And for that you will access the official website of golang and will download and install it. [**Here!**](https://golang.org/dl/)

Once downloaded and installed you will download **wprecon** directly from github with the command:

1. `go get github.com/blackbinn/wprecon`;
2. `cd $(go env GOPATH)/src/github.com/blackbinn/wprecon`;
3. `make build`.

<p align="center" >
  <h2 align="center">🚀 WPrecon running</h2>
  
  <code>$ wprecon --help</code>
  <img alt="wprecon --help" src="https://i.imgur.com/KpuDUy5.png">
  
  <code>$ wprecon fuzz --help</code>
  <img alt="wprecon fuzz --help" src="https://i.imgur.com/UCo3Odu.png">
 
  <code>$ wprecon -u https://xxxx.com/ --agressive-mode --random-agent</code>
  <img alt="wprecon -u https://xxxx.com/ --agressive-mode --random-gent" src="https://i.imgur.com/7pJv2uY.png">
</p>