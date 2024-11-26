[Ondrej Sika (sika.io)](https://sika.io) | <ondrej@sika.io>

# Loki Training: Log Management in Grafana

![Loki Reykjavik](_images/loki_reykjavik.png)

## Install Mac

```
brew install loki
brew install promtail
brew install logcli
```

## Install Linux (using slu)

```
slu ib loki
slu ib promtail
slu ib logcli
```

## Run Loggen

```
slu loggen > /tmp/example.log
```

## Run Loki

```
loki -config.file examples/loki/loki.simple.yml
```

## Run Promtail

```
promtail -config.file examples/promtail/promtail.simple.yml
```

## Run Logcli

```
logcli query '{job="example"}' --follow
```
