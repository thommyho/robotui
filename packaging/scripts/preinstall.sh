#!/bin/sh
#
# Executed before the installation of the new package
#
#   $1=install              : On installation
#   $1=upgrade              : On upgrade

set -e

robotui_USER=robotui
robotui_GROUP=robotui
robotui_HOME="/var/lib/$robotui_USER"
RESTART_FLAG_FILE="/tmp/.restartrobotuiOnUpgrade"

copyDbToUserDir() {
  CURRENT_USER=$(systemctl show -pUser robotui | cut -d= -f2)
  if [ -z "$CURRENT_USER" ]; then
  	CURRENT_USER=root
  fi
  CURRENT_HOME=$(getent passwd "$CURRENT_USER" | cut -d: -f6)
  COPIED_FLAG="$CURRENT_HOME/.robotui/.copiedTorobotuiUser"
  if [ -f "$CURRENT_HOME/.robotui/robotui.db" ] && [ ! -f "$COPIED_FLAG" ]; then
    if [ -f "$robotui_HOME/robotui.db" ]; then
      echo "--------------------------------------------------------------------------------"
      echo "Not copying $CURRENT_HOME/.robotui/robotui.db to $robotui_HOME/robotui.db, since there is"
      echo "already a database there."
      echo "Either delete one of the databases or run 'touch $COPIED_FLAG' to keep both,"
      echo "then restart installation."
      echo "Hint: usually the larger one is the one to keep."
      ls -la "$CURRENT_HOME/.robotui/robotui.db" "$robotui_HOME/robotui.db"
      echo "--------------------------------------------------------------------------------"
      exit 1
    else
      cp -Rp "$CURRENT_HOME"/.robotui/robotui.db "$robotui_HOME"
    fi
    chown "$robotui_USER:$robotui_GROUP" "$robotui_HOME/robotui.db"
    touch "$COPIED_FLAG"
    if [ -n "$(ls -A /etc/systemd/system/robotui.service.d 2>/dev/null)" ]; then
        echo "--------------------------------------------------------------------------------"
		echo "You have overrides defined in /etc/systemd/system/robotui.service.d."
		echo "This update changes the robotui user to 'robotui' (from root) and the database file"
		echo "to '/var/lib/robotui/robotui.db"
		echo "Make sure that you neither override 'User' nor 'ExecStart'"
		echo "Hint: you can delete all overrides with 'systemctl revert robotui'"
		echo "As a precaution, robotui is not started even if it was previously started."
        echo "--------------------------------------------------------------------------------"
		rm -f "$RESTART_FLAG_FILE"
	else
        echo "--------------------------------------------------------------------------------"
		echo "NOTE: robotui user has changed from $CURRENT_USER to $robotui_USER, db has been copied to new"
		echo "directory $robotui_HOME/robotui.db, old db in $CURRENT_USER/.robotui has been retained."
      	echo "--------------------------------------------------------------------------------"
    fi
  fi
  return 0
}

if [ "$1" = "install" ] || [ "$1" = "upgrade" ]; then
	if [ -d /run/systemd/system ] && /bin/systemctl status robotui.service > /dev/null 2>&1; then
	  deb-systemd-invoke stop robotui.service >/dev/null || true
	  touch "$RESTART_FLAG_FILE"
	fi
    if ! getent group "$robotui_GROUP" > /dev/null 2>&1 ; then
      addgroup --system "$robotui_GROUP" --quiet
    fi
    if ! getent passwd "$robotui_USER" > /dev/null 2>&1 ; then
      adduser --quiet --system --ingroup "$robotui_GROUP" \
      --disabled-password --shell /bin/false \
      --gecos "robotui runtime user" --home "$robotui_HOME" "$robotui_USER"
      chown -R "$robotui_USER:$robotui_GROUP" "$robotui_HOME"
      adduser --quiet "$robotui_USER" dialout
    else
      adduser --quiet "$robotui_USER" dialout
      homedir=$(getent passwd "$robotui_USER" | cut -d: -f6)
      if [ "$homedir" != "$robotui_HOME" ]; then
      	mkdir -p "$robotui_HOME"
      	chown "$robotui_USER:$robotui_GROUP" "$robotui_HOME"
        process=$(pgrep -u "$robotui_USER") || true
        if [ -z "$process" ]; then
          usermod -d "$robotui_HOME" "$robotui_USER"
          if [ -f "$homedir/.robotui/robotui.db" ]; then
          	cp "$homedir/.robotui/robotui.db" "$robotui_HOME" && touch "$homedir/.robotui/.copiedTorobotuiUser"
          fi
        else
      	  echo "--------------------------------------------------------------------------------"
          echo "Warning: robotui's home directory is incorrect ($homedir)"
          echo "but can't be changed because at least one other process is using it."
          echo "Stop offending process(es), then restart installation."
          echo "Hint: You can list the offending processes using 'pgrep -u $robotui_USER -a'"
          echo "Note that you should NOT use the robotui user as login user, since that will"
          echo "inevitably lead to this error."
          echo "in that case, please create a different user as login user."
          echo "--------------------------------------------------------------------------------"
          exit 1
        fi
      fi
    fi
fi

if [ "$1" = "upgrade" ]; then
    copyDbToUserDir
fi

exit 0
