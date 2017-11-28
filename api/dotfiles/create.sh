env GOOS=linux GOARCH=386 go build -o horizon
unset HTTP_PROXY
unset HTTPS_PROXY
docker build -t horizon .
export HTTP_PROXY=grc-americas-sanra-pitc-wkcz.proxy.corporate.gtm.ge.com:80
export HTTPS_PROXY=grc-americas-sanra-pitc-wkcz.proxy.corporate.gtm.ge.com:80
docker-machine env --shell cmd default
winpty docker run -p 8080:8080 --rm -it horizon
