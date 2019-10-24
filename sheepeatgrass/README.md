# Sheep eats grass 羊吃草

## A Simulation world - Sheeps and Grass
This is a simulator of a simple world, there are only 2 creatures in this world, Sheeps and Grass.

### Rules 
The world has following rules to perform day by day
- The world composed by pixel, which has dimension `M`x`N` pixels.
- There are `P` Sheeps and `Q` Grass at the begining of the world. They are placed randomly overall the world.
- Each Sheep has 20 healthPoints.
- Pass a day
  - 1 healthPoint is consumed for each creature when one day passed
  - Each creature only has one movement each day. Either Breed or Move or StandAround.  
- Dead: 
  - Sheep runs out of it's healthPoint, it dies.
  - Once the Grass is eaten by a Sheep, it dies.
  - When the creature lives long enough, it dies. Sheep only can live for maximum 70 days, Grass for 6 days.
- Give birth:
  - Between **50**th ~ **55**th day for Sheep, it give birth to a new Sheep each day.
  - Between **3**rd ~ **5**th day for Grass, it gives birth to a new Grass each day.
- Move
  - A Sheep can move to its 1-step nearby empty position.
  - **Nearby** which means they are neighbour, in anyone of 4 direction(East, South, West, North) that distance is 1.
- Eating:
  - If a Sheep is going to move to a position that standing a Grass, then the Sheep will eat this Grass. Eating a Grass will recover 5 healthPoints


