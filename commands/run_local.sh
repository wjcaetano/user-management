#!/usr/bin/env bash
chmod +x /commands/wait-for-it.sh

/commands/wait-for-it.sh testlocal 3306 60
/commands/run_migrations.sh
/commands/run.sh
