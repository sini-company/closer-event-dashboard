FROM grafana/grafana:5.0.0
ENV GF_PATHS_PLUGINS=/opt/grafana-plugins
RUN mkdir -p $GF_PATHS_PLUGINS
RUN grafana-cli --pluginsDir $GF_PATHS_PLUGINS plugins install grafana-clock-panel &&\
    grafana-cli --pluginsDir $GF_PATHS_PLUGINS plugins install btplc-status-dot-panel &&\
    grafana-cli --pluginsDir $GF_PATHS_PLUGINS plugins install mtanda-histogram-panel &&\
    grafana-cli --pluginsDir $GF_PATHS_PLUGINS plugins install grafana-piechart-panel &&\
    grafana-cli --pluginsDir $GF_PATHS_PLUGINS plugins install grafana-worldmap-panel &&\
    grafana-cli --pluginsDir $GF_PATHS_PLUGINS plugins install jdbranham-diagram-panel