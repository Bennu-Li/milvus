# Copyright (C) 2019-2022 Zilliz. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
# with the License. You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software distributed under the License
# is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
# or implied. See the License for the specific language governing permissions and limitations under the License.

FROM rockylinux/rockylinux:8.8

ARG TARGETARCH


RUN dnf update -y && dnf -y install gcc-toolset-9-gcc gcc-toolset-9-gcc-c++ 

RUN dnf --enablerepo=powertools  install -y wget gcc gcc-c++ gdb libatomic libstdc++-static git make zip unzip tar which \
    autoconf automake golang python3 python3-pip perl-interpreter texinfo \
    pkg-config libuuid-devel libaio perl-IPC-Cmd && \
    rm -rf /var/cache/yum/*


RUN pip3 install conan==1.61.0

RUN echo "target arch $TARGETARCH"
RUN wget -qO- "https://cmake.org/files/v3.24/cmake-3.24.4-linux-`uname -m`.tar.gz" | tar --strip-components=1 -xz -C /usr/local

# RUN dnf config-manager --set-enabled powertools && dnf install libstdc++-static -y
#RUN dnf --enablerepo=powertools install libstdc++-static

RUN mkdir /tmp/ccache && cd /tmp/ccache &&\
    wget https://dl.fedoraproject.org/pub/epel/8/Everything/`uname -m`/Packages/h/hiredis-0.13.3-13.el8.`uname -m`.rpm &&\
    wget https://dl.fedoraproject.org/pub/epel/8/Everything/`uname -m`/Packages/c/ccache-3.7.7-1.el8.`uname -m`.rpm &&\
    rpm -i hiredis-0.13.3-13.el8.`uname -m`.rpm ccache-3.7.7-1.el8.`uname -m`.rpm &&\
    rm -rf /tmp/ccache


# refer: https://code.visualstudio.com/docs/remote/containers-advanced#_avoiding-extension-reinstalls-on-container-rebuild
RUN mkdir -p /home/milvus/.vscode-server/extensions \
        /home/milvus/.vscode-server-insiders/extensions \
    && chmod -R 777 /home/milvus

COPY --chown=0:0 build/docker/builder/entrypoint.sh /

ENTRYPOINT [ "/entrypoint.sh" ]
CMD ["tail", "-f", "/dev/null"]
