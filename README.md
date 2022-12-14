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

For now, you need to have Git and Go installed on you system. To run
CleanBook, run the following commands:

```sh
# Clone the repository
git clone https://github.com/samsapti/CleanBook.git
cd CleanBook

# Run CleanBook (-port is optional, default is 8080)
go run cmd/cleanbook/main.go -path /path/to/your/data [-port 1234]
```

## Features

CleanBook currently only supports viewing messages.