# Pandik

Monitoring tool for web services. Self-hosted [pingdom](http://pingdom.com) alternative.

## Installation 

If you have go tools installed to your system, enter the command bellow to your terminal.

```bash
$ go get github.com/oguzbilgic/pandik
```
    
Or you can just download the compiled binary to your computer.

## Configuration

Pandik uses `~/.pandik.json` file for configuration by default, but you can overwrite this by using 
`-c` command file with path to your configuration file. Here is a sample configuration file: 

```json
{
  "api": {
    "format": "json",
    "port": 9571
  },
  "monitors": [
    {
      "type": "http-status",
      "url": "webapp.com",
      "freq": "5m"
    }
  ],
  "notifiers": [
    {
      "type": "web",
      "address": "mydomain.com/callback"
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
    
## License

The MIT License (MIT)
