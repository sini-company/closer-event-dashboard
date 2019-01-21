FROM grafana/grafana:5.4.3
ENV GF_PATHS_DATA=/var/lib/grafana \
    GF_PATHS_LOGS=/var/log/grafana \
    GF_PATHS_PLUGINS=/var/lib/grafana/plugins \ 
    GF_PATHS_PROVISIONING=/etc/grafana/provisioning
ADD ./config.ini /etc/grafana/config.ini
ADD ./provisioning ${GF_PATHS_PROVISIONING}
ADD ./dashboards ${GF_PATHS_DATA}/dashboards

RUN mkdir -p $GF_PATHS_PLUGINS &&\
    grafana-cli --pluginsDir $GF_PATHS_PLUGINS plugins install btplc-status-dot-panel &&\
    grafana-cli --pluginsDir $GF_PATHS_PLUGINS plugins install grafana-clock-panel &&\
    grafana-cli --pluginsDir $GF_PATHS_PLUGINS plugins install grafana-piechart-panel &&\
    grafana-cli --pluginsDir $GF_PATHS_PLUGINS plugins install jdbranham-diagram-panel &&\
    grafana-cli --pluginsDir $GF_PATHS_PLUGINS plugins install mtanda-histogram-panel &&\
    grafana-cli --pluginsDir $GF_PATHS_PLUGINS plugins install ryantxu-ajax-panel
