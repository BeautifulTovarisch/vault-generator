# Simulate K8s cluster

Vagrant.configure("2") do |config|
  config.vm.provider "virtualbox" do |os|
    os.cpus = 2
    os.memory = 1024
  end

  # Kubernetes master host
  config.vm.define "master" do |master|
    ip = "192.168.0.10"

    master.vm.box = "ubuntu/bionic64"
    master.vm.network "private_network", ip: ip
    master.vm.hostname = "master"

    # Setup master node via ansible
    master.vm.provision "ansible" do |ansible|
      ansible.playbook = "ansible/master.playbook.yml"
      ansible.extra_vars = {
        node_ip: ip
      }
    end
  end

  # Nodes
  # TODO :: Structure as roles; run ansible completely from master node
  config.vm.define "client" do |client|
    ip = "192.168.0.20"

    client.vm.box = "ubuntu/bionic64"
    client.vm.network "private_network", ip: ip
    client.vm.hostname = "client"

    client.vm.provision "ansible" do |ansible|
      ansible.playbook = "ansible/node.playbook.yml"
      ansible.extra_vars = {
        node_ip: "192.168.0.20"
      }
    end
  end

  config.vm.define "server" do |server|
    ip = "192.168.0.30"

    server.vm.box = "ubuntu/bionic64"
    server.vm.network "private_network", ip: ip
    server.vm.hostname = "server"

    server.vm.provision "ansible" do |ansible|
      ansible.playbook = "ansible/node.playbook.yml"
      ansible.extra_vars = {
        node_ip: ip
      }
    end
  end
end
