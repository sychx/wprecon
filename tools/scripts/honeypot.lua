local printer = require("printer")
local http = require("http")
local url = require("url")
local net = require("net")
local client = http.client()

-- ————————————————————— PRINTER —————————————————————
-- Printer is a package created by myself to make those cute prints, with special characters, colors and etc.
-- In it you will have the functions: done, danger, warning, fatal (This will end all wprecon)
-- Example using functions: printer.done("Hello World") *Only String*
-- —————————————————————— HTTP ———————————————————————
-- This package is the golang http standard.
-- To help out here is a doc from the people who exported this package to Lua: https://github.com/vadv/gopher-lua-libs/tree/master/http
-- And as the two have theirs in some ways are similar, here's the doc of the standard golang package: https://golang.org/pkg/net/http/
-- —————————————————————— URL ————————————————————————
-- For now this package has only one function, but as I was having more needs I would implement more functions in it.
-- Obviously related to url, example: a url validator.
-- Function : url.host(target) *Only String* — This host function only serves to clear the url and return you only the host.
-- —————————————————————— NET ————————————————————————
-- Just like the url package, it has only one function and as there are more needs, I will implement more gradually.
-- Function : net.lookup_ip(host) *Only String* — The purpose of this function is to do a dns lookup and return the target's ip

-- This variable must be named "script"
script = {
    -- Here is the title of your script/vulnerability/poc/cve
    title = "Honeypot Checker",
    -- Here is the name of the author of the script/vulnerability/poc/cve
    author = "blackcrw (WPrecon)",
    -- If your script has a license
    license = "GPL-3.0",
    -- The existing levels are: Low, medium and high.
    risklevel = "Low",
    -- The type of this script. If he is a little bit of a CVE, a scanner and etc ...
    type = "Checker",
    -- It is important that your script has a well-explained description, so that everyone can understand what your script will do.
    description = "It will check if the target is a honeypot, and give you a percentage based on the shodan.",
    -- And it is also very important that there are references.
    references = {""},
}

-- It is extremely important that the main function is named main, since it will be executed when script is called.
function main(target)
    local uri_host = url.host(target)
    local ip = net.lookup_ip(uri_host)
    
    local request = http.request("GET", "https://api.shodan.io/labs/honeyscore/"..ip.."?key=C23OXE0bVMrul2YeqcL7zxb6jZ4pj2by")

    local response, err = client:do_request(request)

    if err then
        printer.danger(err)
    end

    if response.code == 200 then
        printer.done("With a "..convert(response.body).." chance of this host being a Honeypot.")
    end
end

function convert(text)
    text = text:gsub("0.0", "0%")
    text = text:gsub("0.1", "10%")
    text = text:gsub("0.2", "20%")
    text = text:gsub("0.3", "30%")
    text = text:gsub("0.4", "40%")
    text = text:gsub("0.5", "50%")
    text = text:gsub("0.6", "60%")
    text = text:gsub("0.7", "70%")
    text = text:gsub("0.8", "80%")
    text = text:gsub("0.9", "90%")
    text = text:gsub("1.0", "100%")

    return text
end