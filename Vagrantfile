# Simulate K8s cluster

Vagrant.configure("2") do |config|
  config.vm.provider "virtualbox" do |os|
    os.cpus = 2
    os.memory = 1024
  end


  # Kubernetes master host
  config.vm.define "master" do |master|
    master.vm.box = "ubuntu/bionic64"
    master.vm.network "private_network", ip: "192.168.0.10"
    master.vm.hostname = "master"
  end

  # Nodes
  # TODO :: Structure as roles; run ansible completely from master node
  config.vm.define "client" do |client|
    client.vm.box = "ubuntu/bionic64"
    client.vm.network "private_network", ip: "192.168.0.20"
    client.vm.hostname = "client"
  end

  config.vm.define "server" do |server|
    server.vm.box = "ubuntu/bionic64"
    server.vm.network "private_network", ip: "192.168.0.30"
    server.vm.hostname = "server"
  end

  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "ansible/playbook.yml"
    ansible.groups = {
      "master" => { "ansible_host": "192.168.0.10" },
      "client" => { "ansible_host": "192.168.0.20" },
      "server" => { "ansible_host": "192.168.0.30" },
      "nodes" => ["client", "server"]
    }
  end
end
