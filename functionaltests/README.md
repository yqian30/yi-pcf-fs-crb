# CRB Functional Tests #
# Pre-requisites:
 * docker is installed
 * CRB is deployed on PCF

# Usage:
  * docker build -t crbtestdockerimage:1 .
  * docker run -d -p 1111:22 -v "path to functional tests":/root/functionaltests/ -w /root/functionaltests --name "crbtestscontainer" -t "crbtestdockerimage:1"

  * docker exec crbtestscontainer jasmine-node main_test_spec.js --config host "<crb-server url>" --config dbname "<" --config dbpassword "hkAIW37zg4T8IfMi" --config dbuser "SxzlF1uAmlIMCpbM" --config dbip "10.106.124.194" --config repoip "<local ip>" --config repoport "1111" --config repouser "root" --config repopassword "screencast" --config crbname "crb-server" --config cfusername "CNDPDev-OrgMgr" --config cfpassword "CNDPD3vOrgMgr" --config cfendpoint "api.scf.isus.emc.com" --config version "0.38"
 
 

# Running tests using pipeline on pcfdev:
This can be used to test your changes local before pushing them to remote repo.
 * Ensure gitlab-runner is installed.
 * Ensure pcfdev is installed and is running.
 * Create binaries directory in your git repo and copy a built crb binary.
 * Modify the CRB_BIN_LOC variable value in CI config to the absolute path of binaries directory.
 * Using gitlab-runner run the test job:
 ```
gitlab-runner exec shell pcfdev_test
```