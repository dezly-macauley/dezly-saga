def main():
    team_score = 0

    player1_score = 20

    # NOTE: breakpoints
  
    # All of the code above will run until it gets to the point below 
    # and then it the shell prompt will change to `(Pdb)`,
    # which is the Python debugger

    # While in this shell you can check the current value of variables

    # (Pdb) team_score 
    # 0
    # (Pdb) player1_score 
    # 20

    # NOTE: breakpoint commands
    # `q` quit the debugger and stop the program
    # `n` move to the next line in the code (does not enter function calls)
    # `s` step into a function call if the next line contains a function
    # `w` This will show you what module / function the debugger is inside
    # `l` This will list the line of code around the current breakpoint

    # NOTE: Moving to the cursor to a line does not execute it
    # You have to move the cursor to the line after it

    # -> player2_score = 10
    # (Pdb) player2_score 
    # *** NameError: name 'player2_score' is not defined
    # (Pdb) 

    breakpoint()
    player2_score = 10

    team_score = player1_score + player2_score
    print(team_score)
    
if __name__ == "__main__":
    main()
