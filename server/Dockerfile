FROM golang:1.13.5 AS development

ENV environment=dev

WORKDIR /go/src/heard

# Install entr
RUN curl -L http://eradman.com/entrproject/code/entr-4.3.tar.gz -o /tmp/entr-4.3 \
    && tar xvf /tmp/entr-4.3 -C /tmp \
    && cd /tmp/entr-4.3 \
    && ./configure; make \
    && ln -s /tmp/entr-4.3/entr /usr/local/bin

COPY ./src /go/src/heard

RUN go test -short heard/...
RUN go install -i heard/main

CMD find -name "*.go" | entr -r sh -c "go fmt heard/... \
&& go test -v heard/... \
&& go run heard/main"

# Production #############################
FROM scratch

ENV LOG_LEVEL=ERROR
ENV environment=prod

WORKDIR /usr/local/bin/heard-pos

RUN addgroup pos --system \
    && adduser --system --no-create-home --disabled-password pos pos \
    && chown -R pos:pos /usr/local/bin/heard-pos

COPY --chown=pos:pos --from=development /go/bin/main .

USER pos

CMD [ "/usr/local/bin/heard-pos/main" ]
