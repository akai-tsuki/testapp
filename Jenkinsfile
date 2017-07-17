node {
  try {
    // Need to replace the '%2F' used by Jenkins to deal with / in the path (e.g. story/...)
    // because tests that do getResource will escape the % again, and the test files can't be found.
    // See https://issues.jenkins-ci.org/browse/JENKINS-34564 for more.
    ws("workspace/${env.JOB_NAME}/${env.BRANCH_NAME}".replace('%2F', '_')) {
      // Mark the code checkout 'stage'....
      stage 'Checkout'
      checkout scm

      stage('build') {
        try {
          sh date
        
        } finally {
          sh echo "test"
        }
      }
    }
  } catch (e) {
    // TODO: チャットにビルドエラーを通知する
    throw e
  }
}
