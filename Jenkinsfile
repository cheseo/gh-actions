pipeline {
	agent any
	stages {
		stage('build'){
			agent { dockerfile true }
			steps {
				sh 'go build'
				sh 'go test'
			}
		}
		stage('deploy'){
			steps {
				sh '''
				sudo -Hu ubuntu sh -c "cd ~/gh-actions; ./update"
				'''
			}
		}
	}
}
