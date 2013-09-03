# Pandik

Monitoring tool for web services. Self-hosted [pingdom](http://pingdom.com) alternative.

## 1. Installation 

If you have go tools installes to your system, enter the command bellow to your terminal.

    go get github.com/oguzbilgic/pandik
    
Or you can just download the compiled binary to your computer.

## 2. Configuration

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

## 3. Usage

Locate your configuration file and run the comman bellow

    pandik -c /path/to/configuration.json
    
