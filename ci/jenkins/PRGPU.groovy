#!/usr/bin/env groovy

int total_timeout_minutes = 60 * 5
int e2e_timeout_seconds = 70 * 60
def imageTag=''
int case_timeout_seconds = 10 * 60
def chart_version='4.0.6'
pipeline {
    options {
        timestamps()
        timeout(time: total_timeout_minutes, unit: 'MINUTES')
        // buildDiscarder logRotator(artifactDaysToKeepStr: '30')
        // parallelsAlwaysFailFast()
        // preserveStashes(buildCount: 5)
        // disableConcurrentBuilds(abortPrevious: true)

    }
    agent {
            kubernetes {
                inheritFrom 'default'
                defaultContainer 'main'
                yamlFile 'ci/jenkins/pod/rte-gpu.yaml'
                customWorkspace '/home/jenkins/agent/workspace'
            }
    }
    environment {
        PROJECT_NAME = 'milvus'
        SEMVER = "${BRANCH_NAME.contains('/') ? BRANCH_NAME.substring(BRANCH_NAME.lastIndexOf('/') + 1) : BRANCH_NAME}"
        DOCKER_BUILDKIT = 1
        ARTIFACTS = "${env.WORKSPACE}/_artifacts"
        CI_DOCKER_CREDENTIAL_ID = "harbor-milvus-io-registry"
        MILVUS_HELM_NAMESPACE = "milvus-ci"
        DISABLE_KIND = true
        HUB = 'harbor.milvus.io/milvus'
        JENKINS_BUILD_ID = "${env.BUILD_ID}"
        CI_MODE="pr"
        SHOW_MILVUS_CONFIGMAP= true
    }

    stages {
        stage ('Build'){
            steps {
                container('main') {
                    // sh '''
                    // ./build/builder_gpu.sh /bin/bash -c "make milvus-gpu"
                    // '''
                    dir ('build'){
                            sh './set_docker_mirror.sh'
                    }
                    dir ('tests/scripts') {
                        script {
                            sh 'printenv'
                            def date = sh(returnStdout: true, script: 'date +%Y%m%d').trim()
                            sh 'git config --global --add safe.directory /home/jenkins/agent/workspace'
                            def gitShortCommit = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()    
                            imageTag="gpu-${env.BRANCH_NAME}-${date}-${gitShortCommit}"
                            sh 
                            withCredentials([usernamePassword(credentialsId: "${env.CI_DOCKER_CREDENTIAL_ID}", usernameVariable: 'CI_REGISTRY_USERNAME', passwordVariable: 'CI_REGISTRY_PASSWORD')]){
                                sh """
                                TAG="${imageTag}" \
                                ./e2e-k8s.sh \
                                --skip-export-logs \
                                --skip-install \
                                --skip-cleanup \
                                --skip-setup \
                                --gpu \
                                --skip-test
                                """

                                // stash imageTag info for rebuild install & E2E Test only
                                sh "echo ${imageTag} > imageTag.txt"
                                stash includes: 'imageTag.txt', name: 'imageTag'

                            }
                        }
                    }
                }
            }
        }
    }
    // post{
    //     unsuccessful {
    //             container('jnlp') {
    //                 dir ('tests/scripts') {
    //                     script {
    //                         def authorEmail = sh(returnStdout: true, script: './get_author_email.sh ')
    //                         emailext subject: '$DEFAULT_SUBJECT',
    //                         body: '$DEFAULT_CONTENT',
    //                         recipientProviders: [developers(), culprits()],
    //                         replyTo: '$DEFAULT_REPLYTO',
    //                         to: "${authorEmail},devops@zilliz.com"
    //                     }
    //                 }
    //             }
    //         }
    //     }
}
