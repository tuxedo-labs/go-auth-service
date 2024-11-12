pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                script {
                    sh 'docker build -t auth-service:${BUILD_NUMBER} .'
                }
            }
        }
        stage('Test') {
            steps {
                script {
                    sh 'docker run --rm auth-service:${BUILD_NUMBER} go test ./...'
                }
            }
        }
        stage('Deploy') {
            steps {
                script {
                    sh 'docker tag auth-service:${BUILD_NUMBER} yourdockerhubusername/auth-service:${BUILD_NUMBER}'
                    sh 'docker push yourdockerhubusername/auth-service:${BUILD_NUMBER}'
                }
            }
        }
    }
}

