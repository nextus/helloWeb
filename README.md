# Yet another hello world application

Simple stateless web application tool just to play around with various orchestration tools like docker and k8s.

## Docker (Dockerfile)
Dockerfile using pipeline with two containers. At first, the binary is built using a developer container and then is copied to the tiny runtime container. Alpine docker image does not include Glibc so we have to compile our binary statically. This way keeps our runtime environment away from contamination with useless dev files.

## Kubernetes (k8s.yaml)
Simple manifest starts 3 pods with the web application. HTTP health checks are used for cluster consistency. It prevent cases when application instance stuck internally due to deadlock but the process still exists.

## Packer (packer/centos7-ec2.json)
Prepare Amazon Machine Image (AMI) using Ansible provisioner. It starts Centos 7 official image from AWS Market (you have to apply Centos EULA in advance) and runs Ansible from your host system with the provided playbook. As a result, you get your personal AMI stored based on Centos 7 image. Web application starts automatically during boot stage using systemd. A related unit file is stored in the templates directory.