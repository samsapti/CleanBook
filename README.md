# CleanBook

**:warning: Still in development!**

CleanBook is an alternative way to view your downloaded Facebook data,
without trackers and network requests. It requires you to have the JSON
formatted data download, such that it can easily parse the data and show
it to you.

### Why?

Facebook provides two different data download formats: HTML
(website-like) and JSON (machine-readable). Most people would want to
download the HTML format and open it in a browser. The problem? Well
guess what, it contains trackers! The HTML format makes a lot of network
requests to Meta-owned URLs, which really isn't what a data download
should do. But fear not, CleanBook fixes that problem.

## Usage

For now, you need to have Go installed on you system. To install and run
CleanBook, run the following commands:

```sh
# Install CleanBook
go install github.com/samsapti/CleanBook/cmd/cleanbook@latest

# Run CleanBook (-port is optional, default is 8080)
cleanbook -path /path/to/your/data [-port 1234]

# Get a help menu
cleanbook [-h|-help]
```

## Features

CleanBook currently only supports viewing messages.
