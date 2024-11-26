[Ondrej Sika (sika.io)](https://sika.io) | <ondrej@sika.io>

# Loki Training: Log Management in Grafana

![Loki Reykjavik](_images/loki_reykjavik.png)

## What is Loki

Loki is a horizontally scalable, highly available, multi-tenant log aggregation system inspired by Prometheus. It is designed to be very cost effective and easy to operate. It does not index the contents of the logs, but rather a set of labels for each log stream.

-- https://grafana.com/oss/loki/

## How does Grafana Loki work?

![](_images/loki_diagram.svg)

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

## logcli completion

For zsh, add this to your ~/.zshrc file:

```
eval "$(logcli --completion-script-zsh)"
```

For bash, add this to your ~/.bashrc file:

```
eval "$(logcli --completion-script-bash)"
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

## fluent-bit example

## Install fluent-bit on Mac

```
brew install fluent-bit
```

## Run fluent-bit

```
fluent-bit -c examples/fluentbit/fluentbit.conf
```

See logs in Loki using Logcli:

```
logcli query '{job="fluentbit"}' --follow
```
