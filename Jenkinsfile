node {
  stage('SCM Checkout') {
    git url: https://github.com/agiratech/golang-jwt-auth
  }
  stage('Build Docker Image') {
    sh 'docker build -t reddysai/golang-jwt-auth .'
  }
  stage('Push Docker Image') {
    withCredentials([string(credentialsId: 'docker-pwd', variable: 'dockerHubPwd')]) {
      sh "docker login -u reddysai -p ${dockerHubPwd}"
    }
  }
  stage('Run Container on Dev Server') {
    def dockerRun = 'docker run -p 8000:8000 -d reddysai/golang-jwt-auth'
    sshagent(['dev-server']) {
      sh "ssh -o StrictHostKeyChecking=no ec2-user@x.x.x.x ${dockerRun}"
    }
  }
}