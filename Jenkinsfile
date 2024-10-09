pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/paafff/go-fiber-boilerplate.git'
            }
        }
        stage('Build Docker Image') {
            steps {
                sh 'docker-compose down'
                sh 'docker-compose up --build -d'
            }
        }
    }
}
