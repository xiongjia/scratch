@echo off

python "%~dp0engine\playgame.py" ^
    -So --engine_seed 42 --player_seed 42 --end_wait=0.25 ^
    --verbose --log_dir game_logs --turns 500 ^
    --map_file "%~dp0engine\maps\maze\maze_04p_01.map" %* ^
    "python ""%~dp0sdk\python\sample_bots\LeftyBot.py"""   ^
    "python ""%~dp0sdk\python\sample_bots\RandomBot.py"""  ^
    "python ""%~dp0sdk\python\sample_bots\HunterBot.py"""  ^
    "python ""%~dp0sdk\python\sample_bots\TimeoutBot.py""" | java -jar engine\visualizer.jar

