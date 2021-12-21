FROM  harbor-int.alauda.cn/e2e-automation/helloworld:latest
LABEL Version="1.1.9079079120"
ENV VERSION=1.0.15
CMD echo $VERSION && sleep 60