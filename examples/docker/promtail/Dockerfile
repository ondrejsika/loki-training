FROM  grafana/promtail:3.3.0
COPY promtail.yml /etc/promtail/promtail.yml
CMD [ "-config.file=/etc/promtail/promtail.yml", "-config.expand-env" ]
