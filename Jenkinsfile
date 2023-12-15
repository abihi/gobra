pipeline {
    agent any

    environment {
        GOPATH = '/go'
        PATH+EXTRA = "$GOPATH/bin"
    }

    stages {
        stage('Install Go') {
            steps {
                script {
                    sh 'curl -fsSL https://get.sdkman.io | bash'
                    sh 'source "$HOME/.sdkman/bin/sdkman-init.sh"'
                    sh 'sdk install go'
                }
            }
        }

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
