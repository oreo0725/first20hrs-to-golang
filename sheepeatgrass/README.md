# sheepeatgrass

## A Simulation world - Sheeps and Grass
This is a simulator of a simple world, there are only 2 creatures in this world, Sheeps and Grass.

### Rules 
The world has following rules to perform day by day
- The world composed by pixel, which has dimension `M`x`N` pixels.
- There are `P` Sheeps and `Q` Grass at the begining of the world. They are placed randomly overall the world.
- Each Sheep has 20 healthPoints.
- 1 healthPoint is consumed for each creature when one day passed
- When a Sheep is nearby a Grass, he will eat the grass to refill its healthPoints, each grass ate to increase 2 healthPoints. If a sheep is surround by 2 grass, then he will eat them all, then recover for 4 healthPoints.
- **Nearby** which means they are neighbour, in anyone of 4 direction(East, South, West, North) that distance is 1.
- A Sheep randomly choose a direction to move 1 step to it's **nearby** space if that space is empty. otherwise, this sheep won't move if there aren't empty nearby space.
- A Sheep should eat grass rather than move a step, because only 1 action allowed in 1 day
- Dead: 
  - Sheep runs out of it's healthPoint, it dies.
  - Once the Grass is eaten by a Sheep, it dies.
  - When the creature lives long enough, it dies. Sheep only can live for maximum 70 days, Grass for 6 days.
- Give birth:
  - Between 50th ~ 55th day for Sheep, it give birth to a new Sheep each day.
  - Between 3rd ~ 5th day for Grass, it gives birth to a new Grass each day.
