FROM qiujiahong/chromedp:v1.0

# 安装中文字体和chrome
# RUN yum install -y wget && \
#    yum install -y wqy-microhei-fonts wqy-zenhei-fonts && \
#    wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm && \
#    yum install -y ./google-chrome-stable_current_*.rpm && \
#    google-chrome --version && \
#    rm -rf *.rpm

# 设置go mod proxy 国内代理
# 设置golang path
#ENV GOPROXY=https://goproxy.io GOPATH=/gopath PATH="${PATH}:/usr/local/go/bin"
# 定义使用的Golang 版本 1.13.3
#ARG GO_VERSION=1.15.5

# 安装 golang 1.15.5
#RUN wget "https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz" && \
#    rm -rf /usr/local/go && \
#    tar -C /usr/local -xzf "go$GO_VERSION.linux-amd64.tar.gz" && \
#    rm -rf *.tar.gz && \
#    go version && go env;


COPY apps /apps
WORKDIR /apps

#RUN cd chromegin && go build -o app;

EXPOSE 8080

# 保存图片网页图片截图
#VOLUME /data

CMD ["/apps/page-ss"]