#### Frequency - freq

Implements a count–min-log sketch, a probabilistic data structure that serves as a frequency table of events in a stream of data.

**Creating** a new empty sketch of type Count-Min-Log (freq) with the id "sketch_2" and a capacity of 1000000:
```{r, engine='bash', count_lines}
curl -XPOST http://localhost:3596/freq/sketch_2 -d '{
  "capacity": 1000000
}'
```

* optional arguments:
	* capacity: the max capacity of values (does not apply to hllpp), default is 1000000.


**Adding** values to the sketch with id "sketch_2":
```{r, engine='bash', count_lines}
curl -XPUT http://localhost:3596/freq/sketch_2 -d '{
  "values": ["marvel", "hulk", "marvel"]
}'
```

**Getting** the frequency for the values "marvel" and "hulk" in "sketch_2":
```{r, engine='bash', count_lines}
curl -XGET http://localhost:3596/freq/sketch_2 -d '{
  "values": ["marvel", "hulk"]
}'
```
returns the current count for each of these values:
```json
{
  "result":{
    "hulk":1,
    "marvel":2
  },
  "error":null
}
```

**Deleting** the sketch of type "freq" with id "sketch_2":
```{r, engine='bash', count_lines}
curl -XDELETE http://localhost:3596/freq/sketch_2
```