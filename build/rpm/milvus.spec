%if 0%{!?version:1}
%global version 2.0.0
%endif

%if 0%{!?release:1}
%global release 1%{?dist}
%endif

%if 0%{!?tag_version:1}
%global tag_version 2.0.0
%endif

%if 0%{!?git_commit:1}
%global git_commit 6336e23
%endif


Name:             milvus
Version:          %{version}
Release:          %{release}
Summary:          V2 RPM
License:          ASL 2.0
URL:              https://milvus.io/
BuildRequires:    zlib-devel git python3-devel make automake gcc gcc-c++ gcc-gfortran
BuildRequires:    cmake >= 3.18 golang >= 1.15 tbb-devel openblas-devel boost-devel
ExclusiveArch:    x86_64
Source0:          https://github.com/milvus-io/milvus/archive/refs/tags/v%{tag_version}.tar.gz#/milvus-%{tag_version}.tar.gz

%description
Milvus is an open-source vector database for unstructured data. 


%prep
mkdir -p %{_builddir}
tar -xf %{SOURCE0} -C %{_builddir}/

%build

# build install milvus
cd %{_builddir}/milvus-%{tag_version}
## remove rpath config
cmakeRpathFiles=(
    "internal/core/CMakeLists.txt"
    "internal/core/src/index/CMakeLists.txt"
)
for cmakeRpathFile in "${cmakeRpathFiles[@]}"; do
    # remove all set(CMAKE_INSTALL_RPATH.* lines
    sed -i '/^set( CMAKE_INSTALL_RPATH/d' "$cmakeRpathFile"
    # add CMAKE_SKIP_BUILD_RPATH & CMAKE_SKIP_INSTALL_RPATH
    sed -i 's/# This will set RPATH to all excutable TARGET/set( CMAKE_SKIP_BUILD_RPATH TRUE )/g' "$cmakeRpathFile"
    sed -i 's/# self-installed dynamic libraries will be correctly linked by excutable/set( CMAKE_SKIP_INSTALL_RPATH TRUE )/g' "$cmakeRpathFile"
done

goRpathFiles=(
    "internal/indexnode/index.go"
    "internal/indexnode/indexnode.go"
    "internal/querynode/cgo_helper.go"
    "internal/querynode/collection.go"
    "internal/querynode/collection_replica.go"
    "internal/querynode/load_index_info.go"
    "internal/querynode/partition.go"
    "internal/querynode/plan.go"
    "internal/querynode/query_node.go"
    "internal/querynode/reduce.go"
    "internal/querynode/segment.go"
)

for goRpathFiles in "${goRpathFiles[@]}"; do
    # remove -Wl,-rpath=${SRCDIR}/../core/output/lib configs
    sed -i 's/-Wl,-rpath=${SRCDIR}\/..\/core\/output\/lib//g' "$goRpathFiles"
done

# remove set(CMAKE_BUILD_WITH_INSTALL_RPATH TRUE) for libNGT
sed -i '/set(CMAKE_BUILD_WITH_INSTALL_RPATH TRUE)/d' internal/core/src/index/thirdparty/NGT/CMakeLists.txt

# use go mod
sed -i 's/GO111MODULE=on \$(GO) build/GO111MODULE=on \$(GO) build -mod vendor/g' Makefile

sed -i '$d' scripts/cwrapper_rocksdb_build.sh

export MILVUS_FIU_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/1.00.tar.gz'
export MILVUS_OPENTRACING_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/v1.5.1.tar.gz'
export MILVUS_PROTOBUF_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/protobuf-cpp-3.9.0.zip'
export MILVUS_YAMLCPP_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/yaml-cpp-0.6.3.tar.gz'
export MILVUS_ROCKSDB_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/v6.15.2.tar.gz'
export MILVUS_ARROW_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/apache-arrow-6.0.1.tar.gz'

export ARROW_THRIFT_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/apache_arrow_dep/thrift-0.13.0.tar.gz'
export ARROW_JEMALLOC_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/apache_arrow_dep/jemalloc-5.2.1.tar.bz2'
export ARROW_XSIMD_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/apache_arrow_dep/aeec9c872c8b475dedd7781336710f2dd2666cb2.tar.gz'
export ARROW_UTF8PROC_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/apache_arrow_dep/v2.6.1.tar.gz'
export ARROW_BOOST_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/apache_arrow_dep/boost_1_75_0.tar.gz'

export LD_LIBRARY_PATH=${PWD}/internal/core/output/lib/
make install -e BUILD_TAGS=v%{tag_version} -e GIT_COMMIT=%{git_commit}



%install

# dir
mkdir -p %{buildroot}/usr/bin/
mkdir -p %{buildroot}/lib64/milvus
mkdir -p %{buildroot}/etc/milvus/configs/advanced
mkdir -p %{buildroot}/etc/systemd/system/
mkdir -p %{buildroot}/etc/ld.so.conf.d/
mkdir -p %{buildroot}/usr/share/doc/milvus
mkdir -p %{buildroot}/%{_mandir}/man1

cd %{_builddir}/milvus-%{tag_version}

# bin
echo 'export MILVUSCONF=/etc/milvus/configs/' > %{buildroot}/usr/bin/milvus
echo 'milvus-server $@' >> %{buildroot}/usr/bin/milvus
sed -i -e '1i#!/bin/bash' %{buildroot}/usr/bin/milvus
chmod 755 %{buildroot}/usr/bin/milvus
strip bin/milvus
install -m 755 bin/milvus %{buildroot}/usr/bin/milvus-server
install -m 755 build/rpm/milvus-dependencies %{buildroot}/usr/bin/milvus-dependencies

# lib
strip lib/libmilvus_indexbuilder.so
install -m 755 lib/libmilvus_indexbuilder.so %{buildroot}/lib64/milvus/libmilvus_indexbuilder.so
strip lib/libmilvus_segcore.so
install -m 755 lib/libmilvus_segcore.so %{buildroot}/lib64/milvus/libmilvus_segcore.so
install -m 755 /usr/lib64/libopenblas.so %{buildroot}/lib64/milvus/libopenblas.so.0
install -m 755 /usr/lib64/libtbb.so %{buildroot}/lib64/milvus/libtbb.so
strip lib/libfiu.so.1.00
install -m 755 lib/libfiu.so.1.00 %{buildroot}/lib64/milvus/libfiu.so.0
strip lib/libngt.so.1.12.0 
install -m 755 lib/libngt.so.1.12.0 %{buildroot}/lib64/milvus/libngt.so.1
install -m 755 /usr/lib64/libgfortran.so.5.0.0 %{buildroot}/lib64/milvus/libgfortran.so.5

# conf
install -m 644 configs/milvus.yaml %{buildroot}/etc/milvus/configs/milvus.yaml
install -m 644 configs/advanced/etcd.yaml %{buildroot}/etc/milvus/configs/advanced/etcd.yaml

# service
install -m 644 build/rpm/services/milvus-dependencies.service %{buildroot}/etc/systemd/system/milvus-dependencies.service
install -m 644 build/rpm/services/milvus-etcd.service %{buildroot}/etc/systemd/system/milvus-etcd.service
install -m 644 build/rpm/services/milvus-minio.service %{buildroot}/etc/systemd/system/milvus-minio.service
install -m 644 build/rpm/services/milvus.service %{buildroot}/etc/systemd/system/milvus.service

# ldconf
echo '/usr/lib64/milvus' >> %{buildroot}/etc/ld.so.conf.d/milvus.conf
chmod 644 %{buildroot}/etc/ld.so.conf.d/milvus.conf

#doc
install -m 644 README.md %{buildroot}/usr/share/doc/milvus/README.md
install -m 644 build/rpm/man/milvus.1.gz %{buildroot}%{_mandir}/man1/milvus.1.gz
install -m 644 build/rpm/man/milvus-server.1.gz %{buildroot}%{_mandir}/man1/milvus-server.1.gz
install -m 644 build/rpm/man/milvus-dependencies.1.gz %{buildroot}%{_mandir}/man1/milvus-dependencies.1.gz

%post
# update ld, systemd cache
ldconfig
systemctl daemon-reload

%preun
# disable service before remove
systemctl stop milvus
systemctl disable milvus

%postun
# update ld, systemd cache
ldconfig
systemctl daemon-reload

%files
/usr/bin/milvus
/usr/bin/milvus-server
/usr/bin/milvus-dependencies

/lib64/milvus/libmilvus_indexbuilder.so
/lib64/milvus/libmilvus_segcore.so
/lib64/milvus/libopenblas.so.0
/lib64/milvus/libfiu.so.0
/lib64/milvus/libngt.so.1
/lib64/milvus/libgfortran.so.5
/lib64/milvus/libtbb.so

%config(noreplace) /etc/milvus/configs/milvus.yaml
%config(noreplace) /etc/milvus/configs/advanced/etcd.yaml

%config(noreplace) /etc/systemd/system/milvus.service
%config(noreplace) /etc/systemd/system/milvus-minio.service
%config(noreplace) /etc/systemd/system/milvus-etcd.service
%config(noreplace) /etc/systemd/system/milvus-dependencies.service

%config(noreplace) /etc/ld.so.conf.d/milvus.conf

%doc /usr/share/doc/milvus/README.md

%{_mandir}/man1/milvus.1.gz
%{_mandir}/man1/milvus-server.1.gz
%{_mandir}/man1/milvus-dependencies.1.gz

%changelog
* Sun Feb 13 2022 Yunmei Li <yunmei.li@zilliz.com> - 2.0.0-1 
- Initial version of the package
