# CRB Functional Tests #
# Pre-requisites:
 * docker is installed
 * CRB is deployed on PCF

# Usage:
The tests are executed using a [docker](Dockerfile) container which has the required test tools and the copy repository configured.
  * Create the docker image
  ```
  docker build -t crbtestdockerimage:1 .
  ```
  * Run the docker container
  ```
  docker run -d -p 1111:22 -v "path to functional tests":/root/functionaltests/ -w /root/functionaltests --name "crbtestscontainer" -t "crbtestdockerimage:1"
  ```
  * Run the tests using the above container
  ```
  * docker exec crbtestscontainer jasmine-node main_test_spec.js --config host "<crb-server url>" --config dbname "<" --config dbpassword "hkAIW37zg4T8IfMi" --config dbuser "SxzlF1uAmlIMCpbM" --config dbip "10.106.124.194" --config repoip "<local ip>" --config repoport "1111" --config repouser "root" --config repopassword "screencast" --config crbname "crb-server" --config cfusername "CNDPDev-OrgMgr" --config cfpassword "CNDPD3vOrgMgr" --config cfendpoint "api.scf.isus.emc.com" --config version "0.1"
  ```
 
 

