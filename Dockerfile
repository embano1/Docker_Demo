FROM scratch
MAINTAINER Michael Gasch <michael_gasch@live.com>
ADD Docker_Demo /Docker_Demo
ENTRYPOINT ["/Docker_Demo"]
