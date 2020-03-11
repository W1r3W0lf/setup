# setup

My config files and package lists to bring a distro up from base line.

Not yet working.

This program does two things.
 * It Installs packages
 * It Installs settings

The packages are defined in a yaml file.

It can install system packages, python packages, node packages, any package manager with a CLI interface.
The package install string is defined in the yaml file.


The settings are moved from the target folder to the usres home with a . added onto the front.
This allows for .bashrc, .profile, .tmux.conf, .config/*, .ssh/config. Any user settings, even if they are in a folder.

I am planning on adding a scripts section where you can drop in shell scripts to be run before or after install.

And I would also like to add 

Things I would like to add
 * Install scripts folder
 * Create custom files in home folder
 * Default git cloner
 * Client/Server setup to import settings over the local network.
