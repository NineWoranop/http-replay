# http-replay
This is a web application for replay static files and work as Http server

The features that are:

- read metrics from pain text file (dat file) and expose them on port 9191
- work as period (default - read new metrics every 10 seconds)

# Architecture
![Architecture diagram][architecture]

[architecture]: document/architecture.png "Architecture Diagram"

### Building from source

To build http-replay from source code, You require:
* Go [version 1.17 or greater](https://golang.org/doc/install).
* Git

Build

    $ git clone https://github.com/NineWoranop/http-replay.git
    $ go build main.go

Run

    $ http-replay -path=/metrics/

## Command Arguments

|Argument          | Default | Description|
|------------------|:-------:|:-----------|
|autorepeat        |true     |Auto repeat for replay from start|
|path              |./       |Path for read dat file|
|scrape-interval   |10s      |Scrape interval to fetch metrics and write dat file|
|total-dat-file    |1        |Number of dat files to write|
|web.listen-address|:9191     |Address to listen on for web interface and telemetry|
|web.telemetry-path|/metrics |Address to listen on for web interface and telemetry|

## Examples
#####1.Read single file from "./" folder (Expected to read 000001.dat on current folder)

    $ http-replay -autorepeat=false

#####2.Automatically read 60 data files from "/data/" folder

    $ http-replay -path=/data/ -total-dat-file=60


#####3.Read 3 data files from current folder and run on port 8080 with any ip addresses

    $ http-replay -total-dat-file=3 -web.listen-address=:8080


## License

Unlicense, see [LICENSE](https://github.com/NineWoranop/http-replay/blob/main/LICENSE).
