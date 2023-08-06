
FROM golang:1.19 as build 
COPY ./cmd /app/cmd
COPY ./pkg /app/pkg
COPY ./go.mod /app/go.mod
WORKDIR /app
RUN GOOS=js GOARCH=wasm go build -o out/client.wasm cmd/client/main.go 

FROM nginx:1.25.1
COPY ./assets/nginx.conf /etc/nginx/nginx.conf
COPY ./web/static/index.html /usr/share/nginx/html/index.html
COPY ./web/static/wasm_exec.js /usr/share/nginx/html/wasm_exec.js
COPY --from=build /app/out/client.wasm /usr/share/nginx/html
