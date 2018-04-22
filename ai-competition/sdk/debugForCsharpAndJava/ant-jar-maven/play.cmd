@echo off
cd ..\..\..\..

python "playgame.py" --engine_seed 42 --player_seed 42 --end_wait=0.25 --turntime=1000000 --loadtime=1000000 --log_dir game_logs --log_stream --turns 500 -R -S -I -O -E -v --map_file "maps\maze\maze_p02_01.map" %* "java -jar dist\starter_bots\java\SocketForwarder\target\SocketForwarder.jar -p 18889" "python dist\sample_bots\python\HunterBot.py"

exit 0