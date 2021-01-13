<p align="center" ><img alt="Golpher Ninja by Takuya Ueda" src="https://raw.githubusercontent.com/tenntenn/gopher-stickers/master/png/ninja.png"></p>

<h1 align="center">WPrecon (Wordpress Recon)</h1>
<p align="center"> 
  <a href="https://www.gnu.org/licenses/gpl-3.0">
    <img alt="License: GPL v3" src="https://img.shields.io/badge/License-GPLv3-blue.svg">
  </a>
  <img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/m/blackcrw/wprecon">
  <img alt="GitHub go.mod Go version (branch)" src="https://img.shields.io/github/go-mod/go-version/blackcrw/wprecon/master?label=Go&logo=go">
  
Hello! Welcome. Wprecon (Wordpress Recon), is a vulnerability recognition tool in CMS Wordpress, 100% developed in Go.
</p> 

#### Version: 0.1.3.0a

### Notice:
#### Hello my friends, a new super improved version of wprecon is coming... (13/01/2021)
#### [Branch Dev](https://github.com/blackcrw/wprecon/tree/dev)

### [Compile and Install](https://github.com/blackcrw/wprecon/doc/compile&install.md)

## Features

| Status        | Features              |
|---------------|-----------------------|
|   ✅          | Random Agent          |
|   ✅          | Detection WAF         |
|   ✅          | User Enumerator       |
|   ✅          | Plugin Scanner        |
|   ✅          | Theme Scanner         |
|   ✅          | Tor Proxy's           |
|   ✅          | Detection Honeypot    |
|   ✅          | Fuzzing Backup Files  |
|   🔨          | Fuzzing Passwords     |
|   🔨          | Vulnerability Scanner |

<h2 align="center">Usage</h2>

| Flag(s)                    | Description                                                            |
|----------------------------|------------------------------------------------------------------------|
|  -u, --url string          | Target URL (Ex: http(s)://example.com/). (Required)                    |
|      --users-enumerate     | Use the supplied mode to enumerate Users.                              |
|      --themes-enumerate    | Use the supplied mode to enumerate Themes.                             |
|      --plugins-enumerate   | Use the supplied mode to enumerate Plugins.                            |
|      --detection-waf       | I will try to detect if the target is using any WAF.                   |
|      --detection-honeypot  | I will try to detect if the target is a honeypot, based on the shodan. |
|      --no-check-wp         | Will skip wordpress check on target.                                   |
|      --random-agent        | Use randomly selected HTTP(S) User-Agent header value.                 |
|      --tor                 | Use Tor anonymity network.                                             |
|      --disable-tls-checks  | Disables SSL/TLS certificate verification.                             |
|  -h, --help                | help for wprecon.                                                      |
|  -v, --verbose             | Verbosity mode.                                                        |

<h2 align="center">WPrecon running</h2>

Command: `wprecon --url "https://www.xxxxxxx.com/" --detection-waf`
##### Output:
```
—————————————————————————————————————————————————————————————————————

___       ______________________________________________   __
__ |     / /__  __ \__  __ \__  ____/_  ____/_  __ \__  | / /
__ | /| / /__  /_/ /_  /_/ /_  __/  _  /    _  / / /_   |/ /
__ |/ |/ / _  ____/_  _, _/_  /___  / /___  / /_/ /_  /|  /
____/|__/  /_/     /_/ |_| /_____/  \____/  \____/ /_/ |_/

Github: https://github.com/blackcrw/wprecon
Version: 0.0.1a
—————————————————————————————————————————————————————————————————————
[•] Target: https://www.xxxxxxx.com/
[•] Starting: 09/jan/2020 12:11:17

[•] Listing enable: https://www.xxxxxxx.com/wp-content/plugins/
[•] Listing enable: https://www.xxxxxxx.com/wp-content/themes/
[•••] Status Code: 200 — URL: https://www.xxxxxxx.com/wp-admin/
[•••] I'm not absolutely sure that this target is using wordpress! 37.50% chance. do you wish to continue ? [Y/n]: Y
[•••] Status Code: 200 — WAF: Wordfence Security Detected
[•••] Do you wish to continue ?! [Y/n] : Y
```

### Yes Baby, Thank You! ✋
