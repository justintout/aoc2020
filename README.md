# Advent of Code 2020

## Running

Clone the project, build it, run. You'll need Go installed.

```
git clone github.com/justinout/aoc2020
cd aoc2020
go build
./aoc2020
```

To solve only a specific day, use the `-day` flag. To fetch inputs from [adventofcode.com] instead of reading files from the `./input/` directory, use the `-cookie` flag and pass your [adventofcode.com] session cookie. You can get this by inspecting the [adventofcode.com] page. 

## Inputs 

Puzzle inputs are different for everyone, and can be retrieved once you've logged in to [adventofcode.com]. 

This repo ships with my own puzzle inputs in `./input/`. If you'd like to overwrite the files in `./input/` with your inputs, you can run `getinputs.sh` and pass your cookie. This will download all of the unlocked inputs.

```
./getinputs.sh <COOKIE>
```

You can also pass the executable the `-cookie` flag to download the inputs in memory. This will not overwrite `./input/`

## Code Dependencies 

- [gonum/combin](https://godoc.org/gonum.org/v1/gonum/stat/combin#Combinations) - Easy combinatrics to get all the possible input combinations for Day 1

## Notes

### Pre-start

No real plan or language preference this year. I'd like to mess around with Pony some more but 
might just keep it normal. Going to focus on getting the most minimal answers I can muster but still
test ~~fully~~ decently well ;). 

$5 says we end up building a computer. 


### Day 1

Ah man 503s.

Heeeey we got inputs! Alright. Not the most elegant solution but we're 2 stars closer. Gonna massage
how the project looks a little now. 

#### Done 

I have a feeling I might need that "find a sum" kinda deal in the future. This is gonna go from split days
to a single library kinda deal. At least, that's what usually happens. Already have my many-small-files tendency
showing. Also, going to try to remember to add any sketches/notes I take on my reMarkable to the repo. 

### Day 2

Late start. Was watching "Flavorful Origins" on Netflix. 

I feel like the fun answer would be turning these into regex but I'm not doing that with Go. 

Oh and an internet outage on my end. Weird. 

Update: Ah crap. WOW says it's a 12-3 maintenance window. Almost got part 2 submitted but D/Ced right as I pressed the button :laugh: 

#### Done

That one was cute. I went hamfisted with the solution I think. Let's dumb this down a little.

### Day 3

#### Prerelease

Yeah, this feels a lot better. Also went back and really simplified the two solutions thus far. Need to re-test the helper functions for the days.

Hey why even bother with a shell script amirite just let he script get the inputs.

#### Done

I liked this puzzle too. I got stuck on a few fiddly off-by-ones tbh. I also forgot to reset my pointer in the part 2 loop for a few iterations. 

### Day 4

#### Part 1

Oof. The tailing blank line in the input was getting gobbled up and causing the last passport to be missed. Adding an extra whitespace fixed it. Majorly stuck on that one for a minute. Added a check for that last passport. Nasty little labeled continue there too :O

#### Part 2 

Off by 1 in the other direction?? There's a single passport slipping through the cracks somehow... 
