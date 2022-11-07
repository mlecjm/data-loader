## Design and Building a Data Loading Engine for Blockchain-enabled Big Data Systems using Hyperledger Fabric
![System_Model](https://github.com/mlecjm/data-loader/blob/main/resources/system_model.png)

### Prerequisites
- Hyperledger Fabric 2.2.x
- Certificate Authority (CA) 1.4.x
- Hyperledger Explorer
- Hyperledger Caliper
- IBM Blockchain Console
- [Install Git](https://git-scm.com/downloads)
- [Install cURL](https://curl.se/download.html)
- Install Docker and Docker Compose
   - Install docker-engine: [https://docs.docker.com/engine/install/](https://docs.docker.com/engine/install/)
   - Install docker-compose: [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)
-

### AWS server setting
- Resource Spec (AWS EC2 instance)
    - Region: ap-northeast-2 (서울)
    - Instance Type : t3a.2xlarge (8 vCPUs, 32GB)
    - Volume Type : 1000GB SSD
    - Amazon Machine Image(AMI) : Ubuntu 20.04 LTS (HVM) x64(x86)
- SSH Connection

### Hyperledger Fabric setup
- FABRIC_RELEASE=2.3.0
- CHANNEL_NAME=ethtxChannel
- PEER_DATABASE_TYPE=golevel
- CHAINCODE_LANGUAGE=go
- CHAINCODE_NAME=ethtxcc
- CHAINCODE_VERSION=0.1
- CHAINCODE_INIT_REQUIRED=true

