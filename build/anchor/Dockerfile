#基础镜像
FROM alpine
#添加作者
LABEL maintainer "COSMOPlat-DEV <support@cosmoplat.com>"
RUN apk --no-cache add iproute2
#添加文件
ADD anchor /opt/cni/bin/anchor
ADD octopus /opt/cni/bin/octopus
ADD install-cni.sh /install-cni.sh
#设置环境变量
ENV PATH=$PATH:/opt/cni/bin
VOLUME /opt/cni
#工作目录
WORKDIR /opt/cni/bin
CMD ["/opt/cni/bin/anchor"]
