pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                script {
                    // Set up GitHub credentials and checkout the code
                    checkout([$class: 'GitSCM', branches: [[name: '*/main']], userRemoteConfigs: [[url: 'https://github.com/abihi/gobra.git']]])
                }
            }
        }

        stage('Build') {
            steps {
                script {
                    withEnv(['PATH=$PATH:/usr/local/go/bin']) {
                        // Download Go modules and build the Go app
                        sh 'CGO_ENABLED=0 GOOS=linux go build -o /go/bin/gobra ./cmd/gobra'
                    }
                }
            }
        }

        stage('Unit Test') {
            steps {
                script {
                    // Run Go unit tests
                    sh 'go test ./...'
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
