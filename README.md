aws-ecr-login is a very simple tool to login [ECR] without using [aws-cli]

#### Why?

aws-cli is powerful but too big. Not all functions I will use. Sometimes, I just want to pull images from AWS ECR and run them. 

#### How to Use It?

```bash
$ go build
# aws-ecr-login is created
$ ./aws-ecr-login -region <region_XXXX-XXXXX-1>
# Some output you expect
docker login -u AWS -p gsasdxXDGSGHWJUSTEXAMPLEEHEWSGHSGAGA.....XCGdsfsdfsdf0= https://012345678910.dkr.ecr.us-east-1.amazonaws.com
```

[ECR]:https://aws.amazon.com/ecr/
[aws-cli]:https://aws.amazon.com/cli/
