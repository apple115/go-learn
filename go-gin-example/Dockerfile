FROM scratch

WORKDIR /app

# 假设编译后的可执行文件名为 go-gin-example
COPY go-gin-example /app/
COPY conf/app.ini /app/conf/

EXPOSE 8000

CMD ["/app/go-gin-example"]
