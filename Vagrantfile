PROVIDER_DEFAULTS = {
    libvirt: {
        fedora: {
            box_name: "fedora-22",
            box_url: "https://dl.fedoraproject.org/pub/fedora/linux/releases/22/Cloud/x86_64/Images/Fedora-Cloud-Base-Vagrant-22-20150521.x86_64.vagrant-libvirt.box"
        }
    },
    virtualbox: {
        fedora: {
            box_name: "fedora-22",
            box_url: "https://dl.fedoraproject.org/pub/fedora/linux/releases/22/Cloud/x86_64/Images/Fedora-Cloud-Base-Vagrant-22-20150521.x86_64.vagrant-virtualbox.box"
        },
        ubuntu: {
            box_name: "ubuntu/vivid64",
            box_url: "https://cloud-images.ubuntu.com/vagrant/vivid/current/vivid-server-cloudimg-amd64-vagrant-disk1.box"
        }
    }
}

DISTRO = ENV['VAGRANT_RKT_DISTRO'] || "ubuntu"

Vagrant.configure('2') do |config|
    # grab Ubuntu 15.04 official image
    # config.vm.box = "ubuntu/vivid64" # Ubuntu 15.04

    # fix issues with slow dns http://serverfault.com/a/595010
    config.vm.provider :virtualbox do |vb, override|
        vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
        vb.customize ["modifyvm", :id, "--natdnsproxy1", "on"]
        # add more ram, the default isn't enough for the build
        vb.customize ["modifyvm", :id, "--memory", "1024"]
        override.vm.box = PROVIDER_DEFAULTS[:virtualbox][DISTRO.to_sym][:box_name]
        override.vm.box_url = PROVIDER_DEFAULTS[:virtualbox][DISTRO.to_sym][:box_url]
    end

    config.vm.provider :libvirt do |libvirt, override|
        libvirt.memory = 1024
        override.vm.box = PROVIDER_DEFAULTS[:libvirt][DISTRO.to_sym][:box_name]
        override.vm.box_url = PROVIDER_DEFAULTS[:libvirt][DISTRO.to_sym][:box_url]
    end

    config.vm.synced_folder ".", "/vagrant", type: "rsync"
    config.vm.provision :shell, :privileged => true, :path => "scripts/install-vagrant.sh"
end
