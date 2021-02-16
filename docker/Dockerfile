FROM qiujiahong/chromedp:v1.0


COPY apps /apps
WORKDIR /apps

RUN chmod +x /apps/runDocker.sh

EXPOSE 8080


CMD ["/apps/runDocker.sh"]