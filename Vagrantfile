
Vagrant.configure(2) do |config|
  config.vm.box = 'ubuntu/trusty64'
  config.vm.provision 'file', source: '~/.gitconfig', destination: '.gitconfig'
  config.ssh.forward_agent = true

  config.vm.define 'api_box', primary: true do |api_box|
    api_box.vm.provider "virtualbox" do |v|
      v.name = "Go Playground"
      v.memory = 2048
    end

    api_box.vm.network 'private_network', ip: '192.168.100.100'
    api_box.vm.network 'forwarded_port', guest: 80, host: 8080, auto_correct: true
    api_box.vm.network 'forwarded_port', guest: 3000, host: 3000, auto_correct: true
    api_box.vm.network 'forwarded_port', guest: 6379, host: 6379, auto_correct: true

    api_box.vm.synced_folder '.', '/home/vagrant/workspace/src/'

    api_box.vm.provision 'ansible' do |ansible|
      ansible.playbook = 'ansible/go_playground.yml'
    end
  end

  if Vagrant.has_plugin?('vagrant-proxyconf') && ENV.has_key?('http_proxy')
    config.proxy.http = ENV['http_proxy']
    config.proxy.no_proxy = 'localhost,127.0.0.1,.example.com,.oracle.com,192.168.0.0/16'
    config.apt_proxy.http = ENV['http_proxy']
    config.apt_proxy.https = 'DIRECT'
  end

end
