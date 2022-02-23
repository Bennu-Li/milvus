mkdir download_thirdparty
wget https://github.com/milvus-io/knowhere/archive/refs/tags/v1.0.1.tar.gz -P download_thirdparty/
wget https://github.com/opentracing/opentracing-cpp/archive/v1.5.1.tar.gz -P download_thirdparty/
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.9.0/protobuf-cpp-3.9.0.zip -P download_thirdparty/
wget https://github.com/jbeder/yaml-cpp/archive/yaml-cpp-0.6.3.tar.gz -P download_thirdparty/
wget https://github.com/facebook/rocksdb/archive/v6.15.2.tar.gz -P download_thirdparty/
wget https://github.com/apache/arrow/archive/apache-arrow-6.0.1.tar.gz -P download_thirdparty/

mkdir download_thirdparty/apache_arrow_dep
wget https://github.com/JuliaStrings/utf8proc/archive/v2.6.1.tar.gz -P download_thirdparty/apache_arrow_dep/
wget https://github.com/xtensor-stack/xsimd/archive/aeec9c872c8b475dedd7781336710f2dd2666cb2.tar.gz -P download_thirdparty/apache_arrow_dep/
wget https://github.com/jemalloc/jemalloc/releases/download/5.2.1/jemalloc-5.2.1.tar.bz2 -P download_thirdparty/apache_arrow_dep/
wget https://github.com/ursa-labs/thirdparty/releases/download/latest/boost_1_75_0.tar.gz -P download_thirdparty/apache_arrow_dep/
wget --trust-server-names "https://www.apache.org/dyn/closer.cgi?action=download&filename=/thrift/0.13.0/thrift-0.13.0.tar.gz" -P download_thirdparty/apache_arrow_dep/

# go mod
go mod vendor


export MILVUS_KNOWHERE_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/v1.0.1.tar.gz'
export MILVUS_OPENTRACING_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/v1.5.1.tar.gz'
export MILVUS_PROTOBUF_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/protobuf-cpp-3.9.0.zip'
export MILVUS_YAMLCPP_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/yaml-cpp-0.6.3.tar.gz'
export MILVUS_ROCKSDB_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/v6.15.2.tar.gz'
export MILVUS_ARROW_URL='%{_builddir}/milvus-%{tag_version}/download_thirdparty/apache-arrow-6.0.1.tar.gz'
