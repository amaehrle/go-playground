
Vagrant.configure(2) do |config|
  config.vm.box = 'ubuntu/trusty64'
  config.vm.provision 'file', source: '~/.gitconfig', destination: '.gitconfig'
  config.ssh.forward_agent = true

  config.vm.provider "virtualbox" do |v|
    v.name = "Go Playground"
    v.memory = 2048
  end

  config.vm.network 'forwarded_port', guest: 8000, host: 8000, auto_correct: true
  config.vm.network 'forwarded_port', guest: 8888, host: 8888, auto_correct: true

  config.vm.synced_folder '.', '/home/vagrant/workspace/src/github.com/amaehrle/go-playground'

  config.vm.provision 'ansible' do |ansible|
    ansible.playbook = 'ansible/go_playground.yml'
  end
end
