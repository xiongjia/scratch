import subprocess
import os

os.system('''python "engine/playgame.py" --engine_seed 42 --player_seed 42 --end_wait=0.25 --verbose --log_dir game_logs --turns 500 --map_file="./engine/maps/maze/maze_04p_01.map" "python ./sdk/python/sample_bots/GreedyBot.py" "python ./sdk/python/sample_bots/LeftyBot.py" "python ./sdk/python/sample_bots/HunterBot.py" "python ./sdk/python/sample_bots/RandomBot.py"''')