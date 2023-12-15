pipeline {
    agent any

    environment {
        GOPATH = '/go'
    }

    options {
        // Set up GitHub credentials
        scm([
            $class: 'GitSCM', 
            branches: [[name: '*/main']], 
            userRemoteConfigs: [[url: 'https://github.com/abihi/gobra.git']]
        ])
    }

    stages {
        stage('Build') {
            steps {
                script {
                    // Download Go modules and build the Go app
                    sh 'CGO_ENABLED=0 GOOS=linux go build -o /go/bin/gobra ./cmd/gobra'
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
