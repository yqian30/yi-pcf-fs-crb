# CRB Functional Tests #
# Pre-requisites:
 * docker is installed
 * CRB is deployed on PCF

# Usage:
The tests are executed using a [docker](Dockerfile) container which has the required test tools and the copy repository configured.
  * cd functionaltests

  * Create the docker image
  ```
  docker build -t <crb_test_docker_image>:<tag> .
  ```
  * Find an unused port on the local host, for example 1111. To get the list of used local port, run
  ```
  netstat -ap
  ```

  * Run the docker container
  ```
  docker run -d -p <unused_local_port>:22 -v <path_to_functionaltests>:/root/functionaltests/ -w /root/functionaltests --name <crb_test_container_name> -t "<crb_test_docker_image>:<tag>"
  ```
  * Run the tests using the above container
  ```
  * docker exec <crb_test_container_name> jasmine-node main_test_spec.js --config host "<crb_app_url>" --config dbname "<p-mysql_name>" --config dbpassword "<p-mysql_password>" --config dbuser "<p-mysql_username>" --config dbip "<p-mysql_hostname>" --config repoip "<local_ip_address>" --config repoport "<unused_local_port>" --config repouser "root" --config repopassword "screencast" --config crbname "<crb_app_name>" --config cfusername "<pcf_username>" --config cfpassword "<pcf_password>" --config cfendpoint "<pcf_endpoint>" --config version "<crb_app_version>"
  ```
  Note: 
  <p-mysql_name>, <p-mysql_username>, <p-mysql_password> and <p-mysql_hostname> are the credentials of the p-mysql service which is bound to the crb app. They can be retrieved by running
  ```
  cf env <crb_app_name>
  ```

  Here is an sample output:
  ```
  System-Provided:
  {
   "VCAP_SERVICES": {
    "p-mysql": [
     {
      "credentials": {
      "hostname": "10.106.124.194",
      "jdbcUrl": "jdbc:mysql://10.106.124.194:3306/cf_b5ffc129_539c_42cd_ac9b_5524f93502c9?user=njT6gHoQZrca1Yzw\u0026password=8FIWMg5otqKpWKvM",
      "name": "cf_b5ffc129_539c_42cd_ac9b_5524f93502c9",
      "password": "8FIWMg5otqKpWKvM",
      "port": 3306,
      "uri": "mysql://njT6gHoQZrca1Yzw:8FIWMg5otqKpWKvM@10.106.124.194:3306/cf_b5ffc129_539c_42cd_ac9b_5524f93502c9? reconnect=true",
      "username": "njT6gHoQZrca1Yzw"
     },
     "label": "p-mysql",
     "name": "crb-mysql",
     "plan": "pre-existing-plan",
     "provider": null,
     "syslog_drain_url": null,
     "tags": [
      "mysql",
      "relational"
     ],
     "volume_mounts": []
    }
   ]
  }
  ... 
  ```

 

 

