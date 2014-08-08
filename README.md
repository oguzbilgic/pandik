# Pandik

Monitoring tool for web services. Self-hosted [pingdom](http://pingdom.com) alternative.

## Installation 

If you have go tools installed to your system, enter the command bellow to your terminal.

```bash
$ go get github.com/oguzbilgic/pandik
```
    
## Build from a clone

go get -d 
go build
./pandik

## Configuration

Pandik uses `~/.pandik.json` file for configuration by default, but you can overwrite this by using 
`-c` command file with path to your configuration file. Here is a sample configuration file: 

```json
{

  "monitors": [
    {
      "type": "http-status",
      "url": "http://localhost:8000",
      "name": "My website healthcheck",
      "freq": "10s",
      "timeout": "2s"
    }

  ],
  
  "notifiers": [
    {
      "type": "flapjack",
      "address" : "boot2docker:6380"
    },
    { 
      "type": "stderr"
    }
  ]
}
```

## Usage

Locate your configuration file and run the comman bellow

```bash
$ pandik -c /path/to/configuration.json
```

To run pandik as a deamon on your system use the `-d` flag

```bash
$ pandik -d -c /path/to/configuration.json
```

By default pandik uses `~/.pandik.log` for deamon's log file, but this can be overwritten by `-l` flag

```bash
$ pandik -d -l /path/to/log.file -c /path/to/configuration.json
```

## Usage with Flapjack

http://flapjack.io is a alert routing and event processing system, pandik can feed events into it (flapjack expects a heartbeat of events).

To use this - use the flapjack notifier - with the "address" being the redis hostname and port. 
    
## License

The MIT License (MIT)
