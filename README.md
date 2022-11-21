#rpi-cluster-manager

Helps to bootstrap, setup and manage a rpi cluster to run your workloads. 

This tools assumes that you will be setting up raspberry pi(rpi hereon) to run your Kubernetes workloads. Along with kubernetes, you might also require some basic bootstraping scripts, tools and frameworks to be installed on your rpi's. This tool is made just for that, to help you with all such customizations.


It has 4 parts:

- `Installer`: helps you install various tools using `sudo apt-get` commnad.
- `Scripter`: Runs all scripts present in `/scripts` directory in alphabetical order.
- `Configurer`: Adds/appends text/configuration to existing files. It also helps in changing/reassigning hostnames post bootstraping is done.
- `Notifier`: Notifies a discord channel with the progress being made.


