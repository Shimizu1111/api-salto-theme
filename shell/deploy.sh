#!/bin/sh

# @以降にAWSのIPアドレスを記載する
# ssh -o "StrictHostKeyChecking no" -i "../shmizu1111.pem" ec2-user@13.115.122.171
# ssh -i "../shmizu1111.pem" ec2-user@13.115.122.171 git -C api-salto-theme status


# ファイル転送(ローカルからリモート)
scp  ./hello.txt ec2-user@13.115.122.171:~/api-salto-theme/cmd
# scp -i "../shmizu1111.pem" ./main ec2-user@13.115.122.171:~/api-salto-theme/main

# ファイル転送(リモートからリモート)
# scp user1@192.168.10.1:/home/user/tmp/file1 user2@192.168.10.2:/home/user/tmp/

# ssh -i "../shmizu1111.pem" ec2-user@13.115.122.171 pwd
# ssh -i "../shmizu1111.pem" ec2-user@13.115.122.171 pwd
# ssh -i "../shmizu1111.pem" ec2-user@13.115.122.171 pwd
# exit