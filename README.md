# Advent of Code 2020

## Running

Clone the project, build it, run. You'll need Go installed.

```
git clone github.com/justinout/aoc2020
cd aoc2020
go build
./aoc2020
```

If you'd like to run with your inputs instead of mine, first get get all of your inputs using `getinputs.sh`. Inspect the Advent of Code page to get your session cookie.

```
./getinputs.sh <COOKIE> && ./aoc2020
```

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