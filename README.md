[Ondrej Sika (sika.io)](https://sika.io) | <ondrej@sika.io>

# Loki Training: Log Management in Grafana

![Loki Reykjavik](_images/loki_reykjavik.png)

## What is Loki

Loki is a horizontally scalable, highly available, multi-tenant log aggregation system inspired by Prometheus. It is designed to be very cost effective and easy to operate. It does not index the contents of the logs, but rather a set of labels for each log stream.

-- https://grafana.com/oss/loki/

## How does Grafana Loki work?

![](_images/loki_diagram.svg)

## Example of log line

![](_images/loki_log_example.svg)

## Loki Components

- https://grafana.com/docs/loki/latest/get-started/components/

- **Distributor** - The distributor service is responsible for handling incoming push requests from clients.
- **Ingester** - The ingester service is responsible for persisting data and shipping it to long-term storage.
- **Querier** - The querier service is responsible for executing Log Query Language (LogQL) queries from clients.
- **Query Frontend** - The query frontend is an optional service providing the querier’s API endpoints and can be used to accelerate the read path.
- **Query Scheduler** - The query scheduler is responsible for distributing query load across queriers.
- **Index Gateway** - The index gateway service is responsible for handling and serving metadata queries.
- **Compactor** - The compactor service is used by “shipper stores”, such as single store TSDB or single store BoltDB, to compact the multiple index files produced by the ingesters and shipped to object storage into single index files per day and tenant.
- **Ruler** - The ruler service manages and evaluates rule and/or alert expressions provided in a rule configuration.

## Loki Deployment Models

- https://grafana.com/docs/loki/latest/get-started/deployment-modes/

## Loki Monolithic

![](_images/loki_monolithic.png)

## Loki Scalable

![](_images/loki_scalable.png)

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

## Alloy

## What is Alloy

Grafana Alloy is a vendor-neutral distribution of the OpenTelemetry (OTel) Collector. Alloy uniquely combines the very best OSS observability signals in the community.

Alloy offers native pipelines for OTel, Prometheus, Pyroscope, Loki, and many other metrics, logs, traces, and profile tools.

-- https://grafana.com/docs/alloy/latest/

## Install Alloy

On Mac:

```
brew install grafana/grafana/alloy
```

On Linux (using slu):

```
slu ib alloy
```

VS Code Extension:

- https://marketplace.visualstudio.com/items?itemName=Grafana.grafana-alloy

## Convert Promtail config to Alloy

- https://grafana.com/docs/alloy/latest/set-up/migrate/from-promtail/

```
alloy convert --source-format=promtail --output=<OUTPUT_CONFIG_PATH> <INPUT_CONFIG_PATH>
```

```
alloy convert --source-format=promtail --output=examples/alloy/simple.alloy examples/promtail/promtail.simple.yml
```

## Run Alloy

```
alloy -config.file examples/alloy/simple.alloy
```

check using logcli:

```
logcli query '{job="alloy"}' --follow
```

## Send logs from Go application

```
(cd examples/example_go && go run .)
```

## Grafana

## Run Grafana

```
docker run -p 3000:3000 grafana/grafana
```

## Doggos

[![@wild.loki.appears](./_images/loki1.jpg)](https://www.instagram.com/wild.loki.appears/)
[![@loki](./_images/loki2.jpg)](https://www.instagram.com/loki/)

[@wild.loki.appears](https://www.instagram.com/wild.loki.appears/)

[@loki](https://www.instagram.com/loki/)

## Thank you! & Questions?

That's it. Do you have any questions? **Let's go for a beer!**

### Ondrej Sika

- email: <ondrej@sika.io>
- web: <https://sika.io>
- twitter: [@ondrejsika](https://twitter.com/ondrejsika)
- linkedin: [/in/ondrejsika/](https://linkedin.com/in/ondrejsika/)
- Newsletter, Slack, Facebook & Linkedin Groups: <https://join.sika.io>

_Do you like the course? Write me recommendation on Twitter (with handle `@ondrejsika`) and LinkedIn (add me [/in/ondrejsika](https://www.linkedin.com/in/ondrejsika/) and I'll send you request for recommendation). **Thanks**._

Wanna to go for a beer or do some work together? Just [book me](https://book-me.sika.io) :)
