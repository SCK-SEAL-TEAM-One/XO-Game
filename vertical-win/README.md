# XO-GAME VERTICAL WIN
## Instruction
    เกมส์ XO นี้ เป็นเกมที่ประกอบไปด้วยผู้เล่นจำนวน 2 คน โดยผู้เล่นจะทำการเลือกตำแหน่งเพื่อทำสัญลักษณ์ในตำแหน่งของตารางขนาด 3X3 ช่อง และจะสลับกันเล่นจนกว่าจะมีผู้เล่นคนใดคนหนึ่งมีสัญลักษณ์เป็นแนวเส้นตรงเดียวกันครบ 3 ตำแหน่ง

## WINNING RULE
    ผู้เล่นจะต้องทำสัญลักษณ์เรียงกันในแนวตั้งครบ 3 ตำแหน่ง เช่น

```````````````
x | o |  
¯¯¯¯¯¯¯¯¯¯
x | o |
¯¯¯¯¯¯¯¯¯¯
x |   |

```````````````

## HOW TO RUN PROGRAM
### Set Go Path
``````
set GOPATH=%CD%
``````

### Do it before start program
``````
install model
install verticle
go test ./...
go biuld -o ./bin/vertical-win.exe src/main (on windrow)
go biuld -o ./bin/vertical-win src/main (on mac)
open file bin/vertival-win
``````