# to run: `vhs < demo.tape`
# https://github.com/charmbracelet/vhs

Output demo.gif

Require echo

Set Shell "bash"
Set FontSize 12
Set Width 1200
Set Height 600

Hide
Type "cd ../" Enter
Type "clear" Enter
Show

Type "./time-builder generate" Sleep 2000ms  Enter

Sleep 3s

Hide
Type "cd yourproject" Enter
Show

Type "cat workflows.py" Sleep 2000ms Enter

Sleep 3s

Hide
Type "cd ../" Enter
Type "clear" Enter
Show

Type "./time-builder generate --child-workflow ChildWorkflowName  --signal Hello --signal-count 1 --query Bye --query-count 1" Sleep 3000ms Enter

Sleep 2000ms

Hide
Type "cd yourproject" Enter
Type "clear" Enter
Show

Type "cat workflows.py" Sleep 2000ms Enter

Sleep 3s