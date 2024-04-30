#!/bin/bash

echo -e "This script is made for Linux distributions which are using dpkg.\n---"

### VARIABLES ###

# Desired WireGuard Interface configuration directory
des_conf_path="/etc/wireguard/iface-config"

### CHECKING FOR ROOT PRIVILEGES ###

if [ "$EUID" -ne 0 ]; then
    echo "sudo or root privileges required, please re-run with sufficient privileges."
    exit 1
fi

### BEGIN OF WIREGUARD SCRIPT ###

function setup_wireguard_dir() {
    if [ -d /etc/wireguard ]; then
        echo "Found /etc/wireguard."
    else
        printf "WireGuard directory is not present. Want to create it? [Y/n] "
        read wDir
        if [ "$(echo "$wDir" | tr '[:upper:]' '[:lower:]')" = "y" ] || [ -z "$wDir" ]; then
            mkdir "/etc/wireguard"
            mkdir "/etc/wireguard/iface-config"
            mkdir "/etc/wireguard/iface-client"
        else
            echo "Exiting..."
            exit 1
        fi
        echo "Re-running for validation."
        setup_wireguard_dir
    fi
}

function setup_wireguard_package() {
    if dpkg -s wireguard &> /dev/null; then
        # Tree for already installed WireGuard package.
        echo "Wireguard is installed."
        
    else
        # Tree for WireGuard setup.
        echo "WireGuard not found."
        
        printf "Would you like to install it? [Y/n] "
        read wInstall
        if [ "$(echo "$wInstall" | tr '[:upper:]' '[:lower:]')" = "y" ] || [ -z "$wInstall" ]; then
            echo "Installing..."
            apt install wireguard wireguard-tools &> /dev/null
        else
            echo "Exiting..."
            exit 1
        fi
        echo "Re-running for validation."
        setup_wireguard_package
    fi
}

function modify_wg-quick() {
    if grep -q "$des_conf_path" /usr/bin/wg-quick; then
        echo "wg-quick correctly modified."
    else
        sed -i 's|CONFIG_FILE="/etc/wireguard/\$CONFIG_FILE\.conf"|CONFIG_FILE="'"$des_conf_path"'/\$CONFIG_FILE\.conf"|' /usr/bin/wg-quick
    fi
}

### RUNNING CHAIN ###



echo "Checking for WireGuard package installation..."
setup_wireguard_package

echo "Modifying wg-quick..."
modify_wg-quick

echo "---"
echo "WireGuard should be good to Go for the WireGuard Remote."