pipeline {
    agent any

    environment {
        GOPATH = '/go'
    }

    options {
        // Set up GitHub credentials
        checkout([$class: 'GitSCM', branches: [[name: '*/main']], userRemoteConfigs: [[url: 'https://github.com/abihi/gobra.git']]])
    }

    stages {
        stage('Build') {
            steps {
                script {
                    dir('.') {
                        // Download Go modules
                        sh 'go mod download'

                        // Build the Go app
                        sh 'CGO_ENABLED=0 GOOS=linux go build -o /go/bin/gobra ./cmd/gobra'
                    }
                }
            }
        }

        stage('Unit Test') {
            steps {
                script {
                    dir('.') {
                        // Run Go unit tests
                        sh 'go test ./...'
                    }
                }
            }
        }
    }

    post {
        success {
            echo 'Build and test successful! Deploy your application.'
        }
        failure {
            echo 'Build or test failed! Take corrective actions.'
        }
    }
}
