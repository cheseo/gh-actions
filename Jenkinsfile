pipeline {
	agent { docker { image 'golang:1.24.5-bookworm' } }
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
