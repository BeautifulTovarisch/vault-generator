# Simulate K8s cluster

Vagrant.configure("2") do |config|
  config.vm.provider "virtualbox" do |os|
    os.cpus = 2
    os.memory = 1012
  end

  # Kubernetes master host
  config.vm.define "master" do |master|
    master.vm.box = "ubuntu/bionic64"
    master.vm.network "private_network", ip: "192.168.0.10"
    master.vm.hostname = "master"
  end

  # Nodes
  config.vm.define "client" do |client|
    client.vm.box = "ubuntu/bionic64"
    client.vm.network "private_network", ip: "192.168.0.20"
    client.vm.hostname = "client"
  end

  # config.vm.define "server" do |server|
  #   server.vm.box = "ubuntu/bionic64"
  #   server.vm.network "private_network", ip: "192.168.0.30"
  #   server.vm.hostname = "server"
  # end

  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "ansible/playbook.yml"

    ansible.host_vars = {
      "master" => { "host_ip": "192.168.0.10" },
      "client" => { "host_ip": "192.168.0.20" },
      "server" => { "host_ip": "192.168.0.30" }
    }

    ansible.groups = {
      "controller" => [ "master" ],
      "nodes" => [ "client", "server" ]
    }
  end
end
