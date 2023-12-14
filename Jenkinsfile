pipeline {
    agent { docker { image 'golang:1.21.5-alpine3.19' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}
