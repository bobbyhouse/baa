# baa
Simple AI powered command line utility that reads from stdin.

## categorize
You can categorize strings by providing them as stdin and specifying the categories to choose from.

Single line
```
> echo "Ice Cream" | baa healthy not-healthy
> Ice Cream:not-healthy
```

Multiple lines
```
awk '{print $0 | "baa healthy not-healthy" }' food.txt
Ice Cream:not-healthy
salad:healthy
hamburger:not-healthy
pizza:not-healthy
salmon:healthy
```

### Emergent properties
Because we're using LLMs to do the categorization you get some emergent properties for free.

Non-English
```
>> echo "manzana" | baa healthy not-healthy
> manzana:healthy
```

Unicode
```
>> echo "🍕" | baa healthy not-healthy
🍕:not-healthy
```
