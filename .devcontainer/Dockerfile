ARG VERSION_PIMO=v1.19.0

FROM cgifr/pimo:${VERSION_PIMO} AS pimo

FROM adrienaury/go-devcontainer:v2.0

USER root

RUN apk add --update --progress --no-cache make gomplate

ARG VERSION_GOLICENSE=0.2.0
ARG VERSION_MILLER=6.2.0
ARG VERSION_YQ=4.40.4
RUN    wget -nv -O- https://github.com/mitchellh/golicense/releases/download/v${VERSION_GOLICENSE}/golicense_${VERSION_GOLICENSE}_linux_x86_64.tar.gz | tar xz -C /usr/bin golicense \
    && wget -nv -O- https://github.com/johnkerl/miller/releases/download/v${VERSION_MILLER}/miller-${VERSION_MILLER}-linux-amd64.tar.gz | tar xz --strip-components 1 -C /usr/bin miller-${VERSION_MILLER}-linux-amd64/mlr \
    && wget -nv -O /usr/bin/yq https://github.com/mikefarah/yq/releases/download/v${VERSION_YQ}/yq_linux_amd64 \
    && chmod +x /usr/bin/golicense /usr/bin/mlr /usr/bin/yq

COPY --from=pimo /usr/bin/pimo /usr/bin/pimo

USER vscode
