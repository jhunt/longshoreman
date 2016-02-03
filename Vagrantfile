# docker environment
Vagrant.configure('2') do |config|
  config.vm.box = 'jhunt/longshoreman'
  config.vm.network "forwarded_port", guest: 8080, host: 8184, auto_correct: true
  config.vm.synced_folder 'phase1/', '/home/vagrant/phase1'
  config.vm.synced_folder 'phase2/', '/home/vagrant/phase2'
  config.vm.synced_folder 'phase3/', '/home/vagrant/phase3'
end
