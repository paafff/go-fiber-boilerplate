// pipeline {
//     agent any

//     stages {
//         stage('Checkout') {
//             steps {
//                 git branch: 'main', url: 'https://github.com/paafff/go-fiber-boilerplate.git'
//             }
//         }
//         stage('Build Docker Image') {
//             steps {
//                 sh 'docker-compose down'
//                 sh 'docker-compose up --build -d'
//             }
//         }
//     }
// }



pipeline {
    agent any

    stages {
        stage('Clone repository') {
            steps {
                git branch: 'main', url: 'https://github.com/paafff/go-fiber-boilerplate.git'
            }
        }

        // stage('Build Docker Image') {
        //     steps {
        //         script {
        //             def app = docker.build('your-app-image')
        //         }
        //     }
        // }

        stage('Run Docker Compose') {
            steps {
                sh 'docker-compose -f /path/to/your/docker-compose.yml up -d --build'
            }
        }
    }

    post {
        always {
            cleanWs() // Clean up workspace after build
        }
    }
}
