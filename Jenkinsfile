pipeline {
    triggers {
        pollSCM 'H/5 * * * *'
    }
    agent any
    environment {
        docker_image_name = "fouegereregent39/docker-site"
        docker_credentials = "docker-hub-credentials"
        docker_image = ''
        docker_image_latest = ''
    }

    stages {
        stage('Clonning Repository') {
            steps {
                echo 'Clonning Repository'
                git branch: 'dev', credentialsId: 'b5fb04b9-6284-44ad-b4cb-54104a5b1453', url: 'git@github.com:FougereRegent/docker-site.git'     
            }

        }
        stage('Build Image') {
            steps {
                script {
                    docker_image = docker.build "${docker_image_name}:dev-${env.BUILD_ID}"
                    docker_image_latest = docker.build "${docker_image_name}:dev-latest"
                }
            }
        }
        stage('Push Image') {
            steps {
                script {
                    docker.withRegistry("", docker_credentials) {
                        docker_image.push()
                        docker_image_latest.push()
                    }
                }
            }
        }
    }

}
