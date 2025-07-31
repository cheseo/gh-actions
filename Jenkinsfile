pipeline {
	agent { docker { image 'golang:1.24.5-alpine3.21' } }
	stages {
		stage('build'){
			steps {
				sh 'go build'
			}
		}
		stage('test'){
			steps {
			      sh 'go test'
			}
		}
	}
}
