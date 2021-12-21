FROM harbor-int.alauda.cn/e2e-automation/helloworld:latest
LABEL Version="1.1.9079079990"
ENV VERSION=1.0.15
CMD echo $VERSION && sleep 60
COPY a.sh /
RUN chmod +x /a.sh
CMD /a.sh
