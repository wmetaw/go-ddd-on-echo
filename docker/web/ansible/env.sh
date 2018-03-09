if [ -z "$GOPATH" ]; then
    echo "export GOPATH=/go" >> /root/.bash_profile
    echo "export GOPATH=/go" >> /home/ec2-user/.bash_profile

    echo "export PATH=$PATH:/go/bin:/usr/local/go/bin" >> /root/.bash_profile
    echo "export PATH=$PATH:/go/bin:/usr/local/go/bin" >> /home/ec2-user/.bash_profile

    source /root/.bash_profile
    source /home/ec2-user/.bash_profile
fi
