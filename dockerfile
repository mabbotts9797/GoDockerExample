FROM scratch

COPY . /app/

CMD ["/app/main"]

EXPOSE 8080