FROM golang:1.16
WORKDIR /usr/local/go/src/
COPY ./gitpr /usr/local/go/src/gitpr/
RUN cd /usr/local/go/src/ && go mod tidy && cd ./gitpr/ 
#ENTRYPOINT ["tail -f /app/GitPR/cmd/main.go"]
CMD ["bash"]
