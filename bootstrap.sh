#!/bin/bash

echo "Installing Software"
apt-key adv \
	--keyserver hkp://p80.pool.sks-keyservers.net:80 \
	--recv-keys 58118E89F3A912897C070ADBF76221572C52609D \
	>/dev/null 2>&1

cat >/etc/apt/sources.list.d/docker.list <<EOF
deb https://apt.dockerproject.org/repo ubuntu-vivid main
EOF

apt-get update >/dev/null 2>&1
apt-get install -y \
	linux-image-extra-$(uname -r) \
	docker-engine \
	emacs24-nox >/dev/null 2>&1

docker pull filefrog/sandbox
usermod -aG docker vagrant

curl -LSs https://storage.googleapis.com/golang/go1.5.3.linux-amd64.tar.gz | tar -C /usr/local -xz
cat > /home/vagrant/.bash_aliases <<'EOF'
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
EOF

echo "Setting up User Environment"
cat > /tmp/user-setup.sh <<'EOF'
#!/bin/bash
source /home/vagrant/.bash_aliases

mkdir -p /home/vagrant/go
export GOPATH=/home/vagrant/go
go get github.com/tools/godep
go get github.com/golang/lint/golint
go get github.com/onsi/ginkgo
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega

#go get github.com/starkandwayne/training/go/longshoreman/backend
EOF
sudo -u vagrant bash /tmp/user-setup.sh
rm -f /tmp/user-setup.sh

echo "DONE"
