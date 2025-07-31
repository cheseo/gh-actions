pipeline {
	stages {
		stage('build'){
		agent { dockerfile true }
			steps {
				sh 'go build'
			}
			steps {
			      sh 'go test'
			}
		}
		stage('deploy'){
			steps {
			      sh 'sudo -Hu ubuntu sh -c "cd; cd gh-actions; ./update"'
			}
		}
	}
}
