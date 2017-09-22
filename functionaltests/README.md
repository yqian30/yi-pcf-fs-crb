# CRB Functional Tests #
# Pre-requisites:
 * [Docker](https://docs.docker.com/engine/installation/) is installed and running.
 * CRB is deployed and running on PCF.

# Usage:
The tests are executed using a [container](Dockerfile) which has the required test tools and the copy repository configured.
  * Build the docker image ocopea_crb_test_tool (feel free to change names and tag with version, and change the following commands accordingly)
  ```
  docker build -t ocopea_crb_test_tool .
  ```
  * Find an unused port on the local host that will be used to forward traffic to the container's ssh port i.e. port 22. Ocopea's CRB service uses SFTP to securely transfer copies on this port.
  
  On windows or linux based machines to get the list of *used* local port, run
  ```
  netstat -ap
  ```
  and accordingly choose an unused port for SSH forwarding. This port will be used in later commands.

  * Run the docker container with name ocopea_crb_test_tool
  ```
  docker run -d -p <ssh_fwd_port>:22 -v <absolute_path_to_functionaltests>:/root/functionaltests/ --name ocopea_crb_test_tool ocopea_crb_test_tool
  ```
  * Run the tests using the above container
  ```
  * docker exec ocopea_crb_test_tool jasmine-node main_test_spec.js --config host "<crb_app_url>" --config dbname "<p-mysql_name>" --config dbpassword "<p-mysql_password>" --config dbuser "<p-mysql_username>" --config dbip "<p-mysql_hostname>" --config repoip "<host_address>" --config repoport "<ssh_fwd_port>" --config repouser "root" --config repopassword "screencast" --config crbname "<crb_app_name>" --config cfusername "<pcf_username>" --config cfpassword "<pcf_password>" --config cfendpoint "<pcf_api_endpoint>" --config version "<crb_app_version>"
  ```
  **Note**:
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
  ...
  }
  ... 
  ```